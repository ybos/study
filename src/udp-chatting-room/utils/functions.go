package utils

import (
	"fmt"
	"net"
	"os"
	"time"
)

/**
 * 用于处理错误信息
 * 如果遇到错误，则直接中断并抛出异常
 */
func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

/**
 * 消息解析
 */
func ParseMessage(message []byte, len int) []string {
	res := make([]string, 0)
	strNow := ""

	for i := 0; i < len; i++ {
		if string(message[i:i+1]) == ":" {
			res = append(res, strNow)
			strNow = ""
		} else {
			strNow += string(message[i : i+1])
		}
	}

	res = append(res, strNow)
	return res
}

/**
 * 发送消息
 */
func SendMessage(conn *net.UDPConn, tag, nickname, message string) {
	msg := tag

	if nickname != "" {
		msg += ":" + nickname
	}

	if message != "" {
		msg += ":" + message
	}

	conn.Write([]byte(msg))
}

/**
 * 保持链接状态
 * 向指定的链接发送心跳包
 */
func KeepAlive(conn *net.UDPConn, nickname string) {
	for {
		// 每间隔1s向服务器发送一次在线信息
		SendMessage(conn, "heart-beat", nickname, "")

		sleepTimer := time.NewTimer(time.Second)
		<-sleepTimer.C
	}
}
