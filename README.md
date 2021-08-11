# CrowsNest

[Watchtower](https://github.com/containrrr/watchtower) for Git

## Configuration

### Flags

`--loop` or `-l`: Normally CrowsNest would run once then exit, set this flag to loop forever

`--method` or `-m`: Which method to use for checking / updating, should be `pull` or `checkpull`

`--config` or `-c`: Where to look for your config.yaml file (`.` and `$HOME` are automatically searched)

### config.yaml

Example:

```yaml
respositories:
  peicecost:
    directory: "D:\\Code\\piececost"
    remote: "https://github.com/psidex/PieceCost.git"
    gitflags: "-m me"  # Extra flags to provide to git when pulling
  deploy:
    directory: "D:\\Code\\deploy"
    remote: "https://github.com/SpaceXLaunchBot/deploy"
    interval: 900  # How long to wait between pulls when using --loop

```

## Build

Requires [govvv](https://github.com/ahmetb/govvv) to build correctly.

See `build.ps1` or the `Dockerfile` for build commands.
