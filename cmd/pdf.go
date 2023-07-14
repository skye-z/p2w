/*
子命令 - 转PDF

BetaX Page To What
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"log"
	"os"
	"p2w/global"
	"p2w/service"

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

		code := global.GetCode(pdfConf.code)
		log.Println("Conversion completed ->", code)
		name := code + ".pdf"

		if pdfConf.send != "" {
			log.Println("Start sending ->", pdfConf.send)
			global.SendFile(pdfConf.send, cache, name)
			log.Println("PDF sent successfully")
		} else {
			log.Println("Start saving ->", pdfConf.path+"/"+name)
			if err := os.WriteFile(name, cache, 0o644); err != nil {
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
