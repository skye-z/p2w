/*
子命令 - 转图片

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

type ImageConfig struct {
	url     string
	element string
	quality int
	path    string
	code    string
	send    string
}

var imgConf = new(ImageConfig)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Convert web page to image",
	Long:  "Convert web page to image\n\nWarning: path and send cannot be used at the same time",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start conversion -> image")
		var cache []byte
		if imgConf.element != "" {
			cache = service.ToElementImage(imgConf.url, imgConf.element)
		} else {
			cache = service.ToFullImage(imgConf.url, imgConf.quality)
		}

		code := global.GetCode(pdfConf.code)
		log.Println("Conversion completed ->", code)

		if imgConf.send != "" {
			log.Println("Start sending ->", imgConf.send)
			// http post send
			log.Println("Image sent successfully")
		} else {
			log.Println("Start saving ->", imgConf.path+"/"+code+".png")
			if err := os.WriteFile(code+".png", cache, 0o644); err != nil {
				log.Fatalln("Image saving failed:", err)
			} else {
				log.Println("Image saved successfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().StringVarP(&imgConf.url, "url", "u", "", "Web page address to be converted")
	imageCmd.Flags().StringVarP(&imgConf.element, "element", "e", "", "Dom element to be converted")
	imageCmd.Flags().IntVarP(&imgConf.quality, "quality", "q", 90, "Image quality")
	imageCmd.Flags().StringVarP(&imgConf.path, "path", "p", "./", "Save path of converted image")
	imageCmd.Flags().StringVarP(&imgConf.code, "code", "c", "", "Code identifying this conversion")
	imageCmd.Flags().StringVarP(&imgConf.send, "send", "s", "", "Address to receive converted image")
	imageCmd.MarkFlagRequired("url")
}
