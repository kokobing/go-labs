package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs.exe"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	//打开url
	if err := page.Navigate("https://www.baidu.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	//截图
	if err := page.Screenshot("E:/Workspace/go-labs/src/lab123/lab001/tmp/phantomjs_baidu.jpg"); err != nil {
		log.Fatalf("Failed to screenshot:%v", err)
	}
}
