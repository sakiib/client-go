package cmd

import (
	"fmt"
	"github.com/sakiib/client-go/api"

	"github.com/spf13/cobra"
)

var createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "A brief description -> createDeploymentCmd",
	Long:  `A longer description -> createDeploymentCmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create deployment called")
		api.CreateDeployment()
	},
}

var getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "A brief description -> getDeploymentCmd",
	Long:  `A longer description -> getDeploymentCmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get deployment called")
		api.GetDeployment()
	},
}

func init() {
	rootCmd.AddCommand(createDeploymentCmd)
	rootCmd.AddCommand(getDeploymentCmd)
}
