# dabba

daaba is a dummy container platform built to understand *namespaces* and *cgroups* in linux.

## Directory Structure

This repository follows the standard golang project structure.

[Read more](https://github.com/golang-standards/project-layout)

## Build

```shell
$ ./build.sh
```

The executable binaries are put into the `bin` directory.

## Setup

Setting up the environment prior to running is important.

```shell
$ ./bin/dabba setup
```

This sets up the `alpine` root filesystem under `fs/alpine` directory.

## Run

```shell
$ ./bin/dabba run [command]
```

Eg:
To start a shell

```shell
$ ./bin/dabba run /bin/sh
```
