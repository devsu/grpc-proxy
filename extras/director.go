package extras

import (
  "fmt"
  "strings"
  "google.golang.org/grpc/codes"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/mwitkow/grpc-proxy/proxy"
  "google.golang.org/grpc/credentials"
  "google.golang.org/grpc/grpclog"
)

func GetDirector(config Config) func(context.Context, string) (*grpc.ClientConn, error) {

  credentialsCache := make(map[string] credentials.TransportCredentials)

  return func(ctx context.Context, fullMethodName string) (*grpc.ClientConn, error) {
    for _, backend := range config.Backends {
      if strings.HasPrefix(fullMethodName, backend.Filter) {
        if (config.Verbose) {
          fmt.Printf("Found: %s > %s \n", fullMethodName, backend.Backend)
        }
        if backend.CertFile == "" {
          return grpc.DialContext(ctx, backend.Backend, grpc.WithCodec(proxy.Codec()),
            grpc.WithInsecure())
        }
        creds := GetCredentials(credentialsCache, backend)
        if creds != nil {
          return grpc.DialContext(ctx, backend.Backend, grpc.WithCodec(proxy.Codec()),
            grpc.WithTransportCredentials(creds))
        }
        grpclog.Fatalf("Failed to create TLS credentials")
        return nil, grpc.Errorf(codes.FailedPrecondition, "Backend TLS is not configured properly in grpc-proxy")
      }
    }
    if (config.Verbose) {
      fmt.Println("Not found: ", fullMethodName)
    }
    return nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
  }
}

func GetCredentials(cache map[string] credentials.TransportCredentials, backend Backend) credentials.TransportCredentials {
  if cache[backend.Backend] != nil {
    return cache[backend.Backend]
  }
  creds, err := credentials.NewClientTLSFromFile(backend.CertFile, backend.ServerName)
  if err != nil {
    grpclog.Fatalf("Failed to create TLS credentials %v", err)
    return nil
  }
  cache[backend.Backend] = creds
  return creds
}