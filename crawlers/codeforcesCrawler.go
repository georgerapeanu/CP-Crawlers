package crawlers

//TODO time zone stuff i guess

import (
	"errors"
	"fmt"
	"strings"

	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/georgerapeanu/CP-Crawlers/generic"
)

type CodeforcesCrawler struct {
}

func (crawler CodeforcesCrawler) GetSubmissions(handle string,
	taskLink string,
	beginningTime time.Time) ([]generic.Submission, error) {
	ans := make([]generic.Submission, 0)
	return ans, nil
}
func (crawler CodeforcesCrawler) ParseSubmission(submissionLink string) (generic.Submission, error) {

	timeLayout := "2006-01-02 15:04:05"

	ans := generic.Submission{}

	res, err := http.Get(submissionLink)

	if err != nil {
		return ans, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ans, errors.New("Error: GET submission link return NON-ACCEPTED response")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return ans, err
	}

	dataTable := goquery.NewDocumentFromNode(doc.Find(".datatable").Get(0))
	dataTable2 := goquery.NewDocumentFromNode(dataTable.Children().Get(5))
	dataTable3 := goquery.NewDocumentFromNode(dataTable2.Children().Get(2))
	preInfoRow := goquery.NewDocumentFromNode(dataTable3.Children().Get(0))
	infoRow := goquery.NewDocumentFromNode(preInfoRow.Children().Get(1))

	//getting submission link
	ans.SubmissionLink = submissionLink

	//getting handle
	userData := goquery.NewDocumentFromNode(infoRow.Children().Get(1))
	ans.Handle = strings.TrimSpace(goquery.NewDocumentFromNode(userData.Children().Get(1)).Text())

	//getting language
	ans.Language = strings.TrimSpace(goquery.NewDocumentFromNode(infoRow.Children().Get(3)).Text())

	//getting submission time
	ans.SubmissionTime, err = time.Parse(timeLayout, strings.TrimSpace(goquery.NewDocumentFromNode(infoRow.Children().Get(7)).Text()))
	if err != nil {
		return ans, err
	}

	//getting submission result
	ans.Result = strings.TrimSpace(goquery.NewDocumentFromNode(infoRow.Children().Get(4)).Text())

	//getting submission status
	if ans.Result == "Accepted" {
		ans.Status = "COMPLETED"
	} else {
		ans.Status = "ATTEMPTED"
	}
	//getting task & task link
	goquery.NewDocumentFromNode(infoRow.Children().Get(2)).Find("a").Each(func(i int, s *goquery.Selection) {
		sublink, _ := s.Attr("href")
		ans.TaskLink = "https://codeforces.com" + sublink
		name, _ := s.Attr("title")
		ans.Task = name
	})

	fmt.Printf("%+v\n", ans)
	//fmt.Println(taskData.Html())
	return ans, nil
}
