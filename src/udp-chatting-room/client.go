package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
	"udp-chatting-room/config"
	"udp-chatting-room/utils"

	"github.com/gizak/termui"
)

var buff = make([]byte, config.MAX_MESSAGE_LENGTH)

// ui相关
var uiOnline = termui.NewList()

// 处理消息，当接收到请求之后，立刻将请求丢给 goroutine 去处理
func HandleMessage(listener *net.UDPConn) {
	for {
		n, _, err := listener.ReadFromUDP(buff)
		utils.HandleError(err)

		if n > 0 {
			msg := utils.ParseMessage(buff, n)

			switch msg[0] {
			case "online-user":
				termui.Clear()

				list := msg[1 : len(msg)-1]

				uiOnline.Items = list

				termui.Render(uiOnline)
			}
		}
	}
}

func main() {
	var host, nickname string

	flag.StringVar(&host, "host", "", "请输入服务器ip地址")
	flag.StringVar(&nickname, "nickname", "", "请输入您的昵称")

	flag.Parse()

	if host == "" {
		fmt.Println("服务器ip地址不能为空")
		os.Exit(0)
	}

	if nickname == "" {
		fmt.Println("昵称不能为空")
		os.Exit(0)
	}

	// 创建一个用于请求的 UDP 地址
	udpAddr, err := net.ResolveUDPAddr("udp4", host+":"+strconv.Itoa(config.SERVER_PORT))
	utils.HandleError(err)

	// 通过拨号的方式连接 UDP
	udpConn, err := net.DialUDP("udp4", nil, udpAddr)
	utils.HandleError(err)

	// 本地也需要监听端口，等待服务器的链接
	// 生成一个随机的端口用于监听，端口需要大于 10000， 10000以内的端口是常用端口不能占用
	newSeed := rand.NewSource(int64(time.Now().Second()))
	newRand := rand.New(newSeed)
	randPort := newRand.Intn(30000) + 10000

	// 创建一个用于监听的 UDP 地址
	udpLocalAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:"+strconv.Itoa(randPort))
	utils.HandleError(err)

	// 本地监听 UDP 连接
	udpListener, err := net.ListenUDP("udp4", udpLocalAddr)
	utils.HandleError(err)

	// 向服务器发送注册请求，包含自己的昵称和用于监听事件的端口
	utils.SendMessage(udpConn, "register", nickname, strconv.Itoa(randPort))

	// 关闭链接
	defer udpConn.Close()
	defer udpListener.Close()

	// 发送心跳包，用于保持服务器的链接
	go utils.KeepAlive(udpConn, nickname)

	// 循环处理消息
	go HandleMessage(udpListener)

	// 图形库
	if err := termui.Init(); err != nil {
		panic(err)
	}

	defer termui.Close()

	uiOnline.Items = append(uiOnline.Items, "正在获取列表...")
	uiOnline.Block.Height = 40
	uiOnline.Block.Width = 20
	uiOnline.Block.Float = termui.AlignRight
	uiOnline.Block.BorderFg = termui.ColorCyan

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/timer/1s", func(e termui.Event) {
		// 获取在线列表
		utils.SendMessage(udpConn, "online-user", nickname, "")
	})

	termui.Render(uiOnline)

	termui.Loop()
}
