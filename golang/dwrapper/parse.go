package main

import (
	"dwrapper/context"
)

func parse(args []string) (*context.Context, error) {
	var ctx *context.Context = &context.Context{}
	if args[0] == "-v" {
		ctx.Verbose = true
		args = shift(args)
	}

	switch args[0] {
	case "i":
		ctx.CommandType = context.CMD_IMAGE
		args = shift(args)
	case "c":
		ctx.CommandType = context.CMD_COMPOSE
		args = shift(args)
	case "l":
		ctx.CommandType = context.CMD_LOGIN
		args = shift(args)
	default:
		ctx.CommandType = context.CMD_CONTAINER
	}

	if len(args) > 0 {
		ctx.Args = args
	}

	return ctx, nil
}

func shift(arr []string) []string {
	return arr[1:]
}
