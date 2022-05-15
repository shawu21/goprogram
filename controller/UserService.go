package controller

import (
	"Program/constants"
	"fmt"
	"net"
)

func ListenRoom() {
	listener, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	connMap := make(map[net.Conn]string)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err=", err)
		} else {
			fmt.Println("Accept success")
		}
		connMap[conn] = conn.RemoteAddr().String()
		go HandlerConn(conn, connMap)
	}
}

func HandlerConn(conn net.Conn, connMap map[net.Conn]string) {
	buf := make([]byte, 1024)

	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err=", err)
			return
		}
		msg := string(buf[:n])
		broadcast(conn, connMap, msg)
	}

}

func broadcast(conn net.Conn, connMap map[net.Conn]string, msg string) {
	for curConn := range connMap {
		_, err := curConn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发送广播失败")
			return
		}
	}
}
