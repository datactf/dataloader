package cmd

import (
	"fmt"
	"github.com/datactf/dataloader/pkg/loader"
	"github.com/spf13/cobra"
	"time"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs the loader",
	Long:  "runs the loader for a designated period of time",
	RunE:  RunRun,
}

func RunRun(cmd *cobra.Command, args []string) error {
	seconds, _ := cmd.Flags().GetInt("seconds")
	stopAt := time.Now().Add(time.Second * time.Duration(seconds))
	fmt.Printf("Stopping at %s\n", stopAt)
	for {
		if time.Now().After(stopAt) {
			fmt.Printf("stopping at %s", time.Now())
			return nil
		}
		if err := loader.LoadMongoDBAccounts(1, "data", "accounts"); err != nil {
			return err
		}

	}
	return nil
}

func init() {
	runCmd.Flags().IntP("seconds", "s", 30, "time in seconds to execute")
	rootCmd.AddCommand(runCmd)
}
