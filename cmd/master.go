/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/luoruofeng/DockerApiAgent/model"
	"github.com/luoruofeng/DockerApiAgent/util"

	"github.com/spf13/cobra"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:     "master",
	Aliases: []string{"m"},
	Args:    cobra.ExactArgs(1),
	Short:   "设置docker swarm为master节点",
	Long:    `代理配置本机docker进程为master模式节点，需要设置广播地址。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("master called")
		if len(args) < 1 {
			panic("master没有设置advertise-addr")
		}
		isIpv4 := util.CheckIPv4(args[0])
		if !isIpv4 {
			panic("master设置advertise-addr必须为IPv4格式")
		}
		LoadCnf()
		model.Cnf.AdvertiseAddr = args[0]
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
}
