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

all: clean generate build

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf public/swagger
	@rm -rf public/schema
	@rm -rf public/js
	@rm -f api

bootstrap:
	@goagen main    -d github.com/kawabet/goa-react-ts/design -o controllers

generate:
	@goagen app     -d github.com/kawabet/goa-react-ts/design
	@goagen swagger -d github.com/kawabet/goa-react-ts/design -o public
	@goagen schema  -d github.com/kawabet/goa-react-ts/design -o public
	@goagen client  -d github.com/kawabet/goa-react-ts/design
	@goagen js      -d github.com/kawabet/goa-react-ts/design -o public

build:
	@go build -o api

