GO_OPTS=GOROOT=/usr/local/go-1.7 GOPATH=$(PWD)
GO_BIN=/usr/local/go-1.7/bin/go

all:
	$(GO_OPTS) $(GO_BIN) build -o gantry src/gantry.go

get:
	$(GO_OPTS) $(GO_BIN) get -d

clean:
	rm gantry