#! /usr/bin/make
#
# Makefile for goa-react-ts
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#
DBNAME = dev
MYSQL_USER = root
MYSQL_PASSWORD = root
.PHONY: all start depend bootstrap generate models client build run

all: clean generate build

clean:
	@rm -rf ./back/app
	@rm -rf ./back/client
	@rm -rf ./back/tool
	@rm -rf ./back/public/swagger
	@rm -rf ./back/public/schema
	@rm -rf ./back/public/js
	@rm -f server

bootstrap:
	@goagen main    -d github.com/kawabet/goa-react-ts/design -o ./back/controllers

generate:
	@goagen app     -d github.com/kawabet/goa-react-ts/design
	@goagen swagger -d github.com/kawabet/goa-react-ts/design -o ./back/public
	@goagen schema  -d github.com/kawabet/goa-react-ts/design -o ./back/public
	@goagen client  -d github.com/kawabet/goa-react-ts/design
	@goagen js      -d github.com/kawabet/goa-react-ts/design -o ./back/public

build:
	@go build -o server

models:
	@xo mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@localhost/$(DBNAME)  -o models

run:
	@./server
