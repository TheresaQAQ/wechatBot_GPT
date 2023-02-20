package api

import (
	"awesomeProject/config"
	redis2 "awesomeProject/redis"
	"awesomeProject/utils/constant"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"io"
	"net/http"
	"strings"
)

type Data struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int32  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     string `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Payload struct {
	Prompt           string    `json:"prompt"`
	MaxTokens        int64     `json:"max_tokens"`
	Temperature      float64   `json:"temperature"`
	TopP             int64     `json:"top_p"`
	FrequencyPenalty float32   `json:"frequency_penalty"`
	PresencePenalty  float32   `json:"presence_penalty"`
	Model            string    `json:"model"`
	Stop             [2]string `json:"stop"`
}

func parseStr(str string) string {
	replace := strings.Replace(str, "\\n", "", -1)
	replace = strings.Replace(replace, "\n", "", -1)

	return replace
}

func unParseStr(str string) string {
	replace := strings.Replace(str, "<br>", "\n", -1)
	replace = strings.Replace(replace, "Robot:", "", -1)
	replace = strings.Replace(replace, "机器人：", "", -1)
	replace = strings.Replace(replace, "Bot:", "", -1)
	replace = strings.Replace(replace, "吗？", "", -1)
	replace = strings.Replace(replace, "猫娘：", "", -1)

	return replace
}

func getKeyName(message *openwechat.Message) string {
	if message.IsSendByFriend() == true {
		sender, _ := message.Sender()
		return constant.FriendPro + sender.UserName
	}
	if message.IsSendByGroup() == true {
		sender, _ := message.SenderInGroup()
		return constant.GroupPro + sender.UserName
	}
	return ""
}

func Tell(message *openwechat.Message) (string, error) {

	replace := strings.Replace(message.Content, fmt.Sprintf("@%s ", config.Configuration.Bot.BotName), "", -1)

	redis2.WriteRedis(getKeyName(message), replace, constant.Human)
	//读取旧的
	redis, _ := redis2.ReadRedis(getKeyName(message))

	arr := [2]string{"Human:", "AI:"}
	data := Payload{
		Prompt:           redis,
		MaxTokens:        128,
		Temperature:      0.9,
		TopP:             0,
		FrequencyPenalty: 0,
		PresencePenalty:  0.6,
		Model:            "text-davinci-003",
		Stop:             arr,
	}
	payloadBytes, err := json.Marshal(data)

	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", config.Configuration.MsgAPI.OpenAi.Key))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	jsonStr := string(all)

	str := parseStr(jsonStr)

	var response Data
	jsonErr := json.Unmarshal([]byte(str), &response)

	if jsonErr == nil {
		if len(response.Choices) == 0 {
			return "", errors.New("获取数据失败")
		}
		text := response.Choices[0].Text

		//处理一下字符串
		str := unParseStr(text)

		//保存到redis
		redis2.WriteRedis(getKeyName(message), str, constant.Robot)

		return str, nil
	}

	return "", err
}
