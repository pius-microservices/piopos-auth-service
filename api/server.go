package serve

import (
	"github.com/pius-microservices/piopos-auth-service/api/routes"
	envConfig "github.com/pius-microservices/piopos-auth-service/config"
	"github.com/pius-microservices/piopos-auth-service/package/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "For Running api server",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	return c
}

func serve(cmd *cobra.Command, args []string) error {
	envCfg := envConfig.LoadConfig()

	errorLogger, debugLogger := utils.InitLogger()
	debugLogger.Println("Starting server...")

	mainRoute := gin.Default()
	if err := routes.RouteApp(mainRoute); err != nil {
		errorLogger.Println("Failed to initialize route: ", err)
		return err
	}

	c := corsHandler()
	handler := c.Handler(mainRoute)

	debugLogger.Printf("Server running on port %s", envCfg.Port)
	if err := http.ListenAndServe(":"+envCfg.Port, handler); err != nil {
		errorLogger.Println("Failed to start server: ", err)
		return err
	}

	return nil
}
