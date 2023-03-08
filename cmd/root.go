/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/luoruofeng/DockerApiAgent/model"
	"github.com/spf13/cobra"
)

var Config *string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func LoadCnf() {
	if Config == nil || *Config == "" {
		fmt.Println("配置文件使用默认路径")
	} else {
		fmt.Println("configFile:" + *Config)
	}
	model.CreateConfig(*Config)
	model.Print(*Config)
}

func init() {
	Config = rootCmd.PersistentFlags().StringP("config", "c", "", "指定配置文件")
}
