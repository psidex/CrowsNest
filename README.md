# CrowsNest

[![MainCI](https://github.com/psidex/CrowsNest/actions/workflows/mainci.yml/badge.svg)](https://github.com/psidex/CrowsNest/actions/workflows/mainci.yml)
[![ReleaseCI](https://github.com/psidex/CrowsNest/actions/workflows/releaseci.yml/badge.svg)](https://github.com/psidex/CrowsNest/actions/workflows/releaseci.yml)
[![Docker Pulls](https://img.shields.io/docker/pulls/psidex/crowsnest)](https://hub.docker.com/repository/docker/psidex/crowsnest)
[![Go Report Card](https://goreportcard.com/badge/github.com/psidex/crowsnest)](https://goreportcard.com/report/github.com/psidex/crowsnest)
[![buymeacoffee donate link](https://img.shields.io/badge/Donate-Beer-FFDD00.svg?style=flat&colorA=35383d)](https://www.buymeacoffee.com/psidex)

[Watchtower](https://github.com/containrrr/watchtower) for Git: automatically keep local Git repositories up to date with their remotes.

## Configuration

### Flags

`--run-once` or `-r`: Normally CrowsNest would loop forever, set this flag to run once then exit

`--config` or `-c`: Where to look for your config.yaml file (`.` and `$HOME` are automatically searched)

`--verbose` or `-v`: Write a lot more info to the log, useful for finding errors

`--logfile` or `-l`: Write the log to the given file instead of stdout. Should be a full path ending with the file name

### config.yaml

CrowsNest reads its configuration from a `config.yaml` file.

Example `config.yaml` using every possible option:

```yaml
# The list of repositories we want to watch
respositories:
  # The name of this repsoitory (can be anything, will show up in logs)
  peicecost:
    # The path to the root of the repo
    directory: D:\Code\piececost
    # Extra flags to provide to git when CrowsNest runs `git pull`
    gitpullflags: ["--verbose", "--autostash"]
    # How long to wait between pulls in seconds (defaults to 60), doesn't account for the time it takes to run cmds and pull
    interval: 900
    # A command to run before pulling, if this returns a non-zero exit code, the pull will not happen
    prepullcmd:
      # The binary to execute, e.g. /bin/bash on debian for a shell script
      binarypath: C:\Programs\dosomething.exe
      # Any flags/arguments for the binary
      flags: ["--user", "psidex"]
      # Where to execute this binary
      workingdirectory: D:\Code\piececost\otherdir
    # A command to be run after pulling, same options as prepullcmd, but a non-zero exit code won't change anything
    postpullcmd:
      binarypath: C:\Programs\dosomething.exe
      flags: ["--user", "psidex"]
      workingdirectory: D:\Code\piececost\otherdir
```

The only required option is the `directory` for each repo, and if you have set pre/post cmds then the `binarypath` and `workingdirectory` need to be valid for each of those.

Keep in mind that if you are running CrowsNest in a Docker container, all of the set directories and the pre and post pull binaries will need to be available and exectuable inside the container.

## Docker

If you want to try it out in Docker there are images available on [Docker Hub](https://hub.docker.com/repository/docker/psidex/crowsnest).

2 builds are published, `latest` which is inline with the latest commit to this repository, and the versioned tags that are inline with the GitHub releases.

### Example Docker Run

The crowsnest binary exists in and is run from the `/app` directory in the container.

I chose `/gitrepos/Apollo` in this example for no reason other than its short and descriptive, docker will create that directory automatically.

```bash
docker run -d --name crowsnest --restart unless-stopped \
    -v $(pwd)/config.yaml:/app/config.yaml:ro \
    -v /home/psidex/repos/Apollo:/gitrepos/Apollo \
    psidex/crowsnest:latest
```

My `config.yaml`:

```yaml
repositories:
  apollo:
    directory: /gitrepos/Apollo
```

Notice that the directory is the path inside the container, not the external path.

## Use Cases

This would be useful if you store configuration files or content in a Git repository and want to keep your local copies up to date with the most recent versions.

Personally I use this to keep my website up to date as the files are published to GitHub from my development machine but need to be on my server to be served to the internet.

## Bugs And Feature Requests

If you find a bug or would like to request a new feature please open an [Issue](https://github.com/psidex/CrowsNest/issues/new).

## Development

Requires [govvv](https://github.com/ahmetb/govvv) to build correctly.

See `build.ps1` or the `Dockerfile` for build commands.
