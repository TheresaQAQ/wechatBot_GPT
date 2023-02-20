package tickledReply

import (
	"awesomeProject/utils"
	"github.com/eatmoreapple/openwechat"
)

func IsTickledByGround(msg *openwechat.Message) {
	sender, err := msg.SenderInGroup()
	if err != nil {
		print(err.Error())
		return
	}

	name := sender.String()
	if msg.IsTickledMe() {
		_, err := msg.ReplyText(utils.ParseMessPrefix(name) + "干嘛拍我？")
		if err != nil {
			return
		}
	}
}
