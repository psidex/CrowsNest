# CrowsNest

[Watchtower](https://github.com/containrrr/watchtower) for Git

## Configuration

### Flags

`--run-once` or `-r`: Normally CrowsNest would loop forever, set this flag to run once then exit

`--config` or `-c`: Where to look for your config.yaml file (`.` and `$HOME` are automatically searched)

### config.yaml

Example:

```yaml
respositories:
  peicecost:
    directory: "D:\\Code\\piececost"
    remote: "https://github.com/psidex/PieceCost.git"
    gitflags: ["--verbose", "--autostash"]  # Extra flags to provide to git when runnning git pull
  deploy:
    directory: "D:\\Code\\deploy"
    remote: "https://github.com/SpaceXLaunchBot/deploy"
    interval: 900  # How long to wait between checks and/or pulls
    method: "checkpull" # pull or checkpull, defaults to pull

```

## Build

Requires [govvv](https://github.com/ahmetb/govvv) to build correctly.

See `build.ps1` or the `Dockerfile` for build commands.
