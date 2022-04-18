package auto_reply

import (
	"database/sql"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"strings"
)

func Reply(qqclient *client.QQClient,msg *message.PrivateMessage)  {
	// 查询数据库
	// 查询到了就返回
	if strings.Contains(msg.ToString(),"learn"){
		return
	}
	db,err:=sql.Open("sqlite3","./learn.db")
	if err!=nil{
		logger.WithError(err).Error("open db error")
		return
	}
	defer db.Close()
	rows,err:=db.Query("select answer from learn where keyword=?",msg.ToString())
	if err!=nil{
		logger.WithError(err).Error("select error")
		return
	}
	defer rows.Close()
	for rows.Next(){
		println("I am here")
		//var keyword string
		var answer string
		err=rows.Scan(&answer)
		println(err)
		if err!=nil{
			logger.WithError(err).Error("scan error")
			return
		}
		println(answer)
		answer1:=message.NewSendingMessage().Append(message.NewText(answer))
		qqclient.SendPrivateMessage(msg.Sender.Uin,answer1)
	}
}
