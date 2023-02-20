package api

import (
	"awesomeProject/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TxData struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		DataType string `json:"datatype"`
		Reply    string `json:"reply"`
	} `json:"result"`
}

func TellToTx(msg *openwechat.Message) (string, error) {
	replace := strings.Replace(msg.Content, fmt.Sprintf("@%s ", config.Configuration.Bot.BotName), "", -1)

	println(replace)
	println(msg.Content)

	postValue := url.Values{"key": {config.Configuration.MsgAPI.TianxinAi.Key}, "question": {replace}, "uniqueid": {msg.FromUserName}}
	res, _ := http.PostForm("https://apis.tianapi.com/robot/index", postValue)
	defer res.Body.Close()
	jsonStr, _ := ioutil.ReadAll(res.Body)

	print(string(jsonStr))
	var txData TxData
	json.Unmarshal(jsonStr, &txData)

	if txData.Code == 200 {
		reply := txData.Result.Reply
		replace := strings.Replace(reply, "<br>", "\n", -1)
		return replace, nil
	}
	return "", errors.New(txData.Msg)
}
