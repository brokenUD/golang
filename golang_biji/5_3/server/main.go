package main

import (
	// "chatbox/zdbmodel"
	// "chatbox/zlogs"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 直播间信息
type LiveRoom struct {
	user map[int]*UserInfo //直播间用户
	ch   chan []byte       //发送弹幕
	msg  [][]byte          //历史弹幕
}

// 用户信息
type UserInfo struct {
	id       int             //用户Id
	name     string          //用户名称
	conn     *websocket.Conn //当前用户的连接
	connType int             //用户类型 1-游客 2-会员
}

// 公共函数
func commAtoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return 0
}

var list map[int]*LiveRoom
var roomUser map[int]*UserInfo
var roomCh chan []byte
var roomMsg [][]byte

func svrConnHandler(conn *websocket.Conn) {
	//TODO : 表单数据处理
	r := conn.Request()
	r.ParseForm()
	roomId := r.Form["roomId"][0]
	userId := r.Form["userId"][0]
	userName := r.Form["userName"][0]
	connType := r.Form["connType"][0]
	rId := commAtoi(roomId)
	uId := commAtoi(userId)
	cType := commAtoi(connType)

	//TODO : 构建结构体
	uInfo := &UserInfo{
		id:       uId,
		name:     userName,
		conn:     conn,
		connType: cType,
	}
	_, ok := list[rId]

	if !ok {
		roomUser = make(map[int]*UserInfo, 100)
		roomCh = make(chan []byte, 100)
		roomMsg = make([][]byte, 100)
	}

	roomUser[uId] = uInfo
	list[rId] = &LiveRoom{
		user: roomUser,
		ch:   roomCh,
		msg:  roomMsg,
	}

	//TODO : 接收客户端消息
	go func() {
		var content [1024]byte
		for {
			connUserInfo := list[rId].user[uId]
			n, err := connUserInfo.conn.Read(content[:])
			if err != nil {
				list[rId].user[uId].conn.Close()
				delete(list[rId].user, uId)
				break
			}
			if n > 0 {
				if string(content[:n]) == string("ping www.51websocket.cn") {
					content = [1024]byte{}
					continue
				}
				content2 := append([]byte(fmt.Sprint(userId, ":::", userName, ":::")), content[:n]...)
				fmt.Println(content2)
				list[rId].ch <- content2
				content = [1024]byte{}
			}
		}
	}()

	//TODO : 向客户端发送消息
	if !ok {
		for {
			select {
			case content := <-list[rId].ch:
				message := strings.Split(string(content), ":::")
				list[rId].msg = append(list[rId].msg, []byte("<span class='name'>"+message[1]+"："+"</span>"+message[2]))
				for _, v := range list[rId].user {
					v.conn.Write([]byte("<span class='name'>" + message[1] + "：" + "</span>" + message[2]))
				}
			}
		} //<--向客户端发送消息-->
	} else {
		for {
			time.Sleep(time.Second * 60 * 60)
		}
	}
}

func main() {
	list = make(map[int]*LiveRoom, 10)
	// zlogs.InitLog()
	// zdbmodel.InitDb()
	http.Handle("/message", websocket.Handler(svrConnHandler))
	// err := http.ListenAndServe(":8095", nil)
	http.ListenAndServe(":8095", nil)
	// zlogs.ErrorLog(err)
}
