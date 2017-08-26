## playtime with gRPC; let's see the unicorns fly...

Tools:
  get gRPC:
    `go get google.golang.org/grpc`
  get protobuf and protoc-gen-go plugin:
    `brew install protobuf`
    `go get -u github.com/golang/protobuf/protoc-gen-go`
  ensure the protoc-gen-go plugin is on you path:
    `export PATH=$PATH:$GOPATH/bin`

Update auto-generated GRPC code:
  `protoc -I unicorns-can-fly/ unicorns-can-fly/unicorns-can-fly.proto --go_out=plugins=grpc:unicorns-can-fly`