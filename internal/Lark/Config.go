package Lark

import (
	"Server/internal/Inter"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	URL      string      `json:"url"`
	InfoList []*CronInfo `json:"infoList"`
}

func (config *Config) ToInterConfig() *Inter.Config {
	result := &Inter.Config{
		URL: config.URL,
	}

	for _, cronInfo := range config.InfoList {
		result.InfoList = append(result.InfoList, cronInfo)
	}

	return result
}

func ReadConfig(filePath string) *Inter.Config {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}

	var config *Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
		return nil
	}

	return config.ToInterConfig()
}

func WriteConfig(filePath string, config *Inter.Config) {
	jsonBytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		log.Println(err)
	}
}

type CronInfo struct {
	Cron string   `json:"cron"`
	Info *JobInfo `json:"info"`
}

func (cronInfo *CronInfo) GetCron() string {
	return cronInfo.Cron
}

func (cronInfo *CronInfo) ToJson() string {
	jsonBytes, err := json.Marshal(cronInfo.Info)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(jsonBytes)
}

type JobInfo struct {
	MsgType string                 `json:"msg_type"`
	Content map[string]interface{} `json:"content"`
}
