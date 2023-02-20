package textReplay

import (
	"awesomeProject/api"
	"awesomeProject/utils"
	"github.com/eatmoreapple/openwechat"
)

func IsAt(msg *openwechat.Message, apiType int) {
	sender, _ := msg.SenderInGroup()
	name := sender.NickName
	var (
		tell string
		err  error
	)

	if apiType == CHAT_GPT {
		tell, err = api.Tell(msg)
	}
	if apiType == TIANXING {
		tell, err = api.TellToTx(msg)
	}
	if err != nil {
		_, err := msg.ReplyText("错误")
		if err != nil {
			return
		}
	}
	if err == nil {
		if tell == "" {
			msg.ReplyText(utils.ParseMessPrefix(name) + "好像出错了")
			return
		}
		_, err := msg.ReplyText(utils.ParseMessPrefix(name) + tell)
		if err != nil {
			return
		}
	}
}
