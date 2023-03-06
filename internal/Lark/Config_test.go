package Lark_test

import (
	"Server/internal/Inter"
	"Server/internal/Lark"
	"testing"
)

func TestWriteJsonFile(t *testing.T) {
	config := Inter.Config{
		URL: "https://open.feishu.cn/open-apis/bot/v2/hook/5b8735eb-e784-40bc-95c7-427ffe4ced99",
		InfoList: []Inter.IInfo{
			&Lark.CronInfo{
				Cron: "0 0 11,15,17 * * 1-5",
				Info: &Lark.JobInfo{
					MsgType: "text",
					Content: map[string]interface{}{
						"text": "今天你喝水了吗？",
					},
				},
			},
			&Lark.CronInfo{
				Cron: "0 0 10,12,14,16,18 * * 1-5",
				Info: &Lark.JobInfo{
					MsgType: "text",
					Content: map[string]interface{}{
						"text": "今天你提肛了吗？",
					},
				},
			},
		},
	}
	Lark.WriteConfig("../../configs/config.lark.json", &config)
}

func TestReadConfig(t *testing.T) {
	config := Lark.ReadConfig("../../configs/config.lark.json")
	t.Log(config)
}
