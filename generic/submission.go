package generic

import (
	"fmt"
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

func (this Submission) String() string {
	return fmt.Sprintf(
		` 
		Handle: %s	
		Language: %s
		Result: %s
		SubmissionLink: %s
		SubmissionTime: %s
		Satus: %s
		Task: %s
		TaskLink: %s
		`,
		this.Handle,
		this.Language,
		this.Result,
		this.SubmissionLink,
		this.SubmissionTime.String(),
		this.Status,
		this.Task,
		this.TaskLink)
}

func (this Submission) Equal(other Submission) bool {
	return this.Handle == other.Handle &&
		this.Language == other.Language &&
		this.Result == other.Result &&
		this.SubmissionLink == other.SubmissionLink &&
		this.SubmissionTime.Equal(other.SubmissionTime) &&
		this.Status == other.Status &&
		this.Task == other.Task &&
		this.TaskLink == other.TaskLink
}
