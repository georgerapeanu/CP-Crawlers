package interfaces

import (
	"github.com/georgerapeanu/CP-Crawlers/generic"
	"time"
	"errors"
)

type GenericCrawler interface {
	GetSubmissions(handle string,
		beginTime time.Time, //the begin time point
		endTime time.Time) ([]generic.Submission,err error) // the end time point
	GetSubmissionsForTask(handle string, //handle of the user
		taskLink string, //link to the task
		beginTime time.Time, //the begin time point
		endTime time.Time) ([]generic.Submission,err error) // the end time point
	ParseSubmission(submissionLink string) (generic.Submission,err error)
}
