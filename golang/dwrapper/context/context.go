package context

import (
	"fmt"
	"os"

	"local.packages/myenv/file"
)

type Context struct {
	// trueの場合、組み立てたコマンドを表示するのみで実行はしない
	Verbose     bool
	Args        []string
	CommandType string
	Command     []string
	CurrentDir  string
}

const (
	APP_NAME string = "docker"

	// Context.Commandにセットするコマンドのタイプ定義
	CMD_IMAGE     string = "image"
	CMD_CONTAINER string = "container"
	CMD_COMPOSE   string = "compose"
	CMD_LOGIN     string = "login"
	CMD_OTHER     string = ""
)

func (ctx *Context) BuildCommand() error {
	var err error = nil

	switch ctx.CommandType {
	case CMD_IMAGE:
		ctx.Command = buildImageCommand(ctx)
	case CMD_CONTAINER:
		ctx.Command = buildContainerCommand(ctx)
	case CMD_COMPOSE:
		ctx.Command = buildComposeCommand(ctx)
	case CMD_LOGIN:
		ctx.Command = buildLoginCommand(ctx)
	case CMD_OTHER:
		ctx.Command = normalBuild(ctx)
	default:
		err = fmt.Errorf("unknown command[%s]", ctx.Command)
	}

	return err
}

func buildImageCommand(ctx *Context) []string {
	switch {
	case ctx.Args == nil:
		return []string{APP_NAME, ctx.CommandType, "ls", "-a"}
	case ctx.Args[0] == "l":
		return []string{APP_NAME, ctx.CommandType, "ls"}
	default:
		return normalBuild(ctx)
	}
}

func buildContainerCommand(ctx *Context) []string {
	switch {
	case ctx.Args[0] == "psa":
		return []string{APP_NAME, ctx.CommandType, "ps", "-a"}
	default:
		return normalBuild(ctx)
	}
}

func buildComposeCommand(ctx *Context) []string {
	// カレントディレクトリにdocker-compose.ymlが存在しなかった場合は.devcontainer
	// ディレクトリに移動する。
	// compose用のファイルはdocker-compose.ymlで、プロジェクトルートか.devcontainer
	// に存在していることを前提としている
	// その上で、プロジェクトルートでコマンドを発行してもdocker-compose.ymlを読み込める
	// ようにするのが、この処理の意図
	if !file.IsFile("docker-compose.yml") {
		os.Chdir(".devcontainer")
	}
	return normalBuild(ctx)
}

func buildLoginCommand(ctx *Context) []string {
	var containerName string
	if ctx.Args != nil {
		containerName = ctx.Args[0]
	}
	return []string{APP_NAME, "exec", "-it", containerName, "fish", "--login"}
}

func normalBuild(ctx *Context) []string {
	length := 1 + len(ctx.CommandType) + cap(ctx.Args)
	command := make([]string, length)

	command[0] = APP_NAME
	if len(ctx.CommandType) > 0 {
		command[1] = ctx.CommandType
	}
	return append(command, ctx.Args...)
}
