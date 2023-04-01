package main

import (
	"fmt"
	"github.com/go-rod/rod"
)

func main() {
	// launch a new browser with rod
	browser := rod.New().MustConnect()

	// navigate to the Ticketmaster page you want to scrape
	page := browser.MustPage("https://www.ticketmaster.ca/discover/concerts").MustWaitLoad()

	// wait for the page to load completely
	page.MustWaitLoad()

	// get a list of all the event names
	names := page.MustElements("span[class='event-title__text']")

	// print out the names to the console
	for _, name := range names {
		fmt.Println(name)
	}
}
