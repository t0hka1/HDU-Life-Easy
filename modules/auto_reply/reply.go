package auto_reply

import (
	"database/sql"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"strings"
)

func Reply(qqclient *client.QQClient,msg *message.PrivateMessage)  {
	if msg.ToString()!=""{
		if strings.Contains(msg.ToString(),"learn"){
			return
		}
		db,err:=sql.Open("sqlite3","./learn.db")
		if err!=nil{
			logger.WithError(err).Error("open db error")
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				logger.WithError(err).Error("close db error")
			}
		}(db)
		rows,err:=db.Query("select answer from learn where keyword=?",msg.ToString())
		if err!=nil{
			logger.WithError(err).Error("select error")
			return
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				logger.WithError(err).Error("close rows error")
			}
		}(rows)
		for rows.Next(){ //var keyword string
			var answer string
			err=rows.Scan(&answer)
			if err!=nil{
				logger.WithError(err).Error("scan error")
				return
			}
			answer1:=message.NewSendingMessage().Append(message.NewText(answer))
			qqclient.SendPrivateMessage(msg.Sender.Uin,answer1)
		}
	}
}
