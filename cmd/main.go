// @Title
// @Description
// @Author
// @Update

package main

import "github.com/XyrusTheVirus/utilities/internal/queue"

func main() {
	queue, _ := queue.NewQueue(100)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
	queue.Dequeue()
	queue.Print()
}
