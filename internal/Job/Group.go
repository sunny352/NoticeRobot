package Job

import (
	"Server/internal/Inter"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Group struct {
	URL string

	Jobs []*Job
}

func NewGroup(config *Inter.Config) *Group {
	group := &Group{
		URL:  config.URL,
		Jobs: make([]*Job, 0),
	}

	for _, info := range config.InfoList {
		job := NewJob(group, info)
		group.Jobs = append(group.Jobs, job)
	}

	return group
}

func (group *Group) Post(reader io.Reader) {
	response, err := http.Post(group.URL, "application/json", reader)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response)
	}
}
