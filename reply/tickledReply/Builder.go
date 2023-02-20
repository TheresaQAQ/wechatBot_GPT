package tickledReply

import "github.com/eatmoreapple/openwechat"

func Builder(msg *openwechat.Message) {
	if msg.IsSendByGroup() {
		if msg.IsTickledMe() {
			IsTickledByGround(msg)
		}
	}

	if msg.IsSendByFriend() {
		if msg.IsTickledMe() {
			IsTickledByFriend(msg)
		}
	}

}
