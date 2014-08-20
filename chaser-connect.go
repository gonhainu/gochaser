package main

/**
@author Nobuyuki SAKAI
*/

import (
  "fmt"
  "net"
  "os"
)

const (
  RECV_BUF_LEN = 1024
)

type CHaserConnect struct {
  host string
  port string
  name string
  socket *net.TCPConn
}

func NewCHaserConnect(host, port, name string) *CHaserConnect {
  tcp_addr, err := net.ResolveTCPAddr("tcp",  host + ":" + port)
  if err != nil {
    println("error tcp resolve failed", err.Error())
    os.Exit(1)
  }
  tcp_conn, err := net.DialTCP("tcp", nil, tcp_addr)
  if err != nil {
    println("接続できませんでした．", err.Error())
    os.Exit(1)
  }
  Send(tcp_conn, name + "\n")
  conn := &CHaserConnect{host, port, name, tcp_conn}
  return conn
}

func (conn *CHaserConnect) getReady() []byte {
  info := conn.Receive()
  conn.Command("gr")
  info = conn.Receive()
  return info
}

func (conn *CHaserConnect) TurnEnd() {
  conn.Command("#")
}

func (conn *CHaserConnect) walkRight() []byte {
  conn.Command("wr")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) walkLeft() []byte {
  conn.Command("wl")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) walkUp() []byte {
  conn.Command("wu")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) walkDown() []byte {
  conn.Command("wd")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) searchUp() []byte {
  conn.Command("su")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) searchDown() []byte {
  conn.Command("sd")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) searchRight() []byte {
  conn.Command("sr")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) searchLeft() []byte {
  conn.Command("sl")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) lookUp() []byte {
  conn.Command("lu")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) lookDown() []byte {
  conn.Command("ld")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) lookLeft() []byte {
  conn.Command("ll")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) lookRight() []byte {
  conn.Command("lr")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) putUp() []byte {
  conn.Command("pu")
  info := conn.Receive()
  return info
}

func (conn *CHaserConnect) Close() {
  conn.socket.Close()
}

func (conn *CHaserConnect) Command(cmd string) {
  fmt.Printf("send commend [%s]\n", cmd)
  Send(conn.socket, cmd + "\n")
}

func (conn *CHaserConnect) Receive() []byte {
  result := Recv(conn.socket)
  result = result[:10]
  fmt.Printf("data reveived [%s]\n", string(result))
  return result
}

// func (conn *CHaserConnect) ReceiveInfo() []int64 {
//   result := conn.Receive()
//   info :=
// }

func Send(socket *net.TCPConn, msg string) {
  _, err := socket.Write([]byte(msg))
  if err != nil {
    fmt.Println("Error send request:", err.Error())
  }
}

func Recv(socket *net.TCPConn) []byte {
  buf_recever := make([]byte, RECV_BUF_LEN)
  _, err := socket.Read(buf_recever)
  if err != nil {
    fmt.Println("Error while receive response:", err.Error())
    return []byte("")
  }
  return buf_recever
}

func main() {
  host := "localhost"
  port := "2009"
  name := "hoge"

  connect := NewCHaserConnect(host, port, name)
  fmt.Println(connect.host)
  fmt.Println(connect.port)
  fmt.Println(connect.name)
  for {
    connect.getReady()
    connect.searchUp()
    connect.TurnEnd()
  }
  connect.Close()
}
