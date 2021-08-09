package interfaces

import (
	"github.com/georgerapeanu/CP-Crawlers/generic"
	"time"
	"errors"
)

type GenericCrawler interface {
	GetSubmissions(handle string, //handle of the user
		taskLink string, //link to the task
		beginningTime time.Time) ([]generic.Submission,err error) //the time point from which we want to extract data
	ParseSubmission(submissionLink string) (generic.Submission,err error)
}
