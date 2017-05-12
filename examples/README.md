# grpc-proxy examples

Examples on how to use grpc proxy.

Just [install *grpc-proxy*](https://github.com/devsu/grpc-proxy#installation) and then:

- Start the backends (microservices)

```bash
cd secure-microservice-node
npm i
node .
```

```bash
cd insecure-microservice-node
npm i
node .
```

- Start the proxy. We have two: one insecure that will use port `50051` and one secure that will use `50052`. 

From the examples folder run the following commands, each one on their own terminal:

```bash 
grpc-proxy config-insecure-proxy.json
grpc-proxy config-secure-proxy.json
```

- Test with the client

```bash
cd client-node
npm i
node .
```

## Why are the microservices in node?

I had them at hand ready to use. Send me a PR with examples in go. They should be easy to do, and I'll be happy to merge them.
