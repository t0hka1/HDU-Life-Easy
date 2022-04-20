package cronModule

import (
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"

)


var logger = utils.GetModuleLogger("t0hka.cronModule")

func Reminder(client *client.QQClient,msg *message.PrivateMessage){
	//实现一个清单，记录作业情况
	// 支持主动提醒(定时通知)和被动提醒(自己询问)
	c:=cron.New()
	_, err := c.AddFunc("@every 1m", func() {
		reply:=query()
		client.SendPrivateMessage(msg.Sender.Uin,message.NewSendingMessage().Append(message.NewText("cron success！")))
		client.SendPrivateMessage(msg.Sender.Uin,message.NewSendingMessage().Append(message.NewText(reply)))
	})
	if err != nil {
		return
	}
}

func query() string{
	var tem map[string]string
	var sendString string
	path := config.GlobalConfig.GetString("t0hka.cronModule.path")

	if path == "" {
		path = "./homework.yaml"
	}

	bytes := utils.ReadFile(path)
	err := yaml.Unmarshal(bytes, &tem)
	if err != nil {
		logger.WithError(err).Errorf("unable to read config file in %s", path)
	}
	println(tem)
	sendString+="小澪的作业提醒！\n"
	for key,value :=range tem{
		sendString+=key+":"+value+"\n"
	}
	println(sendString)
	return sendString
}
