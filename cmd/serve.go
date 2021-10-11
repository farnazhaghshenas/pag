package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"pack-and-go/config"
	"pack-and-go/http"
	"pack-and-go/storage"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	e := echo.New()
	e.Use(middleware.Recover())

	configuration, err := config.NewViperConfiguration()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := storage.NewDBConnection(configuration.DB)
	if err != nil {
		logrus.Fatal("sql connection error : ", err.Error())
	}

	api := &http.API{
		Server: e,
		DB:     db,
	}

	api.Server.GET("/trip", api.GetTrips)
	api.Server.POST("/trip", api.CreateTrip)
	api.Server.GET("/trip/:id", api.GetTrip)

	go e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", configuration.GetAppConfig().Port)))
}
