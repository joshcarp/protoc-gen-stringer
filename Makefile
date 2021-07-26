all: ci install
.PHONY: install test demo update-tests

test:
	go test -v ./... -count=1

help:			## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

%.pb.go: %.proto		## Generates .pb.go files from .proto files with protoc.
	protoc --go_out=. $<

install:	## Installs this project as a binary in your go binary directory.
	go install github.com/joshcarp/protoc-gen-stringer

tidy:			## Tidies up Go mod and source files.
	goimports -w $$(find . -type f -name '*.go' -not -name '*.pb.go')
	gofmt -s -w .
	go mod tidy

demo:			## Makes sure the demo directory still builds and compiles
	cd demo && make

ci: test				## Runs the same ci that is on master.
	golangci-lint run

example: 	## Executes proto to generate go code
	protoc -I . example/*.proto --go_out=paths=source_relative:. --stringer_out=paths=source_relative:.

options: 	## Executes proto to generate go code
	protoc -I stringerpb/ stringerpb/*.proto --go_out=paths=source_relative:stringerpb

docker:			## Builds the Docker image.
	docker build -t protoc-gen-stringer .

update-tests:		## Updates the code_generator_request.pb.bin for the go test cases.
	protoc --debug_out="tests/simple:tests/." ./tests/simple/*.proto
	protoc --debug_out="tests/double:tests/." ./tests/double/*.proto