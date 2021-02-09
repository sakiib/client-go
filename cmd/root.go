package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A brief description -> rootCmd",
	Long:  `A longer description -> rootCmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
