package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/contexts"
	"github.com/kensei18/enechain-technical-assignment/app/loggers"
	"github.com/kensei18/enechain-technical-assignment/app/storage"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

const (
	userTokenKeyName = "userToken"

	errorCodeUnauthorized        = "UNAUTHORIZED"
	errorCodeNotFound            = "NOT_FOUND"
	errorCodeDuplicatedKey       = "DUPLICATED_KEY"
	errorCodeInvalidInputField   = "INVALID_INPUT_FIELD"
	errorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
)

type userTokenKey string

type GraphQLServer struct {
	Port    string
	Schema  graphql.ExecutableSchema
	Logger  *loggers.RequestLogger
	Loaders *storage.Loaders
}

func (s *GraphQLServer) Serve() {
	srv := handler.NewDefaultServer(s.Schema)
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = graphql.WithResponseContext(ctx, graphqlErrorPresenter(s.Logger), graphql.DefaultRecover)
		return next(ctx)
	})
	srv.AroundOperations(graphqlLogHandler(s.Logger))
	srv.AroundOperations(graphqlAuthHandler())

	dataloadersHandler := dataloadersHandlerFunc(s.Loaders)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle(
		"/query",
		traceHandler(httpAuthHandler(dataloadersHandler(srv))),
	)

	slog.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", s.Port))
	log.Fatal(http.ListenAndServe(":"+s.Port, nil))

}

func traceHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.NewString()
		next.ServeHTTP(w, r.WithContext(contexts.WithTraceID(r.Context(), traceID)))
	})
}

func graphqlLogHandler(logger *loggers.RequestLogger) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		start := time.Now()

		res := next(ctx)

		operationCtx := graphql.GetOperationContext(ctx)
		attrs := []slog.Attr{
			slog.String("graphql_operation", operationCtx.OperationName),
			slog.String("graphql_query", operationCtx.RawQuery),
			slog.Duration("duration", time.Since(start)),
		}
		errs := graphql.GetErrors(ctx)
		if len(errs) > 0 {
			codes := make([]interface{}, 0, len(errs))
			for _, e := range errs {
				codes = append(codes, e.Extensions["codes"])
			}
			attrs = append(attrs, slog.Any("graphql_error_codes", codes))
		}
		logger.Info(ctx, "", attrs...)

		return res
	}
}

func httpAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		var key userTokenKey = userTokenKeyName
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), key, token)))
	})
}

func graphqlAuthHandler() graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		operationCtx := graphql.GetOperationContext(ctx)
		if operationCtx.OperationName == "IntrospectionQuery" {
			// skip authorization for playground docs
			return next(ctx)
		}
		var key userTokenKey = userTokenKeyName
		value := ctx.Value(key)
		token, ok := value.(string)
		if !ok {
			panic("invalid token")
		}
		// TODO: verify token
		userID, err := uuid.Parse(token)
		if err != nil {
			return func(ctx context.Context) *graphql.Response {
				return &graphql.Response{Errors: gqlerror.List{errorGraphqlUnauthorized()}}
			}
		}
		return next(contexts.WithUserID(ctx, userID))
	}
}

func dataloadersHandlerFunc(loaders *storage.Loaders) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(storage.SetLoaders(r.Context(), loaders)))
		})
	}
}

func graphqlErrorPresenter(logger *loggers.RequestLogger) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		var gqlErr *gqlerror.Error
		var orig error
		var code string
		if x, ok := err.(*gqlerror.Error); ok {
			gqlErr = x
			orig = errors.Unwrap(x)
		} else {
			gqlErr = graphql.DefaultErrorPresenter(ctx, err)
			orig = err
		}
		gqlErr.Extensions = map[string]interface{}{}
		switch true {
		case errors.Is(orig, gorm.ErrRecordNotFound):
			code = errorCodeNotFound
		case errors.Is(orig, gorm.ErrDuplicatedKey):
			code = errorCodeDuplicatedKey
		case errors.Is(orig, gorm.ErrInvalidField):
			code = errorCodeInvalidInputField
		default:
			code = errorCodeInternalServerError
			logger.Error(ctx, gqlErr.Message, slog.String("code", code))
		}
		gqlErr.Extensions["code"] = code
		return gqlErr
	}
}

func errorGraphqlUnauthorized() *gqlerror.Error {
	err := gqlerror.Wrap(errors.New("unauthorized"))
	err.Extensions = map[string]interface{}{}
	err.Extensions["code"] = errorCodeUnauthorized
	return err
}
