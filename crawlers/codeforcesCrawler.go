package crawlers

///TODO maybe investigate the status ajax
///for now to get the submissions between 2 timestamps you have to fetch all submissions
///also write tests

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/georgerapeanu/CP-Crawlers/generic"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

var CodeforcesCrawlerErrorExtractingSubmissionLink = errors.New("CodeforcesCrawler: Error extracting submission link")
var CodeforcesCrawlerErrorExtractingTaskLink = errors.New("CodeforcesCrawler: Error extracting task link")
var CodeforcesCrawlerErrorExtractingLastPage = errors.New("CodeforcesCrawler: Error extracting last page")

type CodeforcesCrawler struct {
}

func myUrlJoin(urlString, subroute string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, subroute)
	return u.String(), nil
}

func (crawler CodeforcesCrawler) GetSubmissions(handle string,
	beginTime time.Time,
	endTime time.Time) ([]generic.Submission, error) {

	submissionTableLink, err := myUrlJoin("https://codeforces.com/submissions", handle)
	if err != nil {
		return make([]generic.Submission, 0), err
	}

	return crawler.ParseSubmissionTable(submissionTableLink, beginTime, endTime)
}

func (crawler CodeforcesCrawler) GetSubmissionsForTask(
	handle string,
	taskLink string,
	beginTime time.Time,
	endTime time.Time) ([]generic.Submission, error) {
	ans, err := crawler.GetSubmissions(handle, beginTime, endTime)

	if err != nil {
		return ans, err
	}

	tmp := make([]generic.Submission, 0)

	for _, submission := range ans {
		if submission.TaskLink != taskLink {
			continue
		}
		tmp = append(tmp, submission)
	}

	ans = tmp

	return ans, nil
}

func (crawler CodeforcesCrawler) ParseSubmissionPage(submissionPageLink string) ([]generic.Submission, error) {
	timeLayout := "Jan/02/2006 15:04"
	ans := make([]generic.Submission, 0)

	res, err := generic.HttpClient.Get(submissionPageLink)
	if err != nil {
		return ans, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ans, generic.ErrNon200Response
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return ans, err
	}

	doc.Find("[data-submission-id]").Each(func(i int, selection *goquery.Selection) {
		var success bool
		var submission generic.Submission
		submission.SubmissionLink, success = goquery.NewDocumentFromNode(selection.Children().Get(0)).Children().Attr("href")
		if success == false {
			err = CodeforcesCrawlerErrorExtractingSubmissionLink
			return
		}
		submission.SubmissionLink = "https://codeforces.com" + submission.SubmissionLink
		timeString := strings.TrimSpace(goquery.NewDocumentFromNode(selection.Children().Get(1)).Text())
		timeLocation := time.FixedZone("GMT+3", +3*60*60)
		submission.SubmissionTime, err = time.ParseInLocation(timeLayout, timeString, timeLocation)
		if err != nil {
			return
		}
		submission.Handle = strings.Trim(goquery.NewDocumentFromNode(selection.Children().Get(2)).Text(), "#@\t\n ")
		submission.TaskLink, success = goquery.NewDocumentFromNode(selection.Children().Get(3)).Children().Attr("href")
		if success == false {
			err = CodeforcesCrawlerErrorExtractingTaskLink
			return
		}
		submission.TaskLink = "https://codeforces.com" + submission.TaskLink
		submission.Task = strings.TrimSpace(goquery.NewDocumentFromNode(selection.Children().Get(3)).Children().Text())
		submission.Language = strings.TrimSpace(goquery.NewDocumentFromNode(selection.Children().Get(4)).Text())
		submission.Result = strings.TrimSpace(goquery.NewDocumentFromNode(selection.Children().Get(5)).Text())
		if submission.Result == "Accepted" {
			submission.Status = "COMPLETED"
		} else {
			submission.Status = "ATTEMPTED"
		}
		ans = append(ans, submission)
	})

	return ans, err
}

