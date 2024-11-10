---
id: grpc-external-service
title: Working with External gRPC Services
sidebar_label: External gRPC Services
---
Oftentimes, you will want to include in your gRPC server, methods that are not automatically generated from
your Ent schema. To achieve this result, define the methods in an additional service in an additional `.proto` file
in your `entpb` directory. 

:::info 

Find the changes described in this section in  [this pull request](https://github.com/rotemtam/fluent-grpc-example/pull/7/files).

:::


For example, suppose you want to add a method named `TopUser` which will return the user with the highest ID number.
To do this, create a new `.proto` file in your `entpb` directory, and define a new service:

```protobuf title="fluent/proto/fluentpb/ext.proto"
syntax = "proto3";

package fluentpb;

option go_package = "github.com/rotemtam/fluent-grpc-example/fluent/proto/fluentpb";

import "fluentpb/fluentpb.proto";

import "google/protobuf/empty.proto";


service ExtService {
  rpc TopUser ( google.protobuf.Empty ) returns ( User );
}
```

Next, update `entpb/generate.go` to include the new file in the `protoc` command input:

```diff title="fluent/proto/fluentpb/generate.go"
- //go:generate protoc -I=.. --go_out=.. --go-grpc_out=.. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --entgrpc_out=.. --entgrpc_opt=paths=source_relative,schema_path=../../schema entpb/fluentpb.proto 
+ //go:generate protoc -I=.. --go_out=.. --go-grpc_out=.. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --entgrpc_out=.. --entgrpc_opt=paths=source_relative,schema_path=../../schema entpb/fluentpb.proto entpb/ext.proto
```

Next, re-run code generation:

```shell
go generate ./...
```

Observe some new files were generated in the `ent/proto/fluentpb` directory:

```shell
tree
.
|-- entpb.pb.go
|-- entpb.proto
|-- entpb_grpc.pb.go
|-- entpb_user_service.go
// highlight-start
|-- ext.pb.go
|-- ext.proto
|-- ext_grpc.pb.go
// highlight-end
`-- generate.go

0 directories, 9 files
```

Now, you can implement the `TopUser` method in `ent/proto/fluentpb/ext.go`:

```go title="fluent/proto/fluentpb/ext.go"
package fluentpb

import (
	"context"

	"github.com/rotemtam/fluent-grpc-example/fluent"
	"github.com/rotemtam/fluent-grpc-example/fluent/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ExtService implements ExtServiceServer.
type ExtService struct {
	client *fluent.Client
	UnimplementedExtServiceServer
}

// TopUser returns the user with the highest ID.
func (s *ExtService) TopUser(ctx context.Context, _ *emptypb.Empty) (*User, error) {
	id := s.client.User.Query().Aggregate(fluent.Max(user.FieldID)).IntX(ctx)
	user := s.client.User.GetX(ctx, id)
	return toProtoUser(user)
}

// NewExtService returns a new ExtService.
func NewExtService(client *fluent.Client) *ExtService {
	return &ExtService{
		client: client,
	}
}

```

### Adding the New Service to the gRPC Server

Finally, update `cmd/server.go` to include the new service:

```go title="cmd/server.go"
package main

import (
	"context"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rotemtam/fluent-grpc-example/fluent"
	"github.com/rotemtam/fluent-grpc-example/fluent/proto/fluentpb"
	"google.golang.org/grpc"
)

func main() {
	// Initialize an fluent client.
	client, err := fluent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated User service.
	svc := entpb.NewUserService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

    // highlight-start
	// Register the User service with the server.
	entpb.RegisterUserServiceServer(server, svc)
	// highlight-end

	// Register the external ExtService service with the server.
	entpb.RegisterExtServiceServer(server, entpb.NewExtService(client))

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}

```