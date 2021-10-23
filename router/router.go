package router

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/topics/controllers"
	"github.com/topics/forms"
)

func Init() {
	// gin.SetMode(gin.ReleaseMode)
	//Start the default gin server
	r := gin.Default()

	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)
		v1.PUT("/user/renew/:uuid", user.Renew)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)
		//Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", auth.Refresh)

		/*** START DASHBOARD ***/
		dashboard := new(controllers.DashboardController)
		v1.GET("/dashboard/system/info", TokenAuthMiddleware(), dashboard.SystemInfo)
		v1.GET("/dashboard/system/info/history/:day", TokenAuthMiddleware(), dashboard.SystemInfoHistory)
		v1.GET("/dashboard/system/client/type/percentage", TokenAuthMiddleware(), dashboard.ClientTypePercentage)
		v1.GET("/dashboard/system/log/:line", TokenAuthMiddleware(), dashboard.SystemLog)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")

	// log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {
		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}
}
