package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/contexts"
	"github.com/kensei18/enechain-technical-assignment/app/graph/web"
	"github.com/kensei18/enechain-technical-assignment/app/graph/web/resolver"
	"github.com/kensei18/enechain-technical-assignment/app/loggers"
	"github.com/kensei18/enechain-technical-assignment/app/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultPort = "8080"

	userTokenKeyName = "userToken"
)

type userTokenKey string

func main() {
	logger := loggers.NewDefaultLogger(os.Stdout, slog.LevelDebug)

	db, err := gorm.Open(
		postgres.Open("dbname=app host=localhost port=5432 user=postgres password=password sslmode=disable"),
		&gorm.Config{Logger: loggers.NewGormLogger(logger)},
	)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(web.NewExecutableSchema(web.Config{Resolvers: &resolver.Resolver{DB: db}}))
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = graphql.WithResponseContext(ctx, graphql.DefaultErrorPresenter, graphql.DefaultRecover)
		return next(ctx)
	})
	srv.AroundOperations(graphqlLogHandler(logger))
	srv.AroundOperations(graphqlAuthHandler())

	loaders := storage.NewLoaders(&storage.Reader{DB: db})
	dataloadersHandler := dataloadersHandlerFunc(loaders)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle(
		"/query",
		traceHandler(httpAuthHandler(dataloadersHandler(srv))),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
			panic(err)
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
