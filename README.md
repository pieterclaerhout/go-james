# go-james

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-james)](https://goreportcard.com/report/github.com/pieterclaerhout/go-james) [![Documentation](https://godoc.org/github.com/pieterclaerhout/go-james?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-james) [![License](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-james/raw/master/LICENSE) [![GitHub Version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-james.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-james) [![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-james.svg)](https://github.com/pieterclaerhout/go-james/issues)

James is your butler and helps you to create, build, test and run your [Go](https://golang.org) projects.

When you often create new apps using [Go](https://golang.org), it quickly becomes annoying when you realize all the steps it takes to configure the basics. You need to manually create the source files, version info requires more steps to be injected into the executable, using [Visual Studio Code](https://code.visualstudio.com) requires you to manually setup the tasks you want to run…

Using the `go-james` tool, you can automate and streamline this process. The tool will take care of initializing your project, running your project, building it and running the tests.

<!-- TOC depthFrom:2 -->

- [Installation](#installation)
- [Creating a new project](#creating-a-new-project)
- [Initializing an existing project](#initializing-an-existing-project)
- [Building a project](#building-a-project)
- [Running a project](#running-a-project)
- [Testing a project](#testing-a-project)
- [The config file `go.json`](#the-config-file-gojson)
- [What is covered during `new` and `init`?](#what-is-covered-during-new-and-init)

<!-- /TOC -->

## Installation

As [Go 1.13](https://golang.org) is required for this tool, the best way to install this tool is to run `go install` as follows:

```
go install github.com/pieterclaerhout/go-james/cmd/go-james
```

This will create the `go-james` command in your `$GOPATH/bin` folder.

## Creating a new project

To create a new project, you can use the `new` subcommand as follows:

```
go-james new --path=<target-path> --package=<package> --name=<name> --description=<description>
```

You can specify the following options:

* `--path` (required): the path where the new project should be created, e.g. `/home/username/go-example`
* `--package` (required): the main package for the new project, e.g. `github.com/pieterclaerhout/go-example`
* `--name` (optional): the name of the project, if not specified, the last part of the path is used
* `--description` (optional): the description of the project, used for the readme

It will automatically create the following project structure:

```
/home/username/go-example
├── .gitignore
├── .vscode
│   └── tasks.json
├── LICENSE
├── README.md
├── build
│   └── go-example
├── cmd
│   └── go-example
│       └── main.go
├── go.json
├── go.mod
├── library.go
└── versioninfo
    └── versioninfo.go
```

## Initializing an existing project

When you already have an existing folder structure, you can run the `init` command to add the missing pieces.

```
go-james init
```

This command is supposed to run from the project's directory and doesn't take any arguments.

## Building a project

From within the project root, run the `build` command to build the executable:

```
go-james build [-v] [--output=<path>] [--goos=<os>] [--goarch=<arch>]
```

By default, the output is put in the `build` subdirectory.

By adding the `-v` flag, the packages which are built will be listed.

By adding the `--output` flag, you can override the default output path as specified in the config file.

By adding the `--goos` flag, you can override the `GOOS` env variable which indicates for which OS you are compiling.

By adding the `--goarch` flag, you can override the `GOARCH` env variable which indicates for which architecture you are compiling.

You can read more about these flags [here](https://golang.org/doc/install/source#environment).

## Running a project

From within the project root, run:

```
go-james run <args>
```

This will build the project and run it's main target passing the `<args>` to the command.

## Testing a project

From within the project root, run:

```
go-james test
```

This will run all the tests defined in the package.

## The config file `go.json`

```json
{
    "project": {
        "name": "go-james",
        "description": "James is your butler and helps you to create, build, test and run your Go projects",
        "package": "github.com/pieterclaerhout/go-james",
        "main_package": "github.com/pieterclaerhout/go-james/cmd/main",
        "repository": "github.com/pieterclaerhout/go-james"
    },
    "build": {
        "ouput_name": "build/go-james",
        "ld_flags": ["-s", "-w"],
        "extra_args": ["-trimpath"]
    }
}
```

## What is covered during `new` and `init`?

Already covered:

* README.md
* LICENSE (defaults to the Apache license)
* Visual Studio Code Tasks file
* Git Revision and Branch name injection in `versioninfo` package.
* `.gitignore`
* `clean`
* Tests extra flags
* go.json file
* `go mod init`
* Creation of the main entrypoint package
* Creation of the library
* `install`
* `uninstall`

Not covered yet:

* Creation of the git repo (optional)

Eventually:

* Running benchmarks
* `publish` (to github)
* `.dockerignore`
* `Dockerfile`
* Listing out-of-date dependencies
* Update of out-of-date dependencies
* Homebrew recipe
