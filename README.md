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
- [ ] Error handling
- [ ] Comments and documentation
- [ ] Configuration from file
- [ ] Debug mode
- [ ] "Second order" features
  - [ ] Notifications
  - [ ] Health
  - [ ] Provision
  - [ ] Load
