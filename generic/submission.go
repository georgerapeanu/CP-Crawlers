package generic

import (
	"time"
)

type Submission struct {
	Handle         string    //the handle of the user that sent the submission
	Language       string    //the language in which the task was solved
	Result         string    //the result(this is what appears on the page)
	SubmissionLink string    //link to the submission
	SubmissionTime time.Time //the time at which the submission was sent
	Status         string    //the status of the task(ATTEMPTED, COMPLETED)
	Task           string    //the task name
	TaskLink       string    //the link to the task
}
