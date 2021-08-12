# CrowsNest

[![Build status](https://github.com/psidex/crowsnest/workflows/CI/badge.svg)](https://github.com/psidex/crowsnest/actions)
![Docker Pulls](https://img.shields.io/docker/pulls/psidex/crowsnest)
[![Go Report Card](https://goreportcard.com/badge/github.com/psidex/crowsnest)](https://goreportcard.com/report/github.com/psidex/crowsnest)
[![buymeacoffee donate link](https://img.shields.io/badge/Donate-Beer-FFDD00.svg?style=flat&colorA=35383d)](https://www.buymeacoffee.com/psidex)

[Watchtower](https://github.com/containrrr/watchtower) for Git: automatically keep local Git repositories up to date with their remotes.

## Configuration

### Flags

`--run-once` or `-r`: Normally CrowsNest would loop forever, set this flag to run once then exit

`--config` or `-c`: Where to look for your config.yaml file (`.` and `$HOME` are automatically searched)

`--verbose`: Write a lot more info to the log, useful for finding errors

### config.yaml

Example using every possible option:

```yaml
# The list of repositories we want to watch
respositories:
  # The name of this repsoitory
  peicecost:
    # The root of the repo
    directory: D:\Code\piececost
    # Extra flags to provide to git when runnning git pull
    gitpullflags: ["--verbose", "--autostash"]
    # How long to wait between checks and/or pulls in seconds (defaults to 60)
    interval: 900
    # A command to run before pulling, if this returns a non-zero exit code, the pull will not happen
    prepullcmd:
      # The binary to execute, e.g. /bin/bash on debian for a script
      binarypath: C:\Programs\dosomething.exe
      # Any flags/arguments for the binary
      flags: ["--user", "psidex"]
      # Where to execute this binary
      workingdirectory: D:\Code\piececost\otherdir
    # A command to be run after pulling, same options as prepullcmd
    postpullcmd:
      binarypath: C:\Programs\dosomething.exe
      flags: ["--user", "psidex"]
      workingdirectory: D:\Code\piececost\otherdir
```

Keep in mind that if you are running CrowsNest in a Docker container, the pre and post pull binaries will need to be exectuable inside the container.

## Build

Requires [govvv](https://github.com/ahmetb/govvv) to build correctly.

See `build.ps1` or the `Dockerfile` for build commands.

## Use Cases

This would be useful if you store configuration files or content in a Git repository and want to keep your local copies up to date with the most recent versions.

Personally I use this to keep my website up to date as the files are published to GitHub from my development machine but need to be on my server to be servered to the internet.
