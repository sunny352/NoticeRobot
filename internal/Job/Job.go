package Job

import (
	"Server/internal/Inter"
	"strings"
)

type Job struct {
	Group *Group
	Cron  string
	Json  string
}

func NewJob(group *Group, config Inter.IInfo) *Job {
	return &Job{
		Group: group,
		Cron:  config.GetCron(),
		Json:  config.ToJson(),
	}
}

func (job *Job) Run() {
	job.Group.Post(strings.NewReader(job.Json))
}
