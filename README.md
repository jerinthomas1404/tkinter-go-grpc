## Understanding gRPC using Golang

### What did I learn?
1.  On superficial level understood why gRPC is scalable in microservice architecture.
2. Its asynchronous nature.
3. Its support for streaming as it works on top of http v2.
4. The sequence is preserved while streaming.

### What did I implement?
1. Created protobuf definition files, and generated code using protoc and its plugin.
2. It generates both : server skeleton and client stub.
3. Implemented multiple streaming - unary, server, client and bi-directional.

#### References
1. - [ ] gRPC Up & Running - Reading
2. - [x] YT @AkhilSharma - Watched
