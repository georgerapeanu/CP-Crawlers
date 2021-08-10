package main

import (
	"fmt"

	"github.com/georgerapeanu/CP-Crawlers/crawlers"
	"github.com/georgerapeanu/CP-Crawlers/generic"
)

func main() {
	generic.Init()
	crawler := crawlers.CodeforcesCrawler{}
	_, err := crawler.ParseSubmission("https://codeforces.com/contest/1554/submission/12549059312312321")
	fmt.Println(err)
}
