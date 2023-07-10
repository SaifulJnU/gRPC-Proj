# gRPC-Proj

---

### What is gRPC?

gRPC is an open source framework that handles RPC(Remote Procedure Call) and helps build scalable and fast APIs.
It ensures maximum API security, performance and scalability.
It allows the client and server applications to communicate transparently
It uses http/2


---

### Basic concepts of gRPC:
1. Protocol Buffers
2. Streaming 
3. HTTP/2
4. Channels


=> Protocol buffers or protobuf is used as interface Definition Language(IDL) and serialization toolset by gRPC.
Client and Server both must need to have same proto file. gRPC stores data and function contracts in the form of proto file. protoc is the compiler of protobuf. It generates client and server code using the dot proto file. It loads the code using the dot proto file into the memory at runtime and uses the in-memory schema to serialize or deserialize the binary message. We can exchange the data faster using protobuf as it requires fewer resources of cpu since data is converted into binary format and encoded message are smaller in size.

=> Streaming: The multiplexing capability that is sending multiple responses or receiving multiple requests together over a single TCP connection of HTTP/2 makes streaming possible

1. Server side streaming 
2. Client side streaming
3. Bi-directional streaming 

=> HTTP/2: Bi-directional streaming was possible full fledged. This kind of request/response multiplexing is made possible in HTTP/2 by introducing a new HTTP/2 layer called binary framing.

In previous http, to make multiple request and response it needs to make multiple connection. It can not able to handle multiple request and responses by single connection.

=> Channels: It supports multiple steams over multiple concurrent connections. HTTP/2 streams allow many simultaneous streams on one connection. Channels extend this concept by supporting multiple streams or multiple concurrent connections. It provides a way to connect to the grpc server on a specified address and port that are used creating the client stuff


---

## gRPC  Architecture:
```
<gRPC Server>---(proto-req/proto-res)---<gRPC Stub(Go Client)>
      |
(proto-req/proto-res)
      |
      |
<gRPC Stub(Ruby Client)>
```

---

### When to use gRPC:
1. Real-time(Communications services) 
2. For internal APIs
3. Multi-language(environments)

---

### Strengths of gRPC:
1. Performance(10 times faster performance). Because http/2 uses advanced compression so message loading is fast

2. Built-in Community features like metadata exchange, encryption, authentication, deadline, timeout or cancellation. It also provides load balancing, service discovery, and many more

3. Streaming, gRPC makes it much easier to build streaming services. 

4. Security

---

### Weaknesses of gRPC:
1. Limited Browser support: As gRPC uses HTTP/2 so it is impossible to call a gRPC service directly from browser. No modern browser provides the control needed all web request to support a gRPC client therefore a proxy layer and gRPC web are required to perform conversion between HTTP1.1 and HTTP/2.

2. Sleeper Learning curve: Many teams find gRPC challenging to learn. This is why more user still rely on RESTAPi

3. Non human readable format: protobuf compresses gRPC messages into a non human readable format. This compiler needs the messages interface description in the file to deserialize correctly so developers need additional tools like the gRPC command line tool to analyze protobuf payload on the wire and they also write manual requests and perform debugging.

4. No edge Caching
---


# Build-First-gRPC Project:
--------------------------
## Prerequisites:
1. Install Go
2. Install Protobuf
   
---

Step 1: At first create mod file. 
```
Ex: go mod init github.com/saifuljnu/demo-grpc
```
Step 2: create a proto file, where proto file holds the api descriptions. 
        Example of file: invoicer.proto

```
// define version
syntax = "proto3";

// where to put the generated code
option go_package = "github.com/saifuljnu/demo-grpc/invoicer";

// define service
message Amount {
    int64 amount = 1;
    string currency = 2;
}

message CreateRequest {
    Amount amount = 1;
    string from = 2;
    string to = 3;
    string VATnum = 4;
}

message CreateResponse {
    bytes pdf = 1;
    bytes docx = 2;
}

service Invoice {
    rpc Create(CreateRequest) returns (CreateResponse);
}
 
```
        
Step 3: create a folder where you want to store auto generated file and code.
Example of folder: invoice

Step 4: Run the following command,
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
Then,
```
protoc \
  --go_out=invoicer \
  --go_opt=paths=source_relative \
  --go-grpc_out=invoicer \
  --go-grpc_opt=paths=source_relative \
  invoicer.proto
```
In this way, you will have two auto generated file in the invoice folder. (Example of auto generated file and code: invoice_grpc.pb.go, invoice.pb.go)

Step 5: In case of auto generated code you will see some error. To resolve this error you just need to run the following command:
```
go get google.golang.org/grpc
```
Step 6: Now It is time to create main.go file

```
package main

import (
	"context"
	"log"
	"net"

	"github.com/saifuljnu/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoiceServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serviceRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	//resistring the server
	invoicer.RegisterInvoiceServer(serviceRegistrar, service)

	err = serviceRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve %s", err)
	}
}
```
Step 7: It is time to run the project
```
protoc --go_out=. invoicer.proto
```
Step 8:
Now to make request and to get response we can use bloomrpc. To do so we need to download the bloomrpc from its official github repo. As I am using ubuntu, so my case I have downloaded "bloomrpc_1.5.3_amd64.deb" then run the following command:
```
sudo dpkg -i bloomrpc_1.5.3_amd64.deb
bloomrpc
```
After that will have the beautiful UI of bloomrpc then you need to import your proto file and click on create option. And int tcp place your localhost and port. In my case it is localhost:8089

then, after clicking on the play button you will have your response message.

![Screenshot from 2023-07-10 23-07-00](https://github.com/SaifulJnU/gRPC-Proj/assets/47039014/912b67e6-9fb3-4c2a-9ea2-b8440852dd6b)

## Happy Coding!

                
        

