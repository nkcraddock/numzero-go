VENDOR_PATH = $(CURDIR)/_vendor
GOPATH := $(VENDOR_PATH):$(GOPATH)
SERVER_FILES := $(shell find cmd/server -type f -name "*.go" ! -name "*_test.go")

default: vendor build

run: 
	go run $(SERVER_FILES)

debug:
	go run $(SERVER_FILES) -r client/build

test:
	go test -v ./...

# download deps to _vendor (and remove the git repos)
vendor:
	GOPATH=$(VENDOR_PATH)
	mkdir -p $(VENDOR_PATH)
	go get -d github.com/emicklei/go-restful
	go get -d github.com/emicklei/go-restful/swagger
	go get -d github.com/onsi/gomega
	go get -d github.com/onsi/ginkgo/ginkgo
	go get -d code.google.com/p/go-uuid/uuid
	go get -d github.com/dgrijalva/jwt-go
	find $(VENDOR_PATH) -type d -name '.git' | xargs rm -rf

clean:
	rm -rf build/

build: clean clientdata
	mkdir -p build/
	CGO_ENABLED=0 go build -a -installsuffix cgo -o build/gooby --ldflags '-s' $(SERVER_FILES)

clientdata: 
	go get -u github.com/jteeuwen/go-bindata/...
	grunt --gruntfile client/Gruntfile.js build
	go-bindata -o "./clientdata.go" -pkg="gooby" -prefix="client/build/" client/build/...
