# Nitch Clone

<!--toc:start-->

- [Nitch Clone](#nitch-clone)
  - [Key differences from nitch](#key-differences-from-nitch)
  - [Dependencies](#dependencies)
  - [Building and running](#building-and-running)
  <!--toc:end-->

This program is inspired by
[nitch](https://github.com/ssleert/nitch) but written in Go.

## Key Differences from nitch

- Written in Go
- Different icons
- No word art (yet)
- Themes
- ~~Colors are randomized~~ (WIP)

## Dependencies

This program requires `uname` and `uptime`
to be available.

## Config and Themes

Themes ~~and styles~~ (WIP) can be set through `config.toml`
placed in `$XDG_HOME/nitch-clone`
(It will be`~/.config/nitch-clone/config.toml` for most users).

Currently 5 themes are available

- catppuccin-mocha
- catppuccin-frappe
- catppuccin-latte
- catppuccin-machiato
- grayscale

## Building and Running

> [!NOTE]
> For building and running `go` must be installed.
> This has only been tested on EndeavourOS.

1. Clone this [repository](https://github.com/sid314/nitch-clone).
2. Run `cd nitch-clone`
3. Run `go build`
4. Run `./nitch-clone` or add it to your path or copy it to `/usr/local/bin/`
