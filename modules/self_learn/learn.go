package self_learn

import (
	"database/sql"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

func Learn(client *client.QQClient,msg *message.PrivateMessage)  {
	box:=strings.Split(msg.ToString(),"`")
	println(box)
	if strings.Contains(msg.ToString(),"learn"){
		str1,str2:=box[1],box[3]
		println(str1,str2)
		db,err:=sql.Open("sqlite3","./learn.db")
		if err!=nil{
			logger.WithError(err).Error("open db error")
			return
		}
		defer db.Close()
		_,err=db.Exec("create table if not exists learn(id integer primary key autoincrement,keyword varchar(20),answer varchar(20))")
		if err!=nil{
			logger.WithError(err).Error("create table error")
			return
		}
		// 先查找是否有相同的keyword，如果有，则更新answer
		rows,err:=db.Query("select * from learn where keyword=?",str1)
		if err!=nil{
			logger.WithError(err).Error("select error")
			return
		}
		defer rows.Close()
		if rows.Next(){
			_,err=db.Exec("update learn set answer=? where keyword=?",str2,str1)
			if err!=nil{
				logger.WithError(err).Error("update error")
				return
			}
		}else{
			_,err=db.Exec("insert into learn(keyword,answer) values(?,?)",str1,str2)
			if err!=nil{
				logger.WithError(err).Error("insert error")
				return
			}
		}
		msg1:=message.NewSendingMessage().Append(message.NewText("小澪已经学会啦！"))
		client.SendPrivateMessage(msg.Sender.Uin,msg1)
	}
}