package main

import (
	"awesomeProject/config"
	"awesomeProject/reply/textReplay"
	"awesomeProject/reply/tickledReply"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

type name struct {
	Abc string
}

var ssss *name = new(name)

func main() {
	//加载配置文件
	config.Configuration, _ = config.Loading(config.Configuration)
	println(config.Configuration.MsgAPI.OpenAi.Key)

	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	defer reloadStorage.Close()

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {

		//如果光at但是没有其它信息，不做处理
		replace := strings.Replace(msg.Content, fmt.Sprintf("@%s ", config.Configuration.Bot.BotName), "", -1)
		if replace == "" {
			return
		}

		if msg.IsText() {
			textReplay.Builder(msg)
		}
		if msg.IsTickled() {
			tickledReply.Builder(msg)
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println(err)
		return
	}

	/*	// 获取登陆的用户
		self, err := bot.GetCurrentUser()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 获取所有的好友
		friends, err := self.Friends()
		fmt.Println(friends, err)

		// 获取所有的群组
		groups, err := self.Groups()
		fmt.Println(groups, err)*/

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
