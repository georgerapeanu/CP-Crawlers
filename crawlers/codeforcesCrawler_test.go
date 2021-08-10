package crawlers

import "testing"
import "time"
import "github.com/georgerapeanu/CP-Crawlers/generic"

func TestCodeforcesCrawlerParseSubmission1(t *testing.T) {
	generic.Init()
	crawler := CodeforcesCrawler{}
	link := "https://codeforces.com/contest/1554/submission/124237552"
	wantTime, timeErr := time.Parse(time.RFC3339, "2021-07-30T10:10:08+03:00")
	if timeErr != nil {
		t.Errorf("critical error in TestCodeforcesCrawlerParseSubmission1, %s\n", timeErr.Error())
		return
	}
	want := generic.Submission{
		Handle:         "georgerapeanu",
		SubmissionLink: link,
		SubmissionTime: wantTime,
		Task:           "E - You",
		TaskLink:       "https://codeforces.com/contest/1554/problem/E",
		Status:         "COMPLETED",
		Result:         "Accepted",
		Language:       "GNU C++11",
	}
	got, err := crawler.ParseSubmission(link)

	if err != nil {
		t.Errorf("received error %s\n", err.Error())
	} else if got.Equal(want) == false {
		t.Errorf("Error CodeforcesCrawlerParseSubmissionTest1, want %s got %s", want.String(), got.String())
	}
}

func TestCodeforcesCrawlerParseSubmission2(t *testing.T) {
	generic.Init()
	crawler := CodeforcesCrawler{}
	link := "https://codeforces.com/contest/1554/submission/125490593"
	wantTime, timeErr := time.Parse(time.RFC3339, "2021-08-10T17:11:33+03:00")
	if timeErr != nil {
		t.Errorf("critical error in TestCodeforcesCrawlerParseSubmission2, %s\n", timeErr.Error())
		return
	}
	want := generic.Submission{
		Handle:         "georgerapeanu",
		SubmissionLink: link,
		SubmissionTime: wantTime,
		Task:           "E - You",
		TaskLink:       "https://codeforces.com/contest/1554/problem/E",
		Status:         "COMPLETED",
		Result:         "Accepted",
		Language:       "GNU C++14",
	}
	got, err := crawler.ParseSubmission(link)

	if err != nil {
		t.Errorf("received error %s\n", err.Error())
	} else if got.Equal(want) == false {
		t.Errorf("Error CodeforcesCrawlerParseSubmissionTest2, want %s got %s", want.String(), got.String())
	}
}

func TestCodeforcesCrawlerParseSubmissionRedirect1(t *testing.T) {
	generic.Init()
	crawler := CodeforcesCrawler{}
	link := "https://codeforces.com/contest/1554/submission/125490593124121244"
	_, err := crawler.ParseSubmission(link)

	if err != generic.ErrNon200Response {
		t.Errorf("Critical test error TestCodeforcesCrawlerParseSubmissionRedirect1")
	}
}
