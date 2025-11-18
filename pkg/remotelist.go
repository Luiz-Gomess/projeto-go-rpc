package pkg

import (
	"fmt"
	"sync"
)

type RemoteList struct {
	mu       sync.Mutex
	listsMap map[int][]int
}

type AppendArgs struct {
	ListId int
	Value  int
}

type GetArgs struct {
	ListId int
	Index  int
}

func NewRemoteList() *RemoteList {
	return &RemoteList{
		listsMap: make(map[int][]int),
	}
}

func (rl *RemoteList) getList(listId int) ([]int, error) {
	list, exists := rl.listsMap[listId]

	if !exists {
		return nil, fmt.Errorf("list %d does not exists", listId)
	}

	return list, nil
}

func (rl *RemoteList) Append(args AppendArgs, reply *bool) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.listsMap[args.ListId] = append(rl.listsMap[args.ListId], args.Value)
	*reply = true

	RegisterLog("Append", args.ListId, args.Value)
	return nil
}

func (rl *RemoteList) Get(args GetArgs, reply_i *int) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	list, _ := rl.getList(args.ListId)

	if args.Index < 0 || args.Index >= len(list) {
		return fmt.Errorf("index out of bounds: %d", args.ListId)
	}

	*reply_i = list[args.Index]
	return nil
}

func (rl *RemoteList) Remove(listId int, reply_i *int) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	list, _ := rl.getList(listId)

	if len(list) == 0 {
		return fmt.Errorf("list %d is empty", listId)
	}

	lastIndex := len(list) - 1
	val := list[lastIndex]

	rl.listsMap[listId] = list[:lastIndex]
	*reply_i = val

	RegisterLog("Remove", listId, " ")
	return nil
}

func (rl *RemoteList) Size(listId int, reply_i *int) error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	list, _ := rl.getList(listId)

	*reply_i = len(list) 
	return nil

}
