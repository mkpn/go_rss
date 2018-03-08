package eventbus

import (
	"github.com/reactivex/rxgo/observer"
	"fmt"
)

var successWatcher = &observer.Observer{

	// Register a handler function for every next available item.
	NextHandler: func(item interface{}) {
		fmt.Printf("Processing: %v\n", item)
	},

	// Register a handler for any emitted error.
	ErrHandler: func(err error) {
		fmt.Printf("Encountered error: %v\n", err)
	},

	// Register a handler when a stream is completed.
	DoneHandler: func() {
		fmt.Println("Done!")
	},
}

var errorWatcher = &observer.Observer{}

var watcher = &observer.Observer{
	NextHandler: func(item interface{}) {
		fmt.Println(item)
	},

	// Register a handler for any emitted error.
	ErrHandler: func(err error) {
		fmt.Printf("Encountered error: %v\n", err)
	},

	// Register a handler when a stream is completed.
	DoneHandler: func() {
		fmt.Println("Done!")
	},
}

//GetInstance returns singleton successWatcher
func GetSuccessWatcher() *observer.Observer {
	return successWatcher
}

func GetErrorWatcher() *observer.Observer {
	return errorWatcher
}

func GetWatcher() *observer.Observer {
	return watcher
}
