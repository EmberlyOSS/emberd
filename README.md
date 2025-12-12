# emberd

emberd is a small CLI for managing systemd units (start/stop/restart/reload/status/logs) and viewing unit logs.

Table of contents
- [Prerequisites](#prerequisites)
- [Build](#build)
- [Install](#install)
- [Usage](#usage)
- [Running emberd as a systemd service](#running-emberd-as-a-systemd-service)
- [Permissions & best practices](#permissions--best-practices)
- [Development](#development)
- [Repository](#repository)

## Prerequisites

- Go 1.20+ (to build the binary)
- A Linux host using systemd with `systemctl` and `journalctl`

## Build

Build the binary from the repository root:

```bash
go build -o emberd ./...
```

## Install

Copy the built binary to a system path (example):

```bash
sudo cp emberd /usr/local/bin/
sudo chown root:root /usr/local/bin/emberd
sudo chmod 0755 /usr/local/bin/emberd
```

## Usage

Common commands:

- Start a unit: `emberd start nginx.service`
- Stop a unit: `emberd stop nginx.service`
- Restart: `emberd restart nginx.service`
- Reload (if supported): `emberd reload nginx.service`
- Show status: `emberd status nginx.service`
- Tail logs: `emberd logs nginx.service -f`
- Show last N lines: `emberd logs nginx.service -n 200`

You can also pass the unit via flag: `emberd start -u nginx.service`

Run `emberd --help` for full command help.

## Running emberd as a systemd service

emberd is mainly a CLI for interactive and automated service management. If you need a daemon that tails logs continuously, create a systemd unit that runs `emberd logs <unit> -f`.

Example unit file (save as `/etc/systemd/system/emberd-logs-nginx.service`):

```ini
[Unit]
Description=emberd log tailer for nginx
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/emberd logs nginx.service -f
Restart=always
RestartSec=5
# Consider running as a non-root user if possible
User=root

[Install]
WantedBy=multi-user.target
```

Enable and start it:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now emberd-logs-nginx.service
sudo journalctl -u emberd-logs-nginx.service -f
```

## Permissions & best practices

- Managing system services requires privileges. Use `sudo` or run as `root` for system units.
- For user-scoped units/journals, use `systemctl --user` and `journalctl --user` instead.
- Prefer ad-hoc CLI usage or scripted automation over running a privileged long-lived process.
- If non-root operators need to control services, use Polkit rules to grant specific permissions rather than running everything as root.

## Development

Helpful make targets are provided in the repository:

```bash
make build   # build the binary
make install # copy to /usr/local/bin (Makefile uses `install`)
make fmt     # run gofmt
make vet     # run go vet
make tidy    # go mod tidy
```
