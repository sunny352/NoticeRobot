package WeCom

import (
	"Server/internal/Inter"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadConfig(filePath string) *Inter.Config {
	jsonFile := ReadJsonFile(filePath)
	if jsonFile == nil {
		return nil
	}
	config := &Inter.Config{
		URL:      jsonFile.URL,
		InfoList: make([]Inter.IInfo, 0),
	}

	for _, jsonInfo := range jsonFile.InfoList {
		jobInfo := &JobInfo{
			Cron: jsonInfo.Cron,
		}

		info := &MsgInfo{
			MsgType: jsonInfo.MsgType,
		}

		switch jsonInfo.MsgType {
		case "text":
			info.Text = &Text{
				Content:             jsonInfo.Text.Content,
				MentionedList:       jsonInfo.Text.MentionedList,
				MentionedMobileList: jsonInfo.Text.MentionedMobileList,
			}
		case "markdown":
			info.Markdown = ReadMarkdown(jsonInfo.Markdown.Path)
			if info.Markdown == nil {
				continue
			}
		case "image":
			info.Image = ReadImage(jsonInfo.Image.Path)
			if info.Image == nil {
				continue
			}
		}

		jobInfo.Info = info
		config.InfoList = append(config.InfoList, jobInfo)
	}

	return config
}

type JobInfo struct {
	Cron string
	Info *MsgInfo
}

func (info *JobInfo) GetCron() string {
	return info.Cron
}

func (info *JobInfo) ToJson() string {
	jsonBytes, err := json.Marshal(info.Info)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(jsonBytes)
}

type MsgInfo struct {
	MsgType  string    `json:"msgtype"`
	Text     *Text     `json:"text,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
	Image    *Image    `json:"image,omitempty"`
}

type JsonFile struct {
	URL      string      `json:"url"`
	InfoList []*JsonInfo `json:"infoList"`
}

func ReadJsonFile(filePath string) *JsonFile {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}

	var jsonFile *JsonFile
	err = json.Unmarshal(fileBytes, &jsonFile)
	if err != nil {
		log.Println(err)
		return nil
	}

	return jsonFile
}

func WriteJsonFile(filePath string, jsonFile *JsonFile) {
	fileBytes, err := json.MarshalIndent(jsonFile, "", "\t")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		panic(err)
	}
}

type JsonInfo struct {
	Cron     string          `json:"cron"`
	MsgType  string          `json:"msgtype"`
	Text     *Text           `json:"text,omitempty"`
	Markdown *MarkdownConfig `json:"markdown,omitempty"`
	Image    *ImageConfig    `json:"image,omitempty"`
}

func (jsonInfo *JsonInfo) ToJson() string {
	jsonBytes, err := json.Marshal(jsonInfo)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(jsonBytes)
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type MarkdownConfig struct {
	Path string `json:"path"`
}

type Markdown struct {
	Content string `json:"content"`
}

func ReadMarkdown(filePath string) *Markdown {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Markdown{
		Content: string(fileBytes),
	}
}

type ImageConfig struct {
	Path string `json:"path"`
}

type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
}

func ReadImage(filePath string) *Image {
	imageBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Image{
		Base64: base64.StdEncoding.EncodeToString(imageBytes),
		MD5:    fmt.Sprintf("%X", md5.Sum(imageBytes)),
	}
}
