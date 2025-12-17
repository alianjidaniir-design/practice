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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		dataset, _ := cmd.Flags().GetString("dataset")
		if dataset == "" {
			fmt.Println("Please enter the dataset flag")
			return
		}

		// create request
		URL := "http://" + SERVER + ":" + PORT + "/search/" + dataset
		//send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println("Error in send request", err)
			return
		}

		// Check HTTP status code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code", data.StatusCode)
			return
		}
		//Read data
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println("Error reading response", err)
		}
		fmt.Println(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("dataset", "d", "", "dataset name to search")
}
