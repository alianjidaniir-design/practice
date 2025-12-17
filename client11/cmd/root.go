/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client11",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("server", "S", "localhost", "Server")
	rootCmd.PersistentFlags().StringP("port", "P", "1234", "Port number")

	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

}
