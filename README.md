# Cross Check

Cross check makes health checks on PostgreSQL and MySQL database servers, it also performs master & slave control for clusters in H/A Active/Passive structure.

# Building

```
Build golang application
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cross-check

```

# Install & Update

Just run installation script. If already installed old version script will be switch upgrade mode and upgrade application.
```
$ git clone --single-branch --branch release https://github.com/ayetkin/cross-check.git
$ cd cross-check
$ chmod +x install.sh
$ ./install.sh [mysql|pgsql|update|remove]
```


# Usage
```
$ cross-check --help
This tool helps to making health check for pgsql and mysql servers.

Usage:
  cross-check [command]

Available Commands:
  help        Help about any command
  mysql       Switch master&slave check for MySQL nodes.
  pgsql       Switch master&slave check for PostgreSQL nodes.
  version     Print the version number of Cross Check

Flags:
  -h, --help   help for cross-check

Use "cross-check [command] --help" for more information about a command.

```

# [Changelog](CHANGELOG.md)



