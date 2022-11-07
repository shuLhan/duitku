## SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
## SPDX-License-Identifier: GPL-3.0-or-later

.PHONY: all test

all: test

test:
	CGO_ENABLED=1 go test -race ./...
	fieldalignment ./...
