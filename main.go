package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sofosrobotics/auth"
	"sofosrobotics/handler"
	"sofosrobotics/helper"
	"sofosrobotics/robotmasters"
	"sofosrobotics/robotorders"
	"sofosrobotics/robotproductpartdevices"
	"sofosrobotics/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load env file ")
	}

	// data env.
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	// database local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	robotmasterRepository := robotmasters.NewRepository(db)
	robotmasterService := robotmasters.NewService(robotmasterRepository)
	robotmasterHandler := handler.RobotMasterHandler(robotmasterService)

	robotproductpartdeviceRepository := robotproductpartdevices.NewRepository(db)
	robotproductpartdeviceService := robotproductpartdevices.NewService(robotproductpartdeviceRepository)
	robotproductpartdeviceHandler := handler.RobotproductpartdeviceHandler(robotproductpartdeviceService)

	robotorderRepository := robotorders.NewRepository(db)
	robotorderService := robotorders.NewService(robotorderRepository)
	robotorderHandler := handler.NewRobotorderHandler(robotorderService)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("apisofosrobotics", cookieStore))

	router.Static("/images/", "./images")

	router.GET("/api/v1/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "success",
		})
	})

	// user
	router.POST("/api/v1/users/register/", userHandler.RegisterUser)
	router.POST("/api/v1/users/update/", userHandler.UpdateUser)
	router.POST("/api/v1/users/login/", userHandler.Login)
	router.POST("/api/v1/users/email_checkers/", userHandler.CheckEmailAvailability)
	router.POST("/api/v1/users/avatars/", authMiddleware(authService, userService), userHandler.UploadAvatar)
	router.GET("/api/v1/users/profile/", authMiddleware(authService, userService), userHandler.FetchUser)

	//robot master
	router.GET("/api/v1/robotmaster/recomendation/", robotmasterHandler.Recomendation)
	router.GET("/api/v1/robotmaster/robotics/", robotmasterHandler.FindAllRobotic)
	router.GET("/api/v1/robotmaster/automation/", robotmasterHandler.FindAllAutomation)
	router.POST("/api/v1/robotmaster/byid/", robotmasterHandler.Byid)

	//robot product part devices
	router.POST("/api/v1/robotproductpartdevice/all/", robotproductpartdeviceHandler.FindAllIDRobotMasterAndRobotPartDevice)
	router.POST("/api/v1/robotproductpartdevice/id/", robotproductpartdeviceHandler.FindAllIDRobotProductPartDevice)

	//robot orders
	router.POST("/api/v1/robotorders/save/", robotorderHandler.Save)
	router.POST("/api/v1/robotorders/userid/", robotorderHandler.ByUserid)
	router.POST("/api/v1/robotorders/kodeinvoice/", robotorderHandler.Kodeinvoice)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
