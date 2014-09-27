.PHONY: test install

prefix=github.com/work-with-hbc/annie/pkg/
packages=brain jsonconfig storage utils/http behaviour/http
clients=client

test: test-packages test-clients

test-packages:
	$(foreach pkg, $(packages), go test $(prefix)$(pkg);)

test-clients:
	cd $(clients) && $(MAKE) test


install: install-deps
	cd $(clients) && $(MAKE) install

install-deps:
	go get code.google.com/p/go-uuid/uuid
	go get code.google.com/p/snappy-go/snappy
	go get github.com/gorilla/context
	go get github.com/gorilla/mux
	go get github.com/syndtr/goleveldb/leveldb
