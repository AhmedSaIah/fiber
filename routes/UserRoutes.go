package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/AhmedSaIah/fiber/controllers"
	"github.com/AhmedSaIah/fiber/middleware"
	"github.com/AhmedSaIah/fiber/services"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {
	//TODO: /users or users with no slash
	router := rg.Group("/users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
