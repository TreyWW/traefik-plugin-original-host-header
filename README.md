# Forward Original Host Header

A [Traefik](https://traefik.io/) middleware plugin that adds a customisable header (defaulting to `X-Original-Host`) 
containing the original `Host` value of the incoming request. Useful when rewriting the `Host` header for upstream services while
preserving the original host information.

## Installation

### Static Configuration (`traefik.yaml`)

```yaml
experimental:
  plugins:
    originalHostHeaders:
      moduleName: github.com/treyww/traefik-plugin-original-host-header
      version: v0.1.0
```

## Usage

### Dynamic Configuration (`rules.yaml`)

```yaml
http:
  routers:
    my-router:
      rule: "Host(`localhost`)"
      entryPoints:
        - http
      service: my-service
      middlewares:
        - originalHostHeaders

  services:
    my-service:
      loadBalancer:
        servers:
          - url: "http://127.0.0.1"

  middlewares:
    originalHostHeaders:
      plugin:
        originalHostHeaders:
          headerName: X-Original-Host
```

> For more information about Traefik plugins, see the [Traefik Plugin Documentation](https://plugins.traefik.io/install).
