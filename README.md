[![buddy pipeline](https://app.buddy.works/timfall/vesh/pipelines/pipeline/45662/badge.svg?token=3ae6c804af4fdb5947b58ba1c544c232bf8d28f6e6d2b07321added2d1cc0bad "buddy pipeline")](https://app.buddy.works/timfall/vesh/pipelines/pipeline/45662)
[![CircleCI](https://circleci.com/gh/vapor-ware/synse-cli.svg?style=shield&circle-token=7e11598b349e1d280c7cd78517ababef0f837bc3)](https://circleci.com/gh/vapor-ware/vesh)

```shell
NAME:
   vesh - Vapor Edge Shell

USAGE:
   vesh [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     status, stat   Get the status of the current deployment
     assets         Manage and get information about physical devices
     zones          List available zones
     racks          List available racks within a given `zone` (or all zones if none is specified)
     health         Check health for a given `zone`, `rack`, or `device`
     notifications  List notifications for a given `zone`, `rack`, or `device`
     load           Get the load by specific metric
     provision      Get (un)provisioned servers and provision new servers
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d             Enable debug mode
   --config file, -c file  Path to config file [$VESH_CONFIG_FILE]
   --host Vapor Host       Address of Vapor Host [$VESH_HOST]
   --help, -h              show help
   --version, -v           print the version
```

**TO DO**:
- [x] Error handling
- [x] Comments and documentation
- [x] Configuration from file
- [x] Debug mode
- [ ] "Second order" features
  - [ ] Notifications
  - [ ] Health
  - [ ] Provision
  - [ ] Load
- [ ] Colors?
