# build CLI
LOCAL_TAG := $(shell git describe --tags --abbrev=0)
LOCAL_BUILD := $(shell date +"%m-%d-%Y_%H_%M_%S")
LOCAL_LDFLAGS = -s -X main.version=$(LOCAL_TAG) -X main.build=$(LOCAL_BUILD)

build-cli:
	go build -o bin/graphiql -ldflags "$(LOCAL_LDFLAGS)" cmd/graphiql/main.go

install-cli: build-cli
	mv bin/graphiql $(GOPATH)/bin
	rm -rf bin