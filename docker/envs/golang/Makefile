NAME := app
.DEFAULT_GOAL := r

.PHONY: b
b:
	go build

.PHONY: r
r:
	go run .

.PHONY: t
t:
	go test
#	go test app/hoge

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	@rm app

.PHONY: help
help:
	@echo b: build
	@echo r: run
	@echo t: test
	@echo install: install
	@echo clean:   clean
