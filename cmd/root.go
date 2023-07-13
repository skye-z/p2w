/*
主命令

BetaX Page to what
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"p2w/global"

	"github.com/spf13/cobra"
)

// 没有任何子命令时调用的基本命令
var rootCmd = &cobra.Command{
	Use:     "p2w",
	Short:   "Convert web page to PDF or image",
	Version: global.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I don't know what I can do? Please use -h to view help :)")
	},
}

// 将所有子命令导入主命令并设置标识
// 此函数由 main 调用,只执行一次
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "prints version")
}
