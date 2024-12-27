.PHONY: build clean rebuild unit

PARALLELISM = 4

ifndef PARALLEL_PROCS
PARALLEL_PROCS = 24
endif

mkfile_path := $(word $(words $(MAKEFILE_LIST)),$(MAKEFILE_LIST))
mkfile_dir:=$(shell cd $(shell dirname $(mkfile_path)); pwd)
goroot = $(mkfile_dir)/..
export GOBIN:=$(mkfile_dir)/bin

build:
	mkdir -p $(GOBIN)
	go install -v -p $(PARALLEL_PROCS) $(GOLANG_FLAGS) ./...

clean:
	rm -rf ${goroot}/pkg
	rm -rf ${goroot}/bin
	rm -rf $(mkfile_dir)/bin

rebuild: clean build

unit:
	go test ./...
