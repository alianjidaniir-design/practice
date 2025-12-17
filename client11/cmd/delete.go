/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		port := viper.GetString("port")

		dataset, _ := cmd.Flags().GetString("dataset")
		if dataset == "" {
			fmt.Println("Please provide a dataset name")
			return
		}
		//create
		URL := "http://" + SERVER + ":" + port + "/delete/" + dataset
		// send
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Check HTTP
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status Code:", data.StatusCode)
			return
		}
		//Read
		resp, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(resp))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("dataset", "d", "", "dataset name to delete ")
}
