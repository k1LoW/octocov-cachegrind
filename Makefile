PKG = github.com/k1LoW/octocov-cachegrind
COMMIT = $$(git describe --tags --always)
OSNAME=${shell uname -s}
ifeq ($(OSNAME),Darwin)
	DATE = $$(gdate --utc '+%Y-%m-%d_%H:%M:%S')
else
	DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
endif

export GO111MODULE=on

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

default: test

ci: depsdev test test-integration

test:
	go test ./... -coverprofile=coverage.out -covermode=count

test-integration:
	setarch `uname -m` -R valgrind --tool=cachegrind --cachegrind-out-file=cachegrind.out --I1=32768,8,64 --D1=32768,8,64 --LL=8388608,16,64 ls
	cat cachegrind.out | go run ./cmd/octocov-cachegrind/main.go --tee > custom_metrics_cachegrind.json

lint:
	golangci-lint run ./...

build:
	go build -ldflags="$(BUILD_LDFLAGS)" -o octocov-cachegrind cmd/octocov-cachegrind/main.go

depsdev:
	go install github.com/Songmu/ghch/cmd/ghch@latest
	go install github.com/Songmu/gocredits/cmd/gocredits@latest

prerelease_for_tagpr: depsdev
	go mod tidy
	gocredits -skip-missing -w .
	cat _EXTRA_CREDITS >> CREDITS
	git add CHANGELOG.md CREDITS go.mod go.sum

release:
	git push origin main --tag
	goreleaser --clean

.PHONY: default test
