package cmd

import (
	"fmt"
	"github.com/datactf/dataloader/pkg/loader"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "intializes database",
	Long:  "Initializes database with n documents",
	RunE:  RunInit,
}

func RunInit(cmd *cobra.Command, args []string) error {
	ndocs, _ := cmd.Flags().GetInt("ndocs")
	fmt.Printf("running the command: initializating with %d documents\n", ndocs)
	loader.LoadMongoDBAccounts(ndocs, "data", "accounts")
	return nil
}

func init() {

	initCmd.Flags().IntP("ndocs", "n", 100, "number of documents to intialize database")
	rootCmd.AddCommand(initCmd)
}
