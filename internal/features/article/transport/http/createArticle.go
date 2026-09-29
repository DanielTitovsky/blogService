package article_http_router

import (
	core_transport_http_utils "blog-service/internal/core/transport/http/utils"
	article_usecases "blog-service/internal/features/article/usecases"
	"net/http"
)

func (a *ArticleRouter) CreateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	responceHandler := core_transport_http_utils.NewHTTPResponceHandler(w)

	article, err := core_transport_http_utils.GetAndFilterJson[article_usecases.CreateArticleRequest](a.validator, r)

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Failed validate data", err.Error())
		return
	}

	createdArticle, err := a.services.CreateArticle(ctx, article_usecases.CreateArticleService{
		AuthorID: article.AuthorID,
		Title:    article.Title,
		Content:  article.Content,
	})

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Failed to create article", err.Error())
		return
	}

	responceHandler.Success(http.StatusAccepted, createdArticle)
}
