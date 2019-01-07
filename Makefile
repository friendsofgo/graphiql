# general
WORKDIR = $(PWD)

# build CLI
LOCAL_TAG := $(shell git describe --tags --abbrev=0)
LOCAL_LDFLAGS = -s -X main.version=$(LOCAL_TAG)

# coverage
COVERAGE_REPORT = coverage.txt
COVERAGE_PROFILE = profile.out
COVERAGE_MODE = atomic

build-cli:
	go build -o bin/graphiql -ldflags "$(LOCAL_LDFLAGS)" cmd/graphiql/main.go

test:
	go test -v ./...

test-coverage:
	@cd $(WORKDIR); \
	echo "" > $(COVERAGE_REPORT); \
	for dir in `find . -name "*.go" | grep -o '.*/' | sort | uniq`; do \
		go test -v  -race $$dir -coverprofile=$(COVERAGE_PROFILE) -covermode=$(COVERAGE_MODE); \
		if [ $$? != 0 ]; then \
			exit 2; \
		fi; \
		if [ -f $(COVERAGE_PROFILE) ]; then \
			cat $(COVERAGE_PROFILE) >> $(COVERAGE_REPORT); \
			rm $(COVERAGE_PROFILE); \
		fi; \
	done; \