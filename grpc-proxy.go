package main

import (
  "github.com/mwitkow/grpc-proxy/proxy"
  "google.golang.org/grpc"
  "log"
  "net"
  "fmt"
  "smartmate.io/proxy/extras"
  "os"
)

func main() {
  configurationFile := "./config.json"

  args := os.Args[1:]
  if len(args) > 0 {
    configurationFile = args[0]
  }

  config := extras.GetConfiguration(configurationFile)

  address := ":50051"
  if config.Address != "" {
    address = config.Address
  }

  lis, err := net.Listen("tcp", address)

  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  fmt.Printf("Proxy running at %q\n", address)

  server := grpc.NewServer(
    grpc.CustomCodec(proxy.Codec()),
    grpc.UnknownServiceHandler(proxy.TransparentHandler(extras.GetDirector(config))))

  if err := server.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}