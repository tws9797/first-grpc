A simple unary RPC demo.

### References
- [How we use gRPC to build a client/server system in Go](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)
- [Service compiling successfully, but message structs not generating - gRPC/Go](https://stackoverflow.com/questions/71777702/service-compiling-successfully-but-message-structs-not-generating-grpc-go)
- [grpc with mustEmbedUnimplemented*** method](https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method)
- [what's meaning of insecure connection in grpc?](https://stackoverflow.com/questions/51008686/whats-meaning-of-insecure-connection-in-grpc)

Notes: The article [How we use gRPC to build a client/server system in Go](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2) was written in 2017, Golang Protobuf has updated since. Therefore, some command lines like `--go_out=plugins=grpc` are not supported in go anymore. 