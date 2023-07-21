# sprox

`sprox` is a single binary tcp-over-websocket proxy.

## Building

You must have `golang` available. Then, simply run:

```bash
make
```

## Usage

First, start the server, pointing to the tcp address you want to proxy and the websocket port you want to listen on:

```bash
$ sprox proxy -listen :8432 1.2.3.4:5432
```

Then, start the client, pointing to the proxy server address, and the local tcp port you want to listen on:

```bash
$ sprox connect -listen :5432 ws://server:8432
```

There's some configuration options available, see `sprox proxy -h` and `sprox connect -h` for more details. Generally, you should be able to use the defaults, unless you have loopback-specific port requirements.

## Authx

Since this relies on websockets, we shift the authx responsibility up to the edge of the network, and assume you will be running this behind some identity-aware gateway such as Istio. When a client mounts a remote volume, they can pass either a `-token` or `-token-cmd` option to the `mount` command. If `-token` is passed, it will be used as the bearer token for the websocket connection. If `-token-cmd` is passed, it will be executed and the output will be used as the bearer token. This is useful if you want to use a dynamic token, such as a JWT, and don't want to have to manage the token yourself.

In the same vein, the example above shows the client connecting directly to the plain-text TCP port of the proxy server (`ws://server:port`) - in practice, it's assumed this will be run behind a TLS-terminating and identity-aware gateway, so your client will actually be connecting to something like `wss://sprox.example.com/my-service`. If you have a direct TCP netpath from the client to the server on arbitrary ports in a trusted network, just connect to it directly.
