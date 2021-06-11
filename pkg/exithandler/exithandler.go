package exithandler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Init initializes and awaits for exit signals to handle
// program termination gracefully
func Init(cb func()) {
	signals := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		log.Println("exit reason: ", sig)
		terminate <- true
	}()

	<-terminate
	cb()
	log.Println("exiting program")
}
