package main

import (
	"fmt"
	"ifpb/remotelist/pkg"
	"net"
	"net/rpc"

	"github.com/carlescere/scheduler"
)

func main() {

	list := pkg.NewRemoteList()
	rpcs := rpc.NewServer()
	rpcs.Register(list)

	l, e := net.Listen("tcp", "[localhost]:5000")
	defer l.Close()

	if e != nil {
		fmt.Println("listen error:", e)
	}

	pkg.LoadData(list)

	_, err := scheduler.Every(1).Minutes().Run(func() {
		pkg.Snapshot(list)
	})

	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		} else {
			break
		}
	}
}
