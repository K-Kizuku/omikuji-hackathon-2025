package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/K-Kizuku/pymon-graphql/internal/app/repository"
	"github.com/K-Kizuku/pymon-graphql/internal/graph"
	"github.com/K-Kizuku/pymon-graphql/pkg/config"
	"github.com/K-Kizuku/pymon-graphql/pkg/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// DI
	dbConfig := config.NewDBConfig()
	db := db.NewDB(dbConfig)
	userRepo := repository.NewUserRepository(db)
	pythonRepo := repository.NewPythonRepository(db)
	pythonStatRepo := repository.NewPythonStatRepository(db)
	pythonSkillRepo := repository.NewPythonSkillRepository(db)

	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserRepository:        userRepo,
		PythonRepository:      pythonRepo,
		PythonStatRepository:  pythonStatRepo,
		PythonSkillRepository: pythonSkillRepo,
	}}))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
