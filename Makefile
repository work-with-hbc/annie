.PHONY: test

prefix=github.com/work-with-hbc/annie/pkg/
packages=brain jsonconfig storage utils/http

test:
	$(foreach pkg, $(packages), go test $(prefix)$(pkg);)

install-deps:
	go get code.google.com/p/go-uuid/uuid
	go get code.google.com/p/snappy-go/snappy
	go get github.com/gorilla/context
	go get github.com/gorilla/mux
	go get github.com/syndtr/goleveldb/leveldb
