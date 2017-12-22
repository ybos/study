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
var uiHistory = termui.NewList()
var uiInput = termui.NewPar("")

// 处理消息，当接收到请求之后，立刻将请求丢给 goroutine 去处理
func HandleMessage(listener *net.UDPConn, nickname string) {
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
			case "message":
				termui.Clear()

				uiHistory.Items = append(uiHistory.Items, msg[1])
			}

			termui.Render(uiOnline, uiHistory, uiInput)
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
	go HandleMessage(udpListener, nickname)

	// 图形库
	if err := termui.Init(); err != nil {
		panic(err)
	}

	defer termui.Close()

	termui.Body.Width = 120

	uiOnline.Items = append(uiOnline.Items, "正在获取列表...")
	uiOnline.Block.BorderLabel = "当前在线用户"
	uiOnline.Block.Height = 40
	uiOnline.Block.Width = 20
	uiOnline.Block.Y = 0
	uiOnline.Block.Float = termui.AlignRight
	uiOnline.Block.BorderFg = termui.ColorCyan

	uiHistory.Block.BorderLabel = "大家正在说:"
	uiHistory.Block.Height = 40
	uiHistory.Block.Width = 100
	uiOnline.Block.Y = 0
	uiHistory.Block.Float = termui.AlignLeft
	uiHistory.Block.BorderFg = termui.ColorCyan

	uiInput.TextFgColor = termui.ColorBlack
	uiInput.Height = 10
	uiInput.Width = 120
	uiInput.Y = 40
	uiInput.BorderLabel = "输入框(回车发送信息)"

	termui.Handle("/sys/kbd", func(e termui.Event) {
		c := e.Data.(termui.EvtKbd)

		b := []rune(c.KeyStr)
		bl := len(b)

		if bl == 1 {
			uiInput.Text += c.KeyStr
		} else if c.KeyStr == "<space>" {
			uiInput.Text += " "
		} else if c.KeyStr == "C-8" && uiInput.Text != "" {
			t := []rune(uiInput.Text)
			uiInput.Text = string(t[:len(t)-1])
		} else if bl == 3 && b[0] != 67 && b[1] != 45 {
			uiInput.Text += c.KeyStr
		}

		termui.Render(uiOnline, uiHistory, uiInput)
	})

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		if uiInput.Text != "" {
			utils.SendMessage(udpConn, "message", nickname, uiInput.Text)

			uiInput.Text = ""

			termui.Render(uiOnline, uiHistory, uiInput)
		}
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/timer/1s", func(e termui.Event) {
		// 获取在线列表
		utils.SendMessage(udpConn, "online-user", nickname, "")
	})

	termui.Render(uiOnline, uiHistory, uiInput)

	termui.Loop()
}
