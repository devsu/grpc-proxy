package extras

import (
  "fmt"
  "strings"
  "google.golang.org/grpc/codes"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/mwitkow/grpc-proxy/proxy"
)

func GetDirector(config Config) func(context.Context, string) (*grpc.ClientConn, error) {
  return func(ctx context.Context, fullMethodName string) (*grpc.ClientConn, error) {
    for _, backend := range config.Backends {
      if strings.HasPrefix(fullMethodName, backend.Filter) {
        if (config.Verbose) {
          fmt.Printf("Found: %s > %s \n", fullMethodName, backend.Backend)
        }
        if (backend.WithInsecure) {
          return grpc.DialContext(ctx, backend.Backend, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
        }
        return grpc.DialContext(ctx, backend.Backend, grpc.WithCodec(proxy.Codec()))
      }
    }
    if (config.Verbose) {
      fmt.Println("Not found: ", fullMethodName)
    }
    return nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
  }
}