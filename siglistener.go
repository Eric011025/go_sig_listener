package siglistener

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	SignalInterruptListener func() error
	SignalTerminateListener func() error
	SignalKillListener      func() error
	ErrorFunc               func(error)
)

func ExcuteListener() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			err := SignalInterruptListener()
			if err != nil {
				ErrorFunc(err)
			}
			os.Exit(0)

		case syscall.SIGTERM:
			err := SignalTerminateListener()
			if err != nil {
				ErrorFunc(err)
			}
			os.Exit(0)

		case syscall.SIGKILL:
			err := SignalKillListener()
			if err != nil {
				ErrorFunc(err)
			}
			os.Exit(0)
		}
	}()
}
