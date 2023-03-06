package Inter

type IInfo interface {
	GetCron() string
	ToJson() string
}

type Config struct {
	URL      string  `json:"url"`
	InfoList []IInfo `json:"infoList"`
}
