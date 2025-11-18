package main

import (
	"fmt"
	"ifpb/remotelist/pkg"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Print("dialing:", err)
	}

	// Synchronous call
	var reply bool
	var reply_i int
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 1, Value: 10}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 1, Value: 20}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 2, Value: 30}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 3, Value: 40}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 3, Value: 50}, &reply)

	fmt.Println("Tamanho da lista 1: ")
	err = client.Call("RemoteList.Get", pkg.GetArgs{ListId: 1, Index: 1}, &reply_i)
	fmt.Println(reply_i)

	fmt.Println()
	fmt.Println("Tamanho da lista 2: ")
	err = client.Call("RemoteList.Size", 3, &reply_i)
	fmt.Println(reply_i)

	err = client.Call("RemoteList.Remove", 0, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}
	
	err = client.Call("RemoteList.Remove", 0, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}
}
