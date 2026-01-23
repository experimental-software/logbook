# Logbook [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

The Logbook project provides a command-line application for Markdown-based chronological note-taking.

## Setup

If a computer has the [go command](https://go.dev/dl/) and [Git](https://git-scm.com) installed, the _Logbook_ can be
installed by cloning its Git repository and then running the `go install` command.

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
# The directory where new logbook entries are added.
logDirectory: ~/Logs

# The directory where logbook entries are moved when they are archived.
archiveDirectory: ~/Archive
```

## Usage

### Add logbook entry

```sh
# Add logbook entry
logbook2 add "${TITLE}"

# Add logbook entry and open its root directory in a text editor
${EDITOR} $(logbook2 add "${TITLE}")
```

### Search logbook entries

```sh
logbook2 search "${SEARCH_TERM}"
```

### Archive logbook entries

```sh
# Archive single logbook entry
logbook2 archive "${PATH}"

# Archive multiple logbook entries
logbook2 archive $(logbook2 search --output-format list "${SEARCH_TERM}")
```

### Remove logbook entries

```sh
# Remove single logbook entry
logbook2 remove "${PATH}"

# Remove multiple logbook entries
logbook2 remove $(logbook2 search --output-format list "${SEARCH_TERM}")
```

### Customization

User-specific utilities may be defined with shell features, e.g. this Bash alias and functions on a macOS computer that has VS Code installed:

```sh
alias log=logbook2

# Creates logbook entry with title "Scratch Note" and opens it in VS Code.
function note() {
  local LOGBOOK_ENTRY_TITLE=$1
  if [[ -z "$LOGBOOK_ENTRY_TITLE" ]]; then
    LOGBOOK_ENTRY_TITLE="Scratch Note"
  fi
  LOGBOOK_ENTRY=$(log add "$LOGBOOK_ENTRY_TITLE")
  code "$LOGBOOK_ENTRY"
  code "$LOGBOOK_ENTRY"/*.md
}

# Searches for logbook entries. If one result, open logbook entry in VS Code.
# Otherwise, print search results.
function notes() {
  local SEARCH_TERM="$1"
  local ENTRY_LIST=$(log search "$SEARCH_TERM" --output-format list)
  local NUMBER_OF_ENTRIES=$(echo "$ENTRY_LIST" | wc | perl -pe 's/.*?(\d+).*/$1/')
  if [[ $NUMBER_OF_ENTRIES -eq 1 ]]; then
    code $ENTRY_LIST
  else
    log search "$SEARCH_TERM"
  fi
}
```

– `~/.bash_profile`

## Testing

### Component test

```sh
go test ./... -coverprofile=./cov.out
```

With the help of the [gremlins](https://gremlins.dev/) program, the tests can be executed with mutations:

```sh
gremlins unleash
```

### Component integration test

```sh
go run main.go
go run main.go search
go run main.go add "Just a test"
go run main.go archive /path/to/2026/01/11/17.28_wip
go run main.go search -a 
```

## Maintenance

### Static code analysis

https://sonarcloud.io/summary/overall?id=experimental-software_logbook&branch=main

## Alternative projects

- [Paper-based engineering logbook](https://github.com/experimental-software/logbook/wiki/Paper%E2%80%90based-engineering-logbook)
- [QOwnNote](https://www.qownnotes.org)
- [Emacs OrgMode](https://orgmode.org)
- [Evernote](https://evernote.com)
- [Roam Research](https://roamresearch.com)
- [Quiver](https://yliansoft.com/)
- [Notion](https://www.notion.so/product)
- [Obsidian](https://obsidian.md/)
- [Joplin](https://joplinapp.org/)
- [Zettelkasten](https://zettelkasten.de/)
- [Dendron](https://wiki.dendron.so/)

