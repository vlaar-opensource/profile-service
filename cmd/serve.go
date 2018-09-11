package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vlaar-opensource/profile-service/app"
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "version",
	Short: "run service",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("APP_PORT")

		if port == "" {
			port = ":8000"
		}
		iport, err := strconv.Atoi(port)
		if err == nil {
			port = fmt.Sprintf(":%d", iport)
		}

		app := app.NewApp()

		app.Initialize()
		app.Run(port)
	},
}
