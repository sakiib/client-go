package cmd

import (
	"fmt"

	"github.com/sakiib/client-go/api"

	"github.com/spf13/cobra"
)

var replica int32
var image string
var name string

var createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "A brief description -> createDeploymentCmd",
	Long:  `A longer description -> createDeploymentCmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create deployment called")
		api.CreateDeployment(replica, image, name)
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

var updateDeploymentCmd = &cobra.Command{
	Use:   "update-deployment",
	Short: "A brief description -> updateDeploymentCmd",
	Long:  `A longer description -> updateDeploymentCmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update deployment called")
		api.UpdateDeployment(replica, image, name)
	},
}

var deleteDeploymentCmd = &cobra.Command{
	Use:   "delete-deployment",
	Short: "A brief description -> deleteDeploymentCmd",
	Long:  `A longer description -> deleteDeploymentCmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete deployment called")
		api.DeleteDeployment(name)
	},
}

func init() {
	rootCmd.AddCommand(createDeploymentCmd)
	rootCmd.AddCommand(getDeploymentCmd)
	rootCmd.AddCommand(updateDeploymentCmd)
	rootCmd.AddCommand(deleteDeploymentCmd)
	createDeploymentCmd.PersistentFlags().Int32VarP(&replica, "replica", "r", 1, "This flag sets the number of replicas")
	createDeploymentCmd.PersistentFlags().StringVarP(&image, "image", "i", "apiserver:1.0.1", "This flag sets the image name with version")
	createDeploymentCmd.PersistentFlags().StringVarP(&name, "name", "n", "apiserver", "This flag sets the deployment name to be created")
	updateDeploymentCmd.PersistentFlags().Int32VarP(&replica, "replica", "r", 1, "This flag sets the number of replicas")
	updateDeploymentCmd.PersistentFlags().StringVarP(&image, "image", "i", "apiserver:1.0.1", "This flag sets the image name with version")
	updateDeploymentCmd.PersistentFlags().StringVarP(&name, "name", "n", "apiserver", "This flag sets the deployment name to be updated")
	deleteDeploymentCmd.PersistentFlags().StringVarP(&name, "name", "n", "apiserver", "This flag sets the deployment name to be deleted")
}
