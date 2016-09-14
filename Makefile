GO_OPTS=GOROOT=/usr/local/go GOPATH=$(PWD) GOBIN=$(PWD)/bin PATH=$(PATH):$(PWD)/bin
GO_BIN=/usr/local/go/bin/go


all: get bind
	$(GO_OPTS) $(GO_BIN) build -o gantry src/gantry.go

get:
	cd src && $(GO_OPTS) $(GO_BIN) get -u github.com/jteeuwen/go-bindata/...
	cd src && $(GO_OPTS) $(GO_BIN) get

bind:
	$(GO_OPTS) ./bin/go-bindata -o src/data.go data/...

install:
	install -m 755 gantry $(PREFIX)/gantry

clean:
	rm gantry