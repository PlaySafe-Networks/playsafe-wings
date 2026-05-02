# PlaySafe Wings

This repository is a fork of [pterodactyl/wings](https://github.com/pterodactyl/wings),
maintained by PlaySafe Software for use as the per-host agent in the PlaySafe
game-server platform.

The upstream project is licensed under the MIT License. See `LICENSE` for the
full text. All upstream copyright notices are preserved.

## PlaySafe modifications

| Commit | Description |
|---|---|
| feat/playsafe-config-paths | Relocate filesystem defaults under `/srv/wings`, `/var/log/playsafe-wings`, `/tmp/playsafe-wings`. |
| chore/strip-telemetry | Diagnostics upload endpoint defaults to `paste.playsafe.dev` instead of `ptero.co`. Wings has no auto-telemetry. |
| feat/playsafe-auth-contract | Default `PanelLocation` points at `https://internal.playsafe.dev`. |
| feat/branding | Banner reads "PlaySafe Wings (fork of pterodactyl/wings)". User-Agent header rebranded. |
| feat/hibernation-hooks | Placeholder routes `POST /api/servers/:server/hibernate` and `POST /api/servers/:server/wake`. Real implementation lands in Phase 4. |

## Tracking upstream

Upstream `pterodactyl/wings:develop` is mirrored at the local `upstream-tracking`
branch. Each upstream release triggers a manual rebase: `git rebase
upstream/develop` on `main`, resolve conflicts per commit, run tests, deploy to
staging.

## Reporting issues

PlaySafe-specific issues belong here. Upstream-affecting bugs (Docker driver,
SFTP, filesystem, install sandbox, console WebSocket, backups, server transfer)
should be filed at `pterodactyl/wings`.
