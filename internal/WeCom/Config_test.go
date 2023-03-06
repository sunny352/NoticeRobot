package WeCom_test

import (
	"Server/internal/WeCom"
	"testing"
)

func TestWriteJsonFile(t *testing.T) {
	configs := []*WeCom.JsonInfo{
		{
			Cron:    "0 0 11,15,17 * * 1-5",
			MsgType: "text",
			Text: &WeCom.Text{
				Content:             "今天你喝水了吗？",
				MentionedList:       []string{"@all"},
				MentionedMobileList: nil,
			},
		},
		{
			Cron:    "0 0 11,15,17 * * 1-5",
			MsgType: "markdown",
			Markdown: &WeCom.MarkdownConfig{
				Path: "./res/drink.md",
			},
		},
		{
			Cron:    "0 0 10,12,14,16,18 * * 1-5",
			MsgType: "text",
			Text: &WeCom.Text{
				Content:             "今天你提肛了吗？",
				MentionedList:       []string{"@all"},
				MentionedMobileList: nil,
			},
		},
		{
			Cron:    "0 0 10,12,14,16,18 * * 1-5",
			MsgType: "image",
			Image: &WeCom.ImageConfig{
				Path: "./res/ti.png",
			},
		},
	}

	WeCom.WriteJsonFile("../../configs/config.wecom.json", &WeCom.JsonFile{
		URL:      "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=1b929807-9803-4998-a1f2-12da611f3165",
		InfoList: configs,
	})
}

func TestReadConfig(t *testing.T) {
	config := WeCom.ReadConfig("../../configs/config.wecom.json")
	t.Log(config)
}

func TestText(t *testing.T) {

}

func TestMarkdown(t *testing.T) {
	markdown := WeCom.ReadMarkdown("../../res/drink.md")
	t.Log(markdown)
}

func TestImage(t *testing.T) {
	image := WeCom.ReadImage("../../res/ti.png")
	t.Log(image)
}
