// @Title
// @Description
// @Author
// @Update

package queue

import (
	"fmt"
	"math"
	"testing"

	"github.com/jaswdr/faker"
	log "github.com/sirupsen/logrus"
)

func TestQueueOrder(t *testing.T) {
	fake := faker.New()
	queue, _ := NewQueue(uint(fake.UIntBetween(2, 10)))
	var expected []uint
	for i := 0; i < int(queue.MaxCapacity); i++ {
		expected = append(expected, fake.UInt())
	}

	for _, elem := range expected {
		err := queue.Enqueue(elem)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	for _, elem := range expected {
		val, err := queue.Dequeue()
		if err != nil {
			t.Fatalf(err.Error())
		}

		if val != elem {
			t.Fatalf("Wrong insertion")
		}
	}
}

func TestMaximumMemoryExcedded(t *testing.T) {
	_, err := NewQueue(uint(math.Pow(2, 30)))
	if err == nil {
		t.Fatalf("Maximun memory exceeded error should raised")
	}
}

func TestMaxCapacityExcedded(t *testing.T) {
	fake := faker.New()
	queue, err := NewQueue(uint(fake.UIntBetween(2, 10)))
	if err != nil {
		t.Fatalf(err.Error())
	}

	for i := 0; i < int(queue.MaxCapacity); i++ {
		err := queue.Enqueue(fake.UInt())
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	err = queue.Enqueue(fake.UInt())
	if err == nil {
		t.Fatalf("No max capacity excedde error has been raised")
	}
}

func TestDequeueToEmptyQueue(t *testing.T) {
	fake := faker.New()
	queue, err := NewQueue(uint(fake.UIntBetween(2, 10)))
	if err != nil {
		t.Fatalf(err.Error())
	}

	for i := 0; i < int(queue.MaxCapacity); i++ {
		err := queue.Enqueue(fake.UInt())
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	numOfElements := queue.NumOfElements
	for i := 0; i < int(numOfElements); i++ {
		_, err := queue.Dequeue()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Fatalf("No empty queue error has been raised")
	}

}

func TestQueueConcurrency(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	defer close(ch1)
	defer close(ch2)
	fake := faker.New()
	testNumber := fake.UInt()
	queue, err := NewQueue(uint(fake.UIntBetween(2, 10)))
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = queue.Enqueue(fake.UInt())
	if err != nil {
		t.Fatalf(err.Error())
	}

	go func(queue *Queue, testNumber uint, ch chan interface{}) {
		dequeueFlag := false
		log.Info("Inside first goroutine")
		if queue.IsEmpty() {
			dequeueFlag = true
		}
		queue.Enqueue(testNumber)
		if dequeueFlag && queue.NumOfElements != 1 {
			ch <- fmt.Errorf("Second goroutine arrived first, but the queue number of elements is wrong")
		}

		if !dequeueFlag && queue.NumOfElements != 2 {
			ch <- fmt.Errorf("The First go routuine arrived first, but the queue number of elements is wrong")
		}

		ch <- true
	}(queue, testNumber, ch1)

	go func(queue *Queue, testNumber uint, ch chan interface{}) {
		enqueueFlag := false
		log.Info("Inside second goroutine")
		if queue.NumOfElements == 2 {
			enqueueFlag = true
		}
		queue.Dequeue()
		if enqueueFlag && queue.NumOfElements != 1 {
			ch <- fmt.Errorf("First goroutine arrived first, but the queue number of elements is wrong")
		}

		if !queue.IsEmpty() {
			ch <- fmt.Errorf("Second go routuine arrived first, but the queue is not empty")
		}

		ch <- true
	}(queue, testNumber, ch2)

	canBreak := false
	select {
	case res := <-ch1:
		if canBreak {
			break
		}
		canBreak = true
		handleResponse(res, t)
	case res := <-ch2:
		if canBreak {
			break
		}
		canBreak = true
		handleResponse(res, t)
	}

}

func handleResponse(res interface{}, t *testing.T) {

	switch res.(type) {
	case error:
		err, _ := res.(error)
		t.Fatalf(err.Error())
	default:
		return
	}
}
