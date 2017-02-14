install:
	go install -v

build:
	go build -v ./...

lint:
	gometalinter --config .linter.conf

test:
	gucumber

deps:
	go get github.com/mitchellh/mapstructure
	go get github.com/ghodss/yaml
	go get github.com/r3labs/graph
	#go get github.com/nats-io/nats
	#go get github.com/r3labs/binary-prefix
	#go get github.com/ernestio/ernest-config-client

dev-deps: deps
	go get golang.org/x/crypto/pbkdf2
	go get github.com/ernestio/crypto
	go get github.com/ernestio/crypto/aes
	go get github.com/smartystreets/goconvey/convey
	go get -u github.com/gucumber/gucumber/cmd/gucumber
	go get github.com/tidwall/gjson
	go get github.com/alecthomas/gometalinter
	gometalinter --install


clean:
	go clean

dist-clean:
	rm -rf pkg src bin
