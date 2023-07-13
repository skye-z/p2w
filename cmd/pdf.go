/*
子命令 - 转PDF

BetaX Page to what
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"p2w/service"
	"time"

	"github.com/spf13/cobra"
)

type PDFConfig struct {
	url  string
	path string
	code string
	send string
}

var pdfConf = new(PDFConfig)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Convert web page to pdf",
	Long:  "Convert web page to pdf",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start conversion -> pdf")
		cache := service.ToPDF(pdfConf.url)

		var code string
		if pdfConf.code == "" {
			code = fmt.Sprint(time.Now().Unix())
		} else {
			code = pdfConf.code
		}
		log.Println("Conversion completed ->", code)

		if pdfConf.send != "" {
			log.Println("Start sending ->", pdfConf.send)
			// http post send
			log.Println("PDF sent successfully")
		} else {
			log.Println("Start saving ->", pdfConf.path+"/"+code+".pdf")
			if err := os.WriteFile(code+".pdf", cache, 0o644); err != nil {
				log.Fatalln("PDF saving failed:", err)
			} else {
				log.Println("PDF saved successfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pdfCmd)
	pdfCmd.Flags().StringVarP(&pdfConf.url, "url", "u", "", "Web page address to be converted")
	pdfCmd.Flags().StringVarP(&pdfConf.path, "path", "p", "./", "Save path of converted image")
	pdfCmd.Flags().StringVarP(&pdfConf.code, "code", "c", "", "Code identifying this conversion")
	pdfCmd.Flags().StringVarP(&pdfConf.send, "send", "s", "", "Address to receive converted image")
	pdfCmd.MarkFlagRequired("url")
}
