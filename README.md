# toolbox-web

A codefly toolbox plugin that performs HTTP fetches behind a domain
allowlist. Canonical replacement for `curl` / `wget` in agent-driven
flows — the `codefly-dev/toolbox-bash` plugin refuses both binaries
and routes here.

## Tools

- `web.fetch(url, method?, timeout_ms?, headers?, body?)` — single HTTP
  request. Returns `{status_code, status_text, headers, body, truncated}`.
  Body capped at 4 MiB; oversized responses come back with
  `truncated: true`.

## Network policy (two layers)

1. **In-process allowlist** — set via `CODEFLY_TOOLBOX_ALLOWED_DOMAINS=foo.com,bar.com`.
   The toolbox refuses any URL whose host isn't on the list. Redirects
   are re-checked: if a request to an allowed host 302s to a non-allowed
   host, the redirect fails closed.
2. **Host OS sandbox** — codefly can additionally constrain outbound
   network at the kernel level (bwrap on Linux, sandbox-exec on macOS).
   The in-process check is authoritative on its own; the OS layer is
   defense in depth.

Auth tokens MUST come from host config (env vars set at plugin spawn),
NEVER from agent-supplied arguments. The toolbox doesn't carry secret
material in the contract.

## Configuration

| Env var                            | Default     | Purpose                                          |
| ---------------------------------- | ----------- | ------------------------------------------------ |
| `CODEFLY_TOOLBOX_VERSION`          | `0.0.0-dev` | Identity version surfaced via `Identity()`       |
| `CODEFLY_TOOLBOX_ALLOWED_DOMAINS`  | _(empty)_   | Comma-separated allowlist. Empty = deny all.     |

## Build & test

```bash
go build ./...
go test ./...
```

## Contract

This plugin implements the codefly Toolbox gRPC contract defined in
[`codefly-dev/core`](https://github.com/codefly-dev/core) at
`proto/codefly/services/toolbox/v0/toolbox.proto`. Loaded by the codefly
host via `agents.Serve` over a Unix domain socket.
