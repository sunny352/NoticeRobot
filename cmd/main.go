package main

import (
	"Server/internal/Inter"
	"Server/internal/Job"
	"Server/internal/Lark"
	"Server/internal/WeCom"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	c := cron.New()
	err := filepath.Walk("./configs", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".json" {
			return nil
		}

		tempPath := strings.TrimSuffix(path, ".json")

		var config *Inter.Config
		switch filepath.Ext(tempPath) {
		case ".lark":
			config = Lark.ReadConfig(path)
			if config == nil {
				return nil
			}
		case ".wecom":
			config = WeCom.ReadConfig(path)
			if config == nil {
				return nil
			}
		default:
			return nil
		}

		fmt.Println("Add group: " + path)

		group := Job.NewGroup(config)
		for _, job := range group.Jobs {
			err = c.AddJob(job.Cron, job)
			if nil != err {
				log.Println(err)
			} else {
				fmt.Println("\tAdd job: " + job.Cron)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("启动定时任务")
	c.Run()
}
