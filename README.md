# tyk-mw-grpcgo-lambda

```
go get -u github.com/asoorm/tyk-mw-grpcgo-lambda
tyk-mw-grpcgo-lambda
```

Modify tyk.conf as follows:

```
"coprocess_options": {
  "enable_coprocess": true,
  "coprocess_grpc_server": "unix:///tmp/foo.sock",
},
```

Modify API Definition as follows:

```
"custom_middleware": {
  "pre": [
    {
      "name": "echo"
    }
  ],
  "driver": "grpc",
}
```

The above will attempt to invoke the `echo` named lambda function within `eu-west-2` (hardcoded in gRPC server main.go).
