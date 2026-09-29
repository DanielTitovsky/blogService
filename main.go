package main

import (
	core_repository_postgres "blog-service/internal/core/repository/postgres"
	core_transport_http_server "blog-service/internal/core/transport/http/server"
	repository_postgres_article "blog-service/internal/features/article/repository/postgres"
	article_services "blog-service/internal/features/article/services"
	article_http_router "blog-service/internal/features/article/transport/http"
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	serverConfig := core_transport_http_server.NewServerMuxConfig()
	serverMux := core_transport_http_server.NewServerMux(serverConfig)

	router := core_transport_http_server.NewRouter("v01/", serverMux.ServerMux)

	postgresConfig := core_repository_postgres.NewConfig()
	pgxPool := core_repository_postgres.NewPgxConnectinPool(ctx, postgresConfig)

	articleRepository := repository_postgres_article.NewArticleRepository(pgxPool.Pool)

	articleService := article_services.NewArticleService(&articleRepository)

	validate := validator.New(validator.WithRequiredStructEnabled())

	articleHandles := article_http_router.NewArticleRouters(&articleService, validate)

	router.RegisterRouter(articleHandles.Routers()...)

	fmt.Print("server started\n")

	err := serverMux.Run(ctx)

	if err != nil {
		fmt.Print(err)
		panic("Failed server started")
	}
}
