package goworker

import (
	"fmt"

	"github.com/benmanns/goworker"
)

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    100,
		Queues:         []string{"myqueue", "delimited", "queues"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "resque:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)

	goworker.Register("MyClass", myFunc1)
	goworker.Register("MyClass2", myFunc2)
}

func myFunc1(queue string, args ...interface{}) error {
	fmt.Println("this is the myFunc1 queue(%s),args(%v)", queue, args)
	return nil
}

func myFunc2(queue string, args ...interface{}) error {
	fmt.Println("this is the myFunc1 queue(%s),args(%v)", queue, args)
	return nil
}

func main() {
	goworker.Enqueue(&goworker.Job{
		Queue: "myqueue",
		Payload: goworker.Payload{
			Class: "MyClass",
			Args:  []interface{}{"hi", "there"},
		},
	})

	goworker.Enqueue(&goworker.Job{
		Queue: "myqueue",
		Payload: goworker.Payload{
			Class: "MyClass2",
			Args:  []interface{}{"hi2", "there2"},
		},
	})

	if err := goworker.Work(); err != nil {
		fmt.Println("err is %s", err.Error())
		return
	}
}
