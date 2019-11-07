REVISION := $(shell git rev-parse --short HEAD)

build:
	@mkdir -p build
	@go build -trimpath -ldflags "-s -w" -o build/genomator github.com/shizonic/go-genomator/cmd/go-genomator

clean:
	@rm -rf build
