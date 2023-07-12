package service

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// 截取页面输出PDF
func ToPDF(url string) []byte {
	var cache []byte
	workBox(toPDF(url, &cache))
	return cache
}

// 截取完整页面输出图片
func ToFullImage(url string, quality int) []byte {
	var cache []byte
	workBox(toImage(url, "", quality, &cache))
	return cache
}

// 截取页面元素输出图片
func ToElementImage(url string, element string) []byte {
	var cache []byte
	workBox(toImage(url, element, 100, &cache))
	return cache
}

// 封装公共作业函数
func workBox(task chromedp.Tasks) {
	// 创建上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	chromedp.AtLeast(0)
	// 截取页面转换PDF
	if err := chromedp.Run(ctx, task); err != nil {
		log.Fatal(err)
	}
}

// 页面地址、字节指针
func toPDF(url string, bp *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cache, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*bp = cache
			return nil
		}),
	}
}

// 页面地址、截取元素、图片质量、字节指针
func toImage(url string, element string, quality int, bp *[]byte) chromedp.Tasks {
	if element == "" {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.FullScreenshot(bp, quality),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.Screenshot(element, bp, chromedp.NodeVisible),
		}
	}
}
