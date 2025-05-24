package routes

import (
	"net/http"

	"golang_starter_kit_2025/app/controllers"
	"golang_starter_kit_2025/app/middleware"
	"golang_starter_kit_2025/app/services"
	"golang_starter_kit_2025/facades"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	controller := controllers.Controller{}
	route.GET("", controller.HelloWorld)

	authService := services.AuthService{}
	authController := controllers.NewAuthController(authService)
	route.PUT("/auth/login", authController.Login)
	authRoutes := route.Group("/auth").Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/logout", authController.Logout)
		authRoutes.GET("/refresh", authController.Refresh)
	}

	userService := services.UserService{}
	userController := controllers.NewUserController(userService)
	userRoutes := route.Group("/users", middleware.AuthMiddleware())
	{
		userRoutes.GET("", userController.List)
		userRoutes.GET("/:id", userController.Get)
		userRoutes.PUT("", userController.Put)
		userRoutes.DELETE("/:id", userController.Delete)
		userRoutes.POST("/:id/roles", userController.AssignRoles)
		userRoutes.GET("/:id/roles", userController.GetRoles)
	}

	roleService := services.RoleService{}
	roleController := controllers.NewRoleController(roleService)
	roleRoutes := route.Group("/roles", middleware.AuthMiddleware())
	{
		roleRoutes.GET("", roleController.List)
		roleRoutes.PUT("", roleController.Put)
		roleRoutes.DELETE("/:id", roleController.Delete)
		roleRoutes.POST("/:id/permissions", roleController.AssignPermissions)
		roleRoutes.GET("/:id/permissions", roleController.GetPermissions)
	}

	permissionService := services.PermissionService{}
	permissionController := controllers.NewPermissionController(permissionService)
	permissionRoutes := route.Group("/permissions", middleware.AuthMiddleware())
	{
		permissionRoutes.GET("", permissionController.List)
		permissionRoutes.PUT("", permissionController.Put)
		permissionRoutes.DELETE("/:id", permissionController.Delete)
	}

	farmerService := services.FarmerService{}
	farmerController := controllers.NewFarmerController(farmerService)
	farmerRoutes := route.Group("/farmers", middleware.AuthMiddleware())
	{
		farmerRoutes.GET("", farmerController.List)
		farmerRoutes.GET("/:id", farmerController.Get)
		farmerRoutes.PUT("", farmerController.Put)
		farmerRoutes.DELETE("/:id", farmerController.Delete)
	}

	fileController := controllers.NewFileController()
	fileRoutes := route.Group("/file")
	{
		fileRoutes.GET("/:key/:filename", fileController.ServeFile)
	}

	route.GET("/health", func(c *gin.Context) {
		sqlDB, err := facades.DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get facades connection",
				"error":   err.Error(),
			})
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "facades connection failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "facades is connected",
			"facades": "supply_chain_retail",
		})
	})
}
