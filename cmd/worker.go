/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/luoruofeng/DockerApiAgent/model"
	"github.com/spf13/cobra"
)

var swarmRemoteAddr *string

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:     "worker",
	Aliases: []string{"w"},
	Args:    cobra.ExactArgs(1),
	Short:   "设置docker swarm为worker节点",
	Long:    `代理配置本机docker进程为worker模式节点，需要设置master token。`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("worker没有设置token")
		}
		if swarmRemoteAddr == nil || *swarmRemoteAddr == "" {
			panic("worker没有设置swarmRemoteAddr")
		}
		model.Cnf.SwarmToken = args[0]
		model.Cnf.SwarmRemoteAddr = *swarmRemoteAddr
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
	swarmRemoteAddr = workerCmd.Flags().StringP("swarm-remote-addr", "r", "", "master的内网IP")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
