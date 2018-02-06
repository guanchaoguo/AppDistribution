package main

import(
	 "game_frame/multiplexer"
	"net"
	"github.com/pkg/errors"
	"fmt"
)

func server()error{
	multiplexer.NewConnection()
}

const Port  = ":8086"
func main(){

}


func Listen()error{
	var err error
	listener, err := net.Listen("tcp", Port)
	if err != nil{
		return errors.Wrap(err , "TCP服务无法监听在端口"+Port)
	}
	fmt.Println(" 服务监听成功：",listener.Addr().String())
	for{
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("心请求监听失败!")
			continue
		}
		// 开始处理新链接数据
		go handleMessage(conn)
	}

}