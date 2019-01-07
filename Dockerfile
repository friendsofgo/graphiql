FROM golang:alpine AS build

LABEL MAINTAINER = 'Friends of Go (it@friendsofgo.tech)'

RUN apk add --update git
WORKDIR /go/src/github.com/friendsofgo/graphiql
COPY . .
RUN TAG=$(git describe --tags --abbrev=0) \
    && LDFLAGS=$(echo "-s -X main.version="$TAG) \
    && go build -o /go/bin/graphiql -ldflags "$LDFLAGS" cmd/graphiql/main.go

# Building image with the binary
FROM alpine
COPY --from=build /go/bin/graphiql /bin/
ENTRYPOINT ["graphiql"]
