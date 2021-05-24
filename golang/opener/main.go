package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	config := LoadConfig(configFilepath())
	item := config.Find(os.Args[1])
	if item == nil {
		return
	}

	fmt.Printf("%s %s", item.Type, item.Uri)
}

func configFilepath() string {
	path := fmt.Sprintf("%s/.config/opener/config.toml", os.Getenv("HOME"))
	return filepath.FromSlash(path)
}
