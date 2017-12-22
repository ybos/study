package main

// https://github.com/yinxin630/gochat

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"udp-chatting-room/config"
	"udp-chatting-room/utils"
)

type User struct {
	LastHeartBeatTime time.Time
	Conn              *net.UDPConn
}

var buff = make([]byte, config.MAX_MESSAGE_LENGTH)
var onlineUser = make(map[string]*User)

// 处理消息，当接收到请求之后，立刻将请求丢给 goroutine 去处理
func HandleMessage(listener *net.UDPConn) {
	n, addr, err := listener.ReadFromUDP(buff)
	utils.HandleError(err)

	if n > 0 {
		msg := utils.ParseMessage(buff, n)

		switch msg[0] {
		case "register":
			// 注册用户函数
			if _, ok := onlineUser[msg[1]]; ok {
				// 加入该昵称已经被人使用了
				fmt.Println(msg[1], "重复注册聊天室，本次注册失败")
			} else {
				// 将用户注册进我们的在线名单内
				// 创建一个链接，保存到我们的映射内
				// 创建一个用于请求的 UDP 地址
				udpAddr, err := net.ResolveUDPAddr("udp4", addr.IP.String()+":"+msg[2])
				utils.HandleError(err)

				// 通过拨号的方式连接 UDP
				udpConn, err := net.DialUDP("udp4", nil, udpAddr)
				utils.HandleError(err)

				onlineUser[msg[1]] = &User{time.Now(), udpConn}
			}
		case "heart-beat":
			if _, ok := onlineUser[msg[1]]; ok {
				onlineUser[msg[1]].LastHeartBeatTime = time.Now()
			}
		case "online-user":
			if _, ok := onlineUser[msg[1]]; ok {
				list := ""
				for k, _ := range onlineUser {
					list += k + "\n"
				}

				utils.SendMessage(onlineUser[msg[1]].Conn, "online-user", "", list)
			}
		}

		// for k, v := range onlineUser {
		// fmt.Println(k, v.IP, v.Port, v.LastHeartBeatTime)
		// }
	}
}

/**
 * 清理已经不在线的用户
 */
func cleanDead() {
	for {
		// 超过10秒钟未能发送心跳包的，就认定是死链，需要被清理
		condition := time.Now().Add(-10 * time.Second)

		for k, v := range onlineUser {
			if v.LastHeartBeatTime.Before(condition) {
				v.Conn.Close()

				delete(onlineUser, k)
			}
		}

		fmt.Println("当前在线用户：")
		for k, _ := range onlineUser {
			fmt.Println(k)
		}

		<-time.After(time.Second * 5)
	}
}

func main() {
	// 监听地址和端口, 获得一个可以使用的 UDP 地址
	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:"+strconv.Itoa(config.SERVER_PORT))
	utils.HandleError(err)

	// 监听连接
	udpListener, err := net.ListenUDP("udp4", udpAddr)
	utils.HandleError(err)

	// 当服务结束的时候，关闭监听连接
	defer udpListener.Close()

	fmt.Println("开始监听：")

	go cleanDead()

	// 循环处理消息
	for {
		HandleMessage(udpListener)
	}
}
