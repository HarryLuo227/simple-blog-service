package routers

import (
	"net/http"
	"time"

	_ "github.com/HarryLuo227/simple-blog-service/docs" // 這行要加，否則驗證時會噴 500 錯誤
	"github.com/HarryLuo227/simple-blog-service/global"
	"github.com/HarryLuo227/simple-blog-service/internal/middleware"
	"github.com/HarryLuo227/simple-blog-service/internal/routers/api"
	v1 "github.com/HarryLuo227/simple-blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", api.GetAuth)

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// Tags
		apiv1.GET("/tags", tag.List)
		apiv1.GET("/tags/:id", tag.Get)
		apiv1.POST("/tags", tag.Create)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)

		// Articles
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.POST("/articles", article.Create)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
