package config

type Config struct {
	Bot struct {
		AdminName       string `yaml:"adminName"`
		AdminRemarkName string `yaml:"adminRemarkName"`
		BotName         string `yaml:"botName"`
	} `yaml:"bot"`
	MsgAPI struct {
		OpenAi struct {
			Key string `yaml:"key"`
		} `yaml:"openAi"`
		TianxinAi struct {
			Key string `yaml:"key"`
		} `yaml:"tianxinAi"`
	} `yaml:"msgApi"`
}

var Configuration *Config
