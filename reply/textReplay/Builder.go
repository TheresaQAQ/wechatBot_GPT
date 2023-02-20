package textReplay

import (
	"awesomeProject/config"
	"awesomeProject/redis"
	"awesomeProject/utils/constant"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

var mod int = TIANXING

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

func Builder(msg *openwechat.Message) {
	//群聊内发送的消息
	var name string
	if msg.IsSendByGroup() {
		group, _ := msg.SenderInGroup()
		name = group.NickName

	}
	//好友私聊
	if msg.IsSendByFriend() {
		sender, _ := msg.Sender()
		name = sender.NickName
	}

	replace := strings.Replace(msg.Content, fmt.Sprintf("@%s ", config.Configuration.Bot.BotName), "", -1)
	if replace == "切换模式" {
		//如果来自我，而且命令是改变模式
		if name == config.Configuration.Bot.AdminName || name == config.Configuration.Bot.AdminRemarkName {
			//切换模式
			if mod >= 2 {
				mod = 1
				msg.ReplyText(fmt.Sprintf("切换到模式%d", mod))
			} else {
				mod += 1
				msg.ReplyText(fmt.Sprintf("切换到模式%d", mod))
			}
		} else {
			msg.ReplyText("你无权操作")
		}
		return
	}

	if replace == "删除缓存" {
		redis.DelRedis(getKeyName(msg))

		msg.ReplyText("好的！")
		return
	}

	if msg.IsSendByGroup() {
		if msg.IsAt() {
			IsAt(msg, mod)
		}
	}

	if msg.IsSendByFriend() {
		SendMess(msg, mod)
	}

}
