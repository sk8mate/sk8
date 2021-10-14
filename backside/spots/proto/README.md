Generate go structs from project root: 

``` protoc --go_out=. ./proto/*.proto ```

``` protoc -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate --go_out=. --validate_out="lang=go:." ./proto/spots_request.proto ```

