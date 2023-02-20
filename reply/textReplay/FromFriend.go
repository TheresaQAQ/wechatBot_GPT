package textReplay

import (
	"awesomeProject/api"
	"github.com/eatmoreapple/openwechat"
)

const CHAT_GPT = 1
const TIANXING = 2

func SendMess(msg *openwechat.Message, apiType int) {
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

	if err == nil {
		if tell == "" {
			msg.ReplyText("好像出错了")
			return
		}
		_, err := msg.ReplyText(tell)
		if err != nil {
			return
		}
	}

	if err != nil {
		_, err2 := msg.ReplyText("错误")
		if err2 != nil {
			return
		}
	}
}
