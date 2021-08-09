package main

import (
	"fmt"

	"github.com/georgerapeanu/CP-Crawlers/crawlers"
)

func main() {
	crawler := crawlers.CodeforcesCrawler{}
	_, err := crawler.ParseSubmission("https://codeforces.com/contest/1554/submission/124237552")
	fmt.Println(err)
}
