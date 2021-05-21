package main

import (
	"dwrapper/context"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func Exec(ctx *context.Context) {
	path, err := exec.LookPath(context.APP_NAME)
	if err != nil {
		log.Fatal(err)
	}
	env := os.Environ()

	err = syscall.Exec(path, ctx.Command, env)
	if err != nil {
		log.Fatal(err)
	}
}
