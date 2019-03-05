# moreutils

Simple command line utilities to help support and extend UNIX pipes workflows.

## Installation

Currently installation is done through the golang tooling provided by the `go` command. For now you must have this installed and your `$PATH` correctly configured.

All commands can be installed and updated with the following command:

```bash
go get -u -v github.com/Nekroze/moreutils/cmd/...
```

To install specific utilities simply adjust the module path:

```bash
go get -u -v github.com/Nekroze/moreutils/cmd/starts-with
```

## Utilities

### starts-with

Checks the input given via stdin (eg pipel) or the first argument, then the next (possible first) argument (called the needle) is checked against the input.

If the input starts with the needle the exit status will be `0`, otherwise it will be a non zero positive integer such as `1`.
