package v1

import (
	"github.com/HarryLuo227/simple-blog-service/global"
	"github.com/HarryLuo227/simple-blog-service/internal/service"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
	"github.com/HarryLuo227/simple-blog-service/pkg/convert"
	"github.com/HarryLuo227/simple-blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 取得多個文章
// @Produce json
// @Param title query string false "文章標題" maxlength(100)
// @Param tag_id query int true "標籤 ID"
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {Object} model.Article "成功"
// @Failure 400 {Object} errcode.Error "請求錯誤"
// @Failure 500 {Object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	response.ToResponseList(articles, totalRows)
	return
}

// @Summary 取得特定文章
// @Produce json
// @Param id path int true "文章 ID"
// @Success 200 {Object} model.Article "成功"
// @Failure 400 {Object} errcode.Error "請求錯誤"
// @Failure 500 {Object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	articleWithTags, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(&articleWithTags)
	return
}

// @Summary 新增文章
// @Produce json
// @Param tag_id body []int true "標籤 ID"
// @Param title body string true "文章標題" minlength(1) maxlength(100)
// @Param desc body string false "文章簡述" maxlength(255)
// @Param content body string true "文章內容" minlength(1) maxlength(4294967295)
// @Param cover_image_url body string true "封面圖片位址"
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param created_by body string true "建立者" minlength(1) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {Object} errcode.Error "請求錯誤"
// @Failure 500 {Object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	articleWithTags, err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(&articleWithTags)
	return
}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章 ID"
// @Param tag_id body int false "標籤 ID"
// @Param title body string false "文章標題" minlength(1) maxlength(100)
// @Param desc body string false "文章簡述" maxlength(255)
// @Param content body string false "文章內容" minlength(1) maxlength(4294967295)
// @Param cover_image_url body string false "封面圖片位址"
// @Param state body int false "狀態" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(1) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {Object} errcode.Error "請求錯誤"
// @Failure 500 {Object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	articleWithTags, err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(&articleWithTags)
	return
}

// @Summary 刪除文章
// @Produce json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {Object} errcode.Error "請求錯誤"
// @Failure 500 {Object} errcode.Error "內部錯誤"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse("Succeeded")
	return
}
