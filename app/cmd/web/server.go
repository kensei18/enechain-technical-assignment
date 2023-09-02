package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/contexts"
	"github.com/kensei18/enechain-technical-assignment/app/graph/web"
	"github.com/kensei18/enechain-technical-assignment/app/graph/web/resolver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultPort = "8080"

	userTokenKeyName = "userToken"
)

type userTokenKey string

func main() {
	// TODO: logger config

	db, err := gorm.Open(postgres.Open("dbname=app host=localhost port=5432 user=postgres password=password sslmode=disable"))
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(web.NewExecutableSchema(web.Config{Resolvers: &resolver.Resolver{DB: db}}))
	srv.AroundOperations(graphqlAuthHandler())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", httpAuthHandler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
