package tickledReply

import (
	"github.com/eatmoreapple/openwechat"
)

func IsTickledByFriend(msg *openwechat.Message) {
	if msg.IsTickledMe() {
		_, err := msg.ReplyText("干嘛拍我？")
		if err != nil {
			return
		}
	}
}
