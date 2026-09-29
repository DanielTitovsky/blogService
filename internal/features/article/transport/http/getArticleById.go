package article_http_router

import (
	core_transport_http_utils "blog-service/internal/core/transport/http/utils"
	article_usecases "blog-service/internal/features/article/usecases"
	"net/http"
)

func (a *ArticleRouter) GetArticleById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	responceHandler := core_transport_http_utils.NewHTTPResponceHandler(w)

	requestArticle, err := core_transport_http_utils.GetAndFilterJson[article_usecases.ArticleId](a.validator, r)

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Failed validate data", err.Error())
		return
	}

	article, err := a.services.GetArticleById(ctx, requestArticle.ArticleId)

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Not faind article", err.Error())
		return
	}

	responceHandler.Success(http.StatusAccepted, article)
}
