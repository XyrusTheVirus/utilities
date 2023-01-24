// @Title
// @Description
// @Author
// @Update

package utilities_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/XyrusTheVirus/utilities"
	"github.com/jaswdr/faker"
	log "github.com/sirupsen/logrus"
)

func TestStackOrder(t *testing.T) {
	fake := faker.New()
	stack, _ := utilities.NewStack(uint(fake.UIntBetween(2, 10)))
	var expected []uint
	for i := 0; i < int(stack.MaxCapacity); i++ {
		expected = append(expected, fake.UInt())
	}

	for _, elem := range expected {
		err := stack.Push(elem)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	for i := len(expected) - 1; i >= 0; i-- {
		val, err := stack.Pop()
		if err != nil {
			t.Fatalf(err.Error())
		}

		if val != expected[i] {
			t.Fatalf("Wrong insertion")
		}
	}
}

func TestStackMaximumMemoryExcedded(t *testing.T) {
	_, err := utilities.NewStack(uint(math.Pow(2, 30)))
	if err == nil {
		t.Fatalf("Maximun memory exceeded error should raised")
	}
}

func TestStackMaxCapacityExcedded(t *testing.T) {
	fake := faker.New()
	stack, err := utilities.NewStack(uint(fake.UIntBetween(2, 10)))
	if err != nil {
		t.Fatalf(err.Error())
	}

	for i := 0; i < int(stack.MaxCapacity); i++ {
		err := stack.Push(fake.UInt())
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	err = stack.Push(fake.UInt())
	if err == nil {
		t.Fatalf("No max capacity excedde error has been raised")
	}
}

func TestStackPopEmptyStack(t *testing.T) {
	fake := faker.New()
	stack, err := utilities.NewStack(uint(fake.UIntBetween(2, 10)))
	if err != nil {
		t.Fatalf(err.Error())
	}

	for i := 0; i < int(stack.MaxCapacity); i++ {
		err := stack.Push(fake.UInt())
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	numOfElements := stack.NumOfElements
	for i := 0; i < int(numOfElements); i++ {
		_, err := stack.Pop()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}

	_, err = stack.Pop()
	if err == nil {
		t.Fatalf("No empty stack error has been raised")
	}

}

func TestStackConcurrency(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	defer close(ch1)
	defer close(ch2)
	fake := faker.New()
	testNumber := fake.UInt()
	stack, err := utilities.NewStack(uint(fake.UIntBetween(2, 10)))
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = stack.Push(fake.UInt())
	if err != nil {
		t.Fatalf(err.Error())
	}

	go func(stack *utilities.Stack, testNumber uint, ch chan interface{}) {
		popFlag := false
		log.Info("Inside first goroutine")
		if stack.IsEmpty() {
			popFlag = true
		}
		stack.Push(testNumber)
		if popFlag && stack.NumOfElements != 1 {
			ch <- fmt.Errorf("Second goroutine arrived first, but the stack number of elements is wrong")
		}

		if !popFlag && stack.NumOfElements != 2 {
			ch <- fmt.Errorf("The First go routuine arrived first, but the stack number of elements is wrong")
		}

		ch <- true
	}(stack, testNumber, ch1)

	go func(stack *utilities.Stack, testNumber uint, ch chan interface{}) {
		pushFlag := false
		log.Info("Inside second goroutine")
		if stack.NumOfElements == 2 {
			pushFlag = true
		}
		stack.Pop()
		if pushFlag && stack.NumOfElements != 1 {
			ch <- fmt.Errorf("First goroutine arrived first, but the stack number of elements is wrong")
		}

		if !stack.IsEmpty() {
			ch <- fmt.Errorf("Second go routuine arrived first, but the stack is not empty")
		}

		ch <- true
	}(stack, testNumber, ch2)

	canBreak := false
	for {
		select {
		case res := <-ch1:
			if canBreak {
				return
			}
			canBreak = true
			handleResponse(res, t)
		case res := <-ch2:
			if canBreak {
				return
			}
			canBreak = true
			handleResponse(res, t)
		}
	}

}

func handleStackResponse(res interface{}, t *testing.T) {

	switch res.(type) {
	case error:
		err, _ := res.(error)
		t.Fatalf(err.Error())
	default:
		return
	}
}
