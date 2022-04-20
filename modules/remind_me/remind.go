package remind_me

import (
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gopkg.in/yaml.v2"
)
var logger = utils.GetModuleLogger("t0hka.autoreply")

func Reminder(client *client.QQClient,msg *message.PrivateMessage){
	//实现一个清单，记录作业情况
	// 支持主动提醒(定时通知)和被动提醒(自己询问)
	println("I come here!")
	if msg.ToString()=="作业" || msg.ToString()=="homework"{
		reply:=query()
		client.SendPrivateMessage(msg.Sender.Uin,message.NewSendingMessage().Append(message.NewText(reply)))
	}
}

func query() string{
	var tem map[string]string
	var sendString string
	path := config.GlobalConfig.GetString("t0hka.autoreply.path")

	if path == "" {
		path = "./homework.yaml"
	}

	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		logger.WithError(err).Errorf("unable to read config file in %s", path)
	}
	sendString+="小澪的作业提醒！\n"
	for key,value :=range tem{
		sendString+=key+":"+value+"\n"
	}
	return sendString
}