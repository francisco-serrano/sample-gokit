SHELL := /bin/bash

models:
	go generate

migrations:
	sql-migrate up -env=local
