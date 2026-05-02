# PlaySafe Wings

PlaySafe Wings is the per-host agent of the PlaySafe game-server platform. It runs on each Wings node, talks to the local Docker daemon, exposes an HTTP control plane plus a built-in SFTP server, and is driven by the PlaySafe Panel (`playsafe-back`).

This repository is a fork of [pterodactyl/wings](https://github.com/pterodactyl/wings), maintained by PlaySafe Software. See [`NOTICE-PLAYSAFE.md`](./NOTICE-PLAYSAFE.md) for the list of fork modifications and upstream tracking strategy. Upstream is licensed under MIT and all upstream copyright notices are preserved.

## What it does

* HTTP API for server lifecycle: install, start, stop, restart, kill, reinstall, transfer.
* WebSocket consumers for live console, logs and resource stats.
* Built-in SFTP server, authenticated against the Panel.
* Filesystem with XFS project quotas, sandboxed install scripts, server transfer between nodes.
* Backups (local and S3), egg ecosystem compatibility for game and app recipes.
* Hibernation hooks (added in this fork): `POST /api/servers/:server/hibernate` and `POST /api/servers/:server/wake`. The Phase 0 build ships placeholders; Phase 4 lands the real implementation.

## Status

This repository is at `v0.1.0-playsafe`, the foundation build. The PlaySafe modifications applied on top of upstream are listed in `NOTICE-PLAYSAFE.md`.

## Building

```
make
```

Produces cross-compiled Linux binaries at `build/wings_linux_amd64` and `build/wings_linux_arm64`. Wings is Linux-only at runtime; building on macOS works via cross-compile but the binary itself does not run there.

## Reporting issues

PlaySafe-specific issues belong here, in [PlaySafe-Networks/playsafe-wings](https://github.com/PlaySafe-Networks/playsafe-wings). For bugs in upstream-shared code (Docker driver, SFTP server, filesystem, install sandbox, console WebSocket, backups, server transfer), please file at [pterodactyl/wings](https://github.com/pterodactyl/wings).

## License

MIT, inherited from upstream. See [`LICENSE`](./LICENSE).
