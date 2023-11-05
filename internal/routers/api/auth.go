package api

import (
	"github.com/HarryLuo227/simple-blog-service/global"
	"github.com/HarryLuo227/simple-blog-service/internal/service"
	"github.com/HarryLuo227/simple-blog-service/pkg/app"
	"github.com/HarryLuo227/simple-blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// @Summary 取得 Token
// @Produce json
// @Param user header string true "使用者帳號"
// @Param password header int true "使用者密碼"
// @Success 200 {Object} model.Tag "成功"
// @Failure 400 {Object} errcode.Error "請求錯誤"
// @Failure 500 {Object} errcode.Error "內部錯誤"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.User, param.Password)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
