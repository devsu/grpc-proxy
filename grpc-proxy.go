package main

import (
  "github.com/mwitkow/grpc-proxy/proxy"
  "google.golang.org/grpc"
  "log"
  "net"
  "fmt"
  "os"
  "github.com/devsu/grpc-proxy/extras"
)

func main() {
  configurationFile := "./config.json"

  args := os.Args[1:]
  if len(args) > 0 {
    configurationFile = args[0]
  }

  config := extras.GetConfiguration(configurationFile)

  listen := ":50051"
  if config.Listen != "" {
    listen = config.Listen
  }

  lis, err := net.Listen("tcp", listen)

  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  fmt.Printf("Proxy running at %q\n", listen)

  server := grpc.NewServer(
    grpc.CustomCodec(proxy.Codec()),
    grpc.UnknownServiceHandler(proxy.TransparentHandler(extras.GetDirector(config))))

  if err := server.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}