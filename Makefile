## SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
## SPDX-License-Identifier: GPL-3.0-or-later

COVER_OUT=cover.out
COVER_HTML=cover.html

.PHONY: all test

all: test

test:
	CGO_ENABLED=1 go test -race -coverprofile=$(COVER_OUT) ./...
	go tool cover -html=$(COVER_OUT) -o $(COVER_HTML)
	go tool cover -func=$(COVER_OUT) | tail -n1
	fieldalignment ./...
