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

type ErrConnectionClose byte

/* コンストラクタ */
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

/* 準備信号 */
func (conn *CHaserConnect) getReady() ([]byte, error) {
  info := conn.Receive()
  conn.Command("gr")
  info, err := conn.ReceiveInfo()
  return info, err
}

/* ターン終了 */
func (conn *CHaserConnect) TurnEnd() {
  conn.Command("#")
}

func (conn *CHaserConnect) walkRight() []byte {
  conn.Command("wr")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) walkLeft() []byte {
  conn.Command("wl")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) walkUp() []byte {
  conn.Command("wu")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) walkDown() []byte {
  conn.Command("wd")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) searchUp() []byte {
  conn.Command("su")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) searchDown() []byte {
  conn.Command("sd")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) searchRight() []byte {
  conn.Command("sr")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) searchLeft() []byte {
  conn.Command("sl")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) lookUp() []byte {
  conn.Command("lu")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) lookDown() []byte {
  conn.Command("ld")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) lookLeft() []byte {
  conn.Command("ll")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) lookRight() []byte {
  conn.Command("lr")
  info, _ := conn.ReceiveInfo()
  return info
}

func (conn *CHaserConnect) putUp() []byte {
  conn.Command("pu")
  info, _ := conn.ReceiveInfo()
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

func (conn *CHaserConnect) ReceiveInfo() ([]byte, error) {
  result := conn.Receive()
  info := ToIntArray(result[:10])
  if info[0] == 0 {
    return info[1:len(info)], ErrConnectionClose(info[0])
  }
  return info[1:len(info)], nil
}

func (e ErrConnectionClose) Error() string {
  return fmt.Sprintf("Connection Closed")
}

func ToIntArray(info []byte) []byte {
  result := make([]byte, len(info))
  for i, char := range info {
    result[i] = char - '0'
  }
  return result
}

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
  name := "Go"

  connect := NewCHaserConnect(host, port, name)
  fmt.Println(connect.host)
  fmt.Println(connect.port)
  fmt.Println(connect.name)
  for {
    _, err := connect.getReady()
    if err != nil {
      fmt.Println(err)
      connect.Close()
      os.Exit(0)
    }
    connect.searchUp()
    connect.TurnEnd()
  }
}
