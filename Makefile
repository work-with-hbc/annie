.PHONY: test

prefix=github.com/bcho/annie/pkg/
packages=brain jsonconfig storage

test:
	$(foreach pkg, $(packages), go test $(prefix)$(pkg);)
