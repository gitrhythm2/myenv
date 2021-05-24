package main

import (
	"testing"
)

const CONFIG_PATH = "./config_sample.toml"

func TestLoadConfig(t *testing.T) {
	config := LoadConfig(CONFIG_PATH)
	if len(config.Item) == 0 {
		t.Fatalf("TestLoadConfig error")
	}
}

func TestFindNotFound(t *testing.T) {
	config := LoadConfig(CONFIG_PATH)
	item := config.Find("hoge")
	if item != nil {
		t.Fatalf("find error. want not found. %v", item)
	}
}

func TestFindGoogle(t *testing.T) {
	config := LoadConfig(CONFIG_PATH)
	item := config.Find("ggl")
	if item.Key != "ggl" {
		t.Fatalf("find[key] error %v", item)
	}
	if item.Type != TYPE_WWW {
		t.Fatalf("find[uri] error %v", item)
	}
	if item.Uri != "www.google.co.jp" {
		t.Fatalf("find[uri] error %v", item)
	}
}

func TestFindBashrc(t *testing.T) {
	config := LoadConfig(CONFIG_PATH)
	item := config.Find("bashrc")
	if item.Key != "bashrc" {
		t.Fatalf("find[key] error %v", item)
	}
	if item.Type != TYPE_FILE {
		t.Fatalf("find[uri] error %v", item)
	}
	if item.Uri != "/home/user/.bashrc" {
		t.Fatalf("find[uri] error %v", item)
	}
}
