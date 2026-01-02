# Logbook [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

The Logbook project provides a command-line application for Markdown-based chronological note-taking.

## Setup

Provided that a computer has the SDK for the [go](https://go.dev) programming language installed, the _Logbook_ can be
installed by cloning its git repository and then running the `go install` command.

```sh
git clone git@github.com:experimental-software/logbook.git && cd ./logbook
go install
```

Then the program can be executed with the `logbook2` command:

```sh
logbook2
```

In the `~/.config/logbook/config.yaml` file it can be configured what directories are used for reading and writing log entries.

The following snippet shows the configuration options with their default values:

```yaml
# The directory where new log entries are added.
logDirectory: ~/Logs
```

## Usage

### Search log entries

```sh
logbook2 search $SEARCH_TERM
```

## Development

```sh
go run main.go
go run main.go search
go run main.go add "Just a test"
```

## Maintenance

### Static code analysis

https://sonarcloud.io/summary/overall?id=experimental-software_logbook&branch=main

## Alternative projects

- [logbook-dart](https://github.com/experimental-software/logbook-dart)
- [Paper-based engineering notebook](https://www.youtube.com/watch?v=xaFqpd7lNM4)
- [QOwnNote](https://www.qownnotes.org)
- [Emacs OrgMode](https://orgmode.org)
- [Evernote](https://evernote.com)
- [Roam Research](https://roamresearch.com)
- [Quiver](https://yliansoft.com/)
- [Notion](https://www.notion.so/product)
- [Obsidian](https://obsidian.md/)
- [Joplin](https://joplinapp.org/)
- [Zettelkasten](https://zettelkasten.de/)
