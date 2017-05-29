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

PID            = .pid
DBNAME         = dev
MYSQL_USER     = root
MYSQL_PASSWORD = root
.PHONY: all start depend bootstrap generate models client build run

########################################
# BACKEND
########################################
all: clean depend generate build

depend:
	@glide install

clean:
	@rm -rf ./backend/app
	@rm -rf ./backend/client
	@rm -rf ./backend/tool
	@rm -rf ./backend/public/swagger
	@rm -rf ./backend/public/schema
	@rm -rf ./backend/public/js
	@rm -f ./server

bootstrap:
	@goagen main    -d github.com/kawabet/goa-react-ts/design -o ./back/controllers

generate:
	@goagen app     -d github.com/kawabet/goa-react-ts/design -o ./backend
	@goagen swagger -d github.com/kawabet/goa-react-ts/design -o ./backend/public
	@goagen schema  -d github.com/kawabet/goa-react-ts/design -o ./backend/public
	@goagen client  -d github.com/kawabet/goa-react-ts/design -o ./backend
	@goagen js      -d github.com/kawabet/goa-react-ts/design -o ./backend/public

build:
	@echo build the app...
	@go build -o server

models:
	@xo mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@localhost/$(DBNAME)  -o ./backend/models

kill:
	@kill `cat $(PID)` || true

restart: kill
	@echo restart the app...
	@./server run & echo $$! > $(PID)

run:
	@./server


########################################
# FRONTEND
########################################

generate-front:
	@swagger-codegen generate -l typescript-fetch -i ./backend/public/swagger/swagger.json -o ./frontapi
	@jq -s '.[0] * .[1]' frontapi/package.json frontapi/package_replace.json > replaced_package.json
	@rm frontapi/package.json
	@mv replaced_package.json frontapi/package.json

build-front:
	@cd frontapi && yarn
	@cd frontapi && npm link
	@cd frontend && npm link frontapi
	@cd frontend && yarn build

