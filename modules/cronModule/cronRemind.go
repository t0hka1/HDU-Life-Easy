package cronModule

import (
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/message"
	"gopkg.in/yaml.v2"
)


var logger = utils.GetModuleLogger("t0hka.cronModule")

func Remind(b *bot.Bot)  {
	b.SendPrivateMessage(1263183073, message.NewSendingMessage().Append(message.NewText(query())))
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
