

SHELL = /bin/bash

HUGO ?= hugo --minify
BASE_PATH ?= /device-repository
BASE_URL ?= https://thethingsnetwork.org
FULL_URL = $(BASE_URL)$(BASE_PATH)

.PHONY: clean
clean:
	rm -f tools/build/build || true
	rm -r content/devices/* || true
	rm -rf public || true

.PHONY: go.deps
go.deps:
	pushd ./tools/build && go mod download && popd

.PHONY: go.build
go.build: clean
	pushd ./tools/build && go build . && ./build && popd

.PHONY: hugo.build.public
hugo.build.public:
	$(HUGO) --baseURL $(FULL_URL)

.PHONY: run
run:
	hugo server -b localhost$(BASE_PATH)
