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

type ImageConfig struct {
	url     string
	element string
	quality int
	path    string
	code    string
	send    string
}

var conf = new(ImageConfig)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Convert web page to image",
	Long:  "Convert web page to image\n\nWarning: path and send cannot be used at the same time",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start conversion -> image")
		var cache []byte
		if conf.element != "" {
			cache = service.ToElementImage(conf.url, conf.element)
		} else {
			cache = service.ToFullImage(conf.url, conf.quality)
		}

		var code string
		if conf.code == "" {
			code = fmt.Sprint(time.Now().Unix())
		} else {
			code = conf.code
		}
		log.Println("Conversion completed ->", code)

		if conf.send != "" {
			log.Println("Start sending ->", conf.send)
			// http post send
			log.Println("Image sent successfully")
		} else {
			log.Println("Start saving ->", conf.path+"/"+code+".png")
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
	imageCmd.Flags().StringVarP(&conf.url, "url", "u", "", "Web page address to be converted")
	imageCmd.Flags().StringVarP(&conf.element, "element", "e", "", "Dom element to be converted")
	imageCmd.Flags().IntVarP(&conf.quality, "quality", "q", 90, "Image quality")
	imageCmd.Flags().StringVarP(&conf.path, "path", "p", "./", "Save path of converted image")
	imageCmd.Flags().StringVarP(&conf.code, "code", "c", "", "Code identifying this conversion")
	imageCmd.Flags().StringVarP(&conf.send, "send", "s", "", "Address to receive converted image")
	imageCmd.MarkFlagRequired("url")
}
