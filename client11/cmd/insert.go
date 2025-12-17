package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("SERVER")
		PORT := viper.GetString("PORT")

		dataset, _ := cmd.Flags().GetString("dataset")
		if dataset == "" {
			fmt.Println("data is empty!")
			return
		}

		val, _ := cmd.Flags().GetString("values")
		if val == "" {
			fmt.Println("No Data")
		}

		VALS := strings.Split(val, ",")
		vSend := ""
		for _, v := range VALS {
			_, err := strconv.ParseFloat(v, 64)
			if err == nil {
				vSend = vSend + "/" + v
			}
		}
		// create request in two steps for readability
		URL := "http://" + SERVER + ":" + PORT + "/insert/"
		URL = URL + "/" + dataset + "/" + vSend + "/"

		//send request
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println("**", err)
			return
		}

		// Check HTTP Status code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}
		// read
		resp, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println("***", err)
			return
		}
		fmt.Println(string(resp))
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	insertCmd.Flags().StringP("dataset", "d", "", "dataset name")
	insertCmd.Flags().StringP("values", "v", "", "List of values")
}
