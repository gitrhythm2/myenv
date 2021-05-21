package main

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"dwrapper/context"
	//"local.packages/myenv/file"
)

func main() {
	ctx, err := parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = ctx.BuildCommand()
	if err != nil {
		log.Fatal(err)
	}

	// -vが指定された場合はコマンドを実行せずに表示のみとする
	// -vが指定されていない場合は、実行するコマンドをターミナルに表示した上でコマンドを実行する
	Verbose(ctx)
	if !ctx.Verbose {
		Exec(ctx)
	}
}

func Verbose(ctx *context.Context) {
	commands := strings.Join(ctx.Command, " ")
	color.Cyan(":%s\n", commands)
}