func (crawler CodeforcesCrawler) ParseSubmissionTable(tableLink string, beginTime time.Time, endTime time.Time) ([]generic.Submission, error) {
	res, err := generic.HttpClient.Get(tableLink)
	ans := make([]generic.Submission, 0)
	if err != nil {
		return ans, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ans, generic.ErrNon200Response
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	left := 1
	right := 1

	doc.Find(".pagination").Each(func(i int, selection *goquery.Selection) {
		if i == 1 {
			pagination := goquery.NewDocumentFromNode(selection.Children().Get(0))
			maxPageLiSelection := goquery.NewDocumentFromNode(pagination.Children().Get(pagination.Children().Length() - 2))
			maxPageSpanSelection := goquery.NewDocumentFromNode(maxPageLiSelection.Children().Get(0))
			ans, success := maxPageSpanSelection.Attr("pageindex")
			if success == false {
				err = CodeforcesCrawlerErrorExtractingLastPage
				return
			}
			right, err = strconv.Atoi(ans)
			if err != nil {
				return
			}
		}
	})

	var l, r int
	l = left - 1
	r = right + 1

	for r-l > 1 {
		mid := int((l + r) / 2)
		var pagesString, midPageString string
		pagesString, err = myUrlJoin(tableLink, "page")
		if err != nil {
			return ans, err
		}
		midPageString, err = myUrlJoin(pagesString, strconv.Itoa(mid))
		if err != nil {
			return ans, err
		}
		var tmp []generic.Submission
		tmp, err = crawler.ParseSubmissionPage(midPageString)
		if err != nil {
			return ans, err
		}
		if len(tmp) == 0 {
			panic("wtf")
		}
		if tmp[len(tmp)-1].SubmissionTime.After(endTime) {
			l = mid
		} else {
			r = mid
		}
	}
	if r > right {
		return ans, nil
	}

	startPage := r

	l = left - 1
	r = right + 1

	for r-l > 1 {
		mid := int((l + r) / 2)
		var pagesString, midPageString string
		pagesString, err = myUrlJoin(tableLink, "page")
		if err != nil {
			return ans, err
		}
		midPageString, err = myUrlJoin(pagesString, strconv.Itoa(mid))
		if err != nil {
			return ans, err
		}
		var tmp []generic.Submission
		tmp, err = crawler.ParseSubmissionPage(midPageString)
		if err != nil {
			return ans, err
		}
		if len(tmp) == 0 {
			panic("wtf")
		}
		if tmp[0].SubmissionTime.Before(beginTime) {
			r = mid
		} else {
			l = mid
		}
	}

	if left == 0 {
		return ans, nil
	}

	endPage := l

	for i := startPage; i <= endPage; i++ {
		var submissions []generic.Submission
		var pagesString, submissionPageString string
		pagesString, err = myUrlJoin(tableLink, "page")
		if err != nil {
			return ans, err
		}
		submissionPageString, err = myUrlJoin(pagesString, strconv.Itoa(i))
		submissions, err = crawler.ParseSubmissionPage(submissionPageString)
		if err != nil {
			return ans, err
		}
		ans = append(ans, submissions...)
	}

	for len(ans) > 0 && ans[0].SubmissionTime.After(endTime) {
		ans = ans[1:]
	}
	for len(ans) > 0 && ans[len(ans)-1].SubmissionTime.Before(beginTime) {
		ans = ans[:len(ans)-1]
	}

	return ans, nil
}

func (crawler CodeforcesCrawler) ParseSubmission(submissionLink string) (generic.Submission, error) {

	timeLayout := "2006-01-02 15:04:05"

	ans := generic.Submission{}

	res, err := generic.HttpClient.Get(submissionLink)
	if err != nil {
		return ans, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ans, generic.ErrNon200Response
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
	ans.Handle = strings.Trim(goquery.NewDocumentFromNode(userData.Children().Get(1)).Text(), "#@\t\n ")

	//getting language
	ans.Language = strings.TrimSpace(goquery.NewDocumentFromNode(infoRow.Children().Get(3)).Text())

	//getting submission time
	timeLocation := time.FixedZone("GMT+3", +3*60*60)

	timeString := strings.TrimSpace(goquery.NewDocumentFromNode(infoRow.Children().Get(7)).Text())
	ans.SubmissionTime, err = time.ParseInLocation(timeLayout, timeString, timeLocation)
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

	return ans, nil
}
