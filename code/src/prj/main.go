package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"log"
	"strings"
)

func main() {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
			"--window-size=1280,1024",
			"--disable-dev-shm-usage",
			"--no-sandbox",
		}),
		agouti.Debug,
	)

	if err := driver.Start(); err != nil {
		log.Printf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	// Access to a target page
	url := "https://www.google.co.jp/"
	err = page.Navigate(url)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	// Wait for a DOM element to exist
	page.Session().SetImplicitWait(30)

	// Get screen shot
	page.Screenshot("Google.png")

	// Get title
	title, err := page.Title()
	if err != nil {
		log.Printf("Failed to get Title: %v", err)
	}
	log.Printf(title)

	// Get current url
	current_url, err := page.URL()
	if err != nil {
		log.Printf("Failed to get current url: %v", err)
	}
	log.Printf(current_url)

	// Get HTML
	getSource, err := page.HTML()
	//log.Printf(getSource)
	if err != nil {
		log.Fatalf("Failed to get HTML:%v", err)
	}

	// Parse DOM
	readerCurContents := strings.NewReader(getSource)
	doc, err := goquery.NewDocumentFromReader(readerCurContents)
	if err != nil {
		log.Fatal(err)
	}

	// Get title from DOM
	log.Printf(doc.Find("title").Text())
}
