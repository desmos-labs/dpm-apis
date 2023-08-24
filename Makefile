#!/usr/bin/make -f
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

###############################################################################
###                                   All                                   ###
###############################################################################

all: lint test-unit


###############################################################################
###                               Build flags                               ###
###############################################################################

build_tags = netgo

# These lines here are essential to include the muslc library for static linking of libraries
# (which is needed for the wasmvm one) available during the build. Without them, the build will fail.
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

ldflags =
ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

ldflags := $(strip $(ldflags))
BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

###############################################################################
###                                 Build                                   ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

###############################################################################
###                                Linting                                  ###
###############################################################################
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	@echo "--> Running linter"
	@go run $(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go run $(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' -not -name '*_mock.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' -not -name '*_mock.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' -not -name '*_mock.go' | xargs goimports -w -local github.com/desmos-labs/dpm-apis
.PHONY: format

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

stop-test-db:
	@echo "Stopping test database..."
	@docker stop apis-test-db || true && docker rm apis-test-db || true
.PHONY: stop-test-db

start-test-db: stop-test-db
	@echo "Starting test database..."
	@docker run --name apis-test-db -e POSTGRES_USER=apis -e POSTGRES_PASSWORD=password -e POSTGRES_DB=apis -d -p 6433:5432 postgres
.PHONY: start-test-db

test-unit: start-test-db
	@echo "Executing unit tests..."
	@go test -p=1 -mod=readonly -v -coverprofile coverage.txt ./...
.PHONY: test-unit
