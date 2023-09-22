## Generate Buffer Protocol Schema

1.- Install dependencies

```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@lates
```

2.- Agregar Buffer Protocol Schema

```sh
echo "
syntax = "proto3";
option go_package = "/gen";

message Metadata {
    string id = 1;
    string title = 2;
    string description = 3;
    string director = 4;
}

" > movie.proto
```

3.- Generar entidades y servicios.

```sh
$ protoc -I=api --go_out=. --go-grpc_out=. movie.proto
```
