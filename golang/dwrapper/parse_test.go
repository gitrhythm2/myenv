package main

import (
	"dwrapper/context"
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	expect := &context.Context{
		Verbose:     false,
		CommandType: context.CMD_OTHER,
		Args:        nil,
	}

	app, _ := parse(nil)
	if !equal(app, expect) {
		t.Fatalf("parse error: [%#v]", app)
	}

	expect.Verbose = true
	app, _ = parse([]string{"-v"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%v]", app)
	}
}

func TestImage(t *testing.T) {
	expect := &context.Context{
		Verbose:     true,
		CommandType: context.CMD_IMAGE,
		Args:        []string{"l", "other"},
	}

	app, _ := parse([]string{"-v", "i", "l", "other"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%v]", app)
	}

	expect.Verbose = false
	app, _ = parse([]string{"i", "l", "other"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%v]", app)
	}

	expect.Args = []string{"la"}
	app, _ = parse([]string{"i", "la"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%v]", app)
	}

	expect.Args = nil
	app, _ = parse([]string{"i"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%#v][%#v]", app, expect)
	}
}

func TestContainer(t *testing.T) {
	expect := &context.Context{
		Verbose:     true,
		CommandType: context.CMD_CONTAINER,
		Args:        []string{"ps"},
	}

	app, _ := parse([]string{"-v", "ps"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%#v][%#v]", app, expect)
	}

	expect.Verbose = false
	app, _ = parse([]string{"ps"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%#v][%#v]", app, expect)
	}

	expect.Args = []string{"rm", "docker"}
	app, _ = parse([]string{"rm", "docker"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%#v][%#v]", app, expect)
	}
}

func TestLogin(t *testing.T) {
	expect := &context.Context{
		Verbose:     true,
		CommandType: context.CMD_LOGIN,
		Args:        []string{"docker"},
	}

	app, _ := parse([]string{"-v", "l", "docker"})
	if !equal(app, expect) {
		t.Fatalf("parse error: [%#v][%#v]", app, expect)
	}
}

func equal(result *context.Context, expect *context.Context) bool {
	if result.Verbose != expect.Verbose ||
		result.CommandType != expect.CommandType ||
		!reflect.DeepEqual(result.Args, expect.Args) {
		return false
	} else {
		return true
	}
}
