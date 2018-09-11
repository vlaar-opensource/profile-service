package cmd

import (
	"fmt"
	"os"

	"github.com/vlaar-opensource/profile-service/db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run service",
	Run: func(cmd *cobra.Command, args []string) {
		DB, err := db.NewDBPostresImpl("localhost", 5432, "postgres", "mydb", "secret")
		defer DB.Close()
		if err != nil {
			fmt.Println("whoops database error", err)
			os.Exit(22)
		}
		DB.MigrateAll()

	},
}
