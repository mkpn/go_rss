package controller

import (
	"net/http"
	"github.com/reactivex/rxgo/observable"
	"fmt"
	"github.com/reactivex/rxgo/observer"
)

func RxHandler(w http.ResponseWriter, r *http.Request) {

	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			fmt.Println(item.(string))
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

	emitter := make(chan interface{})
	source := observable.Observable(emitter)
	sub := source.Subscribe(watcher)

	emitter <- "aiueo"
	close(emitter)
	// wait for the channel to emit a Subscription
	<-sub
}
