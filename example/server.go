package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"ace"
)

func main() {
	env := new(ace.Environ)
	err := env.SetOpt()
	if err != nil {
		os.Exit(1)
	}

	ctx := ace.NewContext()
	ctx.Setup(env)
	ctx.AddPath(
		ace.Router{
			Method:  "GET",
			Path:    "/test",
			Context: ctx,
		},
		func(trans *ace.TransPort) error {
			trans.HttpWriter.Write([]byte("<h1>404 not found!</h1>"))
			return nil
		},
	)
	ctx.AddPath(
		ace.Router{
			Method:  "POST",
			Path:    "/test2",
			Context: ctx,
		},
		func(trans *ace.TransPort) error {
			trans.HttpWriter.Write([]byte("<h1>500 error!</h1>"))
			return errors.New("internal error")
		},
	)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		switch <-exit {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ctx.Shutdown()
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}
