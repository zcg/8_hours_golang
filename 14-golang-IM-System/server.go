package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播channel
	Message chan string
}

// 创建一个Server接口

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// ListenMessager ListenMessage 监听Message广播消息channel的go routine,一旦有消息就发送给全部在线的
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		//将msg发送给所有的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// 链接当前的业务
	// fmt.Println("链接建立成功")
	user := NewUser(conn, this)
	user.Online()
	//监听用户是否活跃
	isLive := make(chan bool)
	//// 用户上线,将用户加入到onlineMap中
	//this.mapLock.Lock()
	//this.OnlineMap[user.Name] = user
	//this.mapLock.Unlock()
	////广播当前用户上线消息
	//this.BoardCast(user, "is online")
	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 1024*4)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err", err)
				return
			}

			//  提取用户的消息 去除\n
			msg := string(buf[:n-1])
			//用户针对msg进行处理
			user.DoMessage(msg)

			//用户发送任意消息,代表当前用户是活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//

		case <-time.After(time.Second * 300):
			//已经超时
			//将当前的user强制关闭
			user.SendMsg("超过300秒未操作\n")
			close(user.C)
			//conn.Close()
			return
			//runtime.Goexit()
		}
	}
}

func (this *Server) Start() {

	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	// accept
	go this.ListenMessager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		go this.Handler(conn)

	}

}
