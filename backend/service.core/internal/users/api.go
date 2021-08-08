package users

import (
	"github.com/gin-gonic/gin"

	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/entity"
	"github.com/COMP3900-9900-Capstone-Project/capstoneproject-comp3900-w16a-jamar/pkg/errors"
)

// RegisterHandlers sets up the api routes for the core users endpoints
func RegisterHandlers(r *gin.RouterGroup, service Service, authHandler gin.HandlerFunc) {
	res := resource{service}

	unauthGroup := r.Group("")
	unauthGroup.POST("/register", res.register)
	unauthGroup.POST("/login", res.login)
	unauthGroup.POST("/login/guest", res.loginGuest)
	unauthGroup.PUT("/reset", res.resetPassword)

	userGroup := r.Group("")
	userGroup.Use(authHandler)

	userGroup.POST("/logout", res.logout)
	userGroup.GET("/users/:id", res.get)
	userGroup.PUT("/users/:id", res.update)
}

type resource struct {
	service Service
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ItemResponse struct {
	Item *entity.User `json:"item"`
}

// @Router /register [post]
// @Tags Users
// @Summary Used to register a new user account
// @Param request body users.RegisterUserRequest true "The new user's details"
// @Success 200 {object} TokenResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 409 {object} errors.ErrorResponse "Duplicate Account"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) register(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	token, err := r.service.Register(c, req)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, TokenResponse{
		Token: token,
	})
}

// @Router /login [post]
// @Tags Users
// @Summary We all know what /login does
// @Param request body users.LoginRequest true "The user's login details"
// @Success 200 {object} TokenResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 409 {object} errors.ErrorResponse "Duplicate Account"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	token, err := r.service.Login(c, req)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, TokenResponse{
		Token: token,
	})
}

// @Router /login/guest [post]
// @Tags Users
// @Summary We all know what /login does, this does that for a guest
// @Success 200 {object} TokenResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 409 {object} errors.ErrorResponse "Duplicate Account"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) loginGuest(c *gin.Context) {
	token, err := r.service.LoginGuest(c)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, TokenResponse{
		Token: token,
	})
}

// @Router /logout [post]
// @Tags Users
// @Summary Logs out the current user and invalidates the session token
// @Success 200 {object} errors.ErrorResponse "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 409 {object} errors.ErrorResponse "Duplicate Account"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) logout(c *gin.Context) {
	if err := r.service.Logout(c); err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

// @Router /users/{id} [get]
// @Tags Users
// @Summary Returns the details for the specified user
// @Param id path string true "The id of the user, or current for the currently authenticated user"
// @Success 200 {object} entity.User "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) get(c *gin.Context) {

	userID := c.Param("id")

	if userID == "current" {
		userID = c.GetString("userID") // our authenticated user id based on the auth token

	} else if c.GetString("userID") != userID {
		errors.Abort(c, errors.Forbidden(""))
		return
	}

	user, err := r.service.Get(c, userID)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: user,
	})
}

// @Router /users/{id} [put]
// @Tags Users
// @Summary Updates a user's details
// @Param id path string true "The id of the user, or current for the currently authenticated user"
// @Success 200 {object} entity.User "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) update(c *gin.Context) {
	userID := c.Param("id")
	if userID == "current" {
		userID = c.GetString("userID") // our authenticated user id based on the auth token
	} else if c.GetString("userID") != userID {
		errors.Abort(c, errors.Forbidden(""))
		return
	}

	var req UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	user, err := r.service.Update(c, req)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: user,
	})
}

// @Router /reset [put]
// @Tags Users
// @Summary Used to update the password of a user
// @Param request body users.ResetPasswordRequest true "The new user's details"
// @Success 200 {object} entity.User "success"
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 409 {object} errors.ErrorResponse "Duplicate Account"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
func (r resource) resetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.BindJSON(&req); err != nil {
		errors.Abort(c, errors.BadRequest(err.Error()))
		return
	}

	user, err := r.service.ResetPassword(c, req)
	if err != nil {
		errors.Abort(c, err)
		return
	}

	c.JSON(200, ItemResponse{
		Item: user,
	})
}
