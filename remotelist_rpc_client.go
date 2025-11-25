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

	// Lista 1
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 1, Value: 10}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 1, Value: 20}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 1, Value: 30}, &reply)

	// Lista 2
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 2, Value: 40}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 2, Value: 50}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 2, Value: 60}, &reply)
	
	// Lista 3 
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 3, Value: 70}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 3, Value: 80}, &reply)
	err = client.Call("RemoteList.Append", pkg.AppendArgs{ListId: 3, Value: 90}, &reply)


	err = client.Call("RemoteList.Get", pkg.GetArgs{ListId: 1, Index: 1}, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Lista 1 índice 1: ", reply_i)
	}

	
	err = client.Call("RemoteList.Size", 2, &reply_i)
	if err != nil {
		fmt.Print("Erro: ", err)
		} else {
		fmt.Println("Tamanho da lista 2: ", reply_i)
	}

	err = client.Call("RemoteList.Remove", 1, &reply_i)

	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}

	err = client.Call("RemoteList.Remove", 2, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}
}
