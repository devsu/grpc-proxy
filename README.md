# grpc-proxy
A runnable, configurable Go reverse proxy that allows for rich routing of gRPC calls with minimum overhead. Built on top of [mwitkow/grpc-proxy](https://github.com/mwitkow/grpc-proxy)

## Status

Alpha. It's working, but it hasn't been battle tested.

## Installation

Make sure you [installed go](https://golang.org/doc/install), then just run:

```bash
go get github.com/devsu/grpc-proxy
```

## Usage

```bash
grpc-proxy [configFile]
```

If no `configFile` is specified, it will try to read `./config.json`
 
## Features

- Redirects the request to the backend using a prefix of `serviceFullName`.
- Support TLS or insecure connections in both the proxy, and the backends.
 
## Configuration

Use this as an example.

```json
{
  "verbose": true,
  "listen": "grpc.localhost:50051",
  "certFile": "./ssl/grpc.localhost.pem",
  "keyFile": "./ssl/grpc.localhost.key",
  "backends": [
    {
      "backend": "grpc.localhost:3000",
      "filter": "/myapp.Greeter",
      "certFile": "./ssl/grpc.localhost.pem",
      "serverName": "grpc.localhost"
    },
    {
      "backendEnv": "ANOTHER_BACKEND",
      "filter": "/com.anotherApp."
    }
  ]
}
```

The proxy will listen in a secure channel. The first backend will be secure as well, whereas the second backend is not.

### Options

| Option | Description | Default |
|--------|-------------|---------|
| verbose | Prints a log of every request | false |
| listen  | What address should the proxy listen to, in the `host:port` format. If you are not using TLS you can use the `":port"` format (without the address) to listen all interfaces | `":50051"` |
| certFile | Path to the `.pem` file to be used by the proxy. If no file is specified, it will create an insecure proxy server. | |
| keyFile  | Path to the `.key` file to be used by the proxy. If no file is specified, it will create an insecure proxy server. | |
| backends | An array of backends. See below | |

### Backend 

The backend options are:

| Option | Description |
|--------|-------------|
| backend | The address that the proxy should connect to. |
| backendEnv | Environment variable to read the value of backend from. |
| filter  | The prefix of the `fullServiceName` that will be used to match calls against the backends. <br><br> The *service full name* it's compound by the package name + the service name. You can use the full name of the service or the full name of the package, or just part of it. Also note that it always starts with an slash `/`.|
| certFile | Path to the `.pem` file used to connect to the backend (if the backend has TLS configured). |
| serverName | Server name of the backend. Used to create the TLS client. Must match with the certificate. |

## Examples

Can be found in the examples folder in this repo.

## Roadmap

- Support other strategies for choosing the backend (besides matching by *prefix*).
- Add tests
- Add examples in go

## License and Credits

Devsu LLC. Copyright 2017. 

It uses [mwitkow/grpc-proxy](https://github.com/mwitkow/grpc-proxy) which is licensed under Apache2 too.

Built by the [GRPC experts](https://devsu.com) at Devsu.