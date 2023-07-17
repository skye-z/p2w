/*
子命令 - Http服务

BetaX Page To What
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"p2w/global"
	"p2w/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type ServerConfig struct {
	port int
}

var serConf = new(ServerConfig)

// serverCmd represents the pdf command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "HTTP server",
	Long:  "HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		const portTag = "server.port"
		// 初始化配置
		global.InitConfig()
		// 关闭调试
		gin.SetMode(gin.ReleaseMode)
		// 判断端口号是否可用
		if serConf.port == 0 {
			serConf.port = global.GetInt(portTag)
		} else if serConf.port != global.GetInt(portTag) {
			global.Set(portTag, serConf.port)
		}
		// 加载路由
		log.Println("HTTP server startup in progress ->", serConf.port)
		r := gin.Default()
		http := service.HTTP{}
		r.GET("/api/pdf", http.ApiToPDF)
		r.GET("/api/img", http.ApiToImage)
		// 启动服务
		r.Run(fmt.Sprint(":", serConf.port))
		log.Println("HTTP server has stopped")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&serConf.port, "port", "p", 0, "HTTP server port")
}
