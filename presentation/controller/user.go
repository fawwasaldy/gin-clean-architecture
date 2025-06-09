package controller

import (
	"github.com/gin-gonic/gin"
	"kpl-base/application/request"
	"kpl-base/application/service"
	"kpl-base/platform/pagination"
	"kpl-base/presentation"
	"kpl-base/presentation/message"
	"net/http"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Me(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var req request.UserRegister
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	result, err := c.userService.Register(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedRegister, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessRegister, result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *userController) Login(ctx *gin.Context) {
	var req request.UserLogin
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.Verify(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedLogin, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessLogin, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Me(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)

	result, err := c.userService.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetUser, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessGetUser, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) RefreshToken(ctx *gin.Context) {
	var req request.RefreshToken
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.RefreshToken(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedRefreshToken, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusForbidden, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessRefreshToken, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetAll(ctx *gin.Context) {
	var req pagination.Request
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.GetAllUsersWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetAllUsers, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.Response{
		Status:  true,
		Message: message.SuccessGetAllUsers,
		Data:    result.Data,
		Meta:    result.Response,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Update(ctx *gin.Context) {
	var req request.UserUpdate
	if err := ctx.ShouldBind(&req); err != nil {
		res := presentation.BuildResponseFailed(message.FailedGetDataFromBody, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.MustGet("user_id").(string)

	result, err := c.userService.Update(ctx.Request.Context(), userID, req)
	if err != nil {
		res := presentation.BuildResponseFailed(message.FailedUpdateUser, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessUpdateUser, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Delete(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(string)

	if err := c.userService.Delete(ctx.Request.Context(), userID); err != nil {
		res := presentation.BuildResponseFailed(message.FailedDeleteUser, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := presentation.BuildResponseSuccess(message.SuccessDeleteUser, nil)
	ctx.JSON(http.StatusOK, res)
}
