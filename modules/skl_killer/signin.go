package skl_killer

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	//"github.com/Logiase/MiraiGo-Template/server"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
	"fmt"
)
type information struct {
	name  string
	tel   string
	pwd   string
	token string
}


func Signin(qqclient *client.QQClient,msg *message.PrivateMessage)  {
	//返回qq消息
	fmt.Println("T0")
	if msg.Sender.Uin==1263183073||msg.Sender.Uin==1792138307{
		rrr,_:=regexp.MatchString("^([1-9][1-9][1-9][1-9])$",msg.ToString())
		if rrr{
			//从前端拿到签到码
			code,_:=strconv.Atoi(msg.ToString())
			//从flask服务器拿到token
			token:=getToken()
			//向上课啦发送code+token
			reply:=Sendcode(code,token)
			if reply==""{
				reply="签到失败啦！"
			}
			msg1:=message.NewSendingMessage().Append(message.NewText(reply))
			qqclient.SendPrivateMessage(msg.Sender.Uin,msg1)
		}else {
			//msg1:=message.NewSendingMessage().Append(message.NewText("没有签到码的我只会傻笑🤣"))
			//qqclient.SendPrivateMessage(msg.Sender.Uin,msg1)
		}
	}else{
		//msg1:=message.NewSendingMessage().Append(message.NewText("很遗憾，你并非指定之人！"))
		//qqclient.SendPrivateMessage(msg.Sender.Uin,msg1)
	}
}


func Sendcode(code int,token string) string{
	rawUrl :="https://skl.hdu.edu.cn/api/checkIn/code-check-in"
	client :=&http.Client{}
	params := url.Values{}
	Url, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	params.Set("userid", "21273110")
	params.Set("code", strconv.Itoa(code))
	params.Set("latitude", "30.31958")
	params.Set("longitude", "120.3391")
	params.Set("t", strconv.FormatInt(time.Now().UnixMilli(), 10))
	//println(strconv.FormatInt(time.Now().UnixMilli(), 10))
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("User-Agent", "YiBan/5.0.1")
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("skl-ticket", getUuid()) //生成的uuid
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	//考虑把reply返回到前端
	//fmt.Println(string(body),code)
	err = resp.Body.Close()
	if err != nil {
		return  ""
	}else {
		return string(body)
	}
}

func getUuid() string {
	u, _ := uuid.NewV4()
	return u.String()
}

func getToken() string {
	//从flask服务器拿到token,访问127.0.0.1:5000/token
	resp, err := http.Get("http://127.0.0.1:5000/token")
	if err != nil {
		fmt.Println("get token error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read token error:", err)
		return ""
	}
	return string(body)
}
