package article_http_router

import (
	core_transport_http_utils "blog-service/internal/core/transport/http/utils"
	article_usecases "blog-service/internal/features/article/usecases"
	"net/http"
)

func (a *ArticleRouter) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	responceHandler := core_transport_http_utils.NewHTTPResponceHandler(w)

	requestArticle, err := core_transport_http_utils.GetAndFilterJson[article_usecases.ArticleRequest](a.validator, r)

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Failed validate data", err.Error())
		return
	}

	updatedArticle, err := a.services.UpdateArticle(ctx, article_usecases.ArticleService{
		Id:       requestArticle.Id,
		AuthorID: requestArticle.AutorId,
		Title:    requestArticle.Title,
		Content:  requestArticle.Content,
	})

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Failed to updated article", err.Error())
		return
	}

	responceHandler.Success(http.StatusAccepted, updatedArticle)
}
