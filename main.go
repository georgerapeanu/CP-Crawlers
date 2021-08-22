package main

import (
	"fmt"
	"time"

	"github.com/georgerapeanu/CP-Crawlers/crawlers"
	"github.com/georgerapeanu/CP-Crawlers/generic"
)

func main() {
	generic.Init()
	crawler := crawlers.CodeforcesCrawler{}
	timeLocation := time.FixedZone("GMT+3", +3*60*60)
	var err error
	endTime, err := time.ParseInLocation("Jan/02/2006 15:04", "Aug/22/2021 16:23", timeLocation)
	if err != nil {
		panic(err)
	}
	startTime, err := time.ParseInLocation("Jan/02/2006 15:04", "Jun/01/2021 10:00", timeLocation)
	if err != nil {
		panic(err)
	}
	fmt.Println(crawler.ParseSubmissionTable("https://codeforces.com/submissions/georgerapeanu", startTime, endTime))
	//fmt.Println(crawler.ParseSubmissionPage("https://codeforces.com/submissions/georgerapeanu"))
	fmt.Println("ok")
}
