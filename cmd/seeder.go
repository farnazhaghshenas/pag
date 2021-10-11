package cmd

import (
	"github.com/sirupsen/logrus"
	"pack-and-go/config"
	"pack-and-go/models"
	"pack-and-go/storage"

	"github.com/spf13/cobra"
)

// seederCmd represents the seeder command
var seederCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed cities in database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		seeder()
	},
}

func init() {
	rootCmd.AddCommand(seederCmd)
}

var cities = []string{"Barcelona", "Seville", "Madrid", "Valencia", "Andorra la Vella", "Malaga"}

func seeder(){
	configuration, err := config.NewViperConfiguration()
	if err != nil {
		logrus.Fatal(err)
	}


	db, err := storage.NewDBConnection(configuration.DB)
	if err != nil {
		logrus.Fatal("SQL Database Connection Error: ", err.Error())
	}

	for _, city := range cities {
		_, err := db.CreateCity(models.City{Name: city})
		if err != nil {
			logrus.Fatal(err.Error())
		}
	}
}