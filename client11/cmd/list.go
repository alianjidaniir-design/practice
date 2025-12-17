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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		//create Request
		URL := "http://" + SERVER + ":" + PORT + "/list"
		//send Request
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Cheak HTTP Status Code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status Code:", data.StatusCode)
			return
		}
		//Read request
		Resp, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("List of entries:")
		fmt.Println(string(Resp))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
