# go-james

James is your butler and helps you to create, build, test and run your Go projects

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-james)](https://goreportcard.com/report/github.com/pieterclaerhout/go-james) [![Documentation](https://godoc.org/github.com/pieterclaerhout/go-james?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-james) [![License](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-james/raw/master/LICENSE) [![GitHub Version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-james.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-james) [![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-james.svg)](https://github.com/pieterclaerhout/go-james/issues)

## Installation

Via `go install`:

```
go install github.com/pieterclaerhout/go-james/cmd/go-james
```

Via [Homebrew](http://brew.sh):

```
brew install go-james
```

## Creating a new project

```
go-james new
```

## Initializing an existing project

```
go-james init
```

## Building a project

From within the project root, run:

```
go-james build [-v]
```

By default, the output is put in the `build` subdirectory.

By adding the `-v` flag, the packages which are built will be listed.

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

Not covered yet:

* `go mod init`
* Tests extra flags
* Creation of the main entrypoint package
* Creation of the library
* Running benchmarks
* Creation of the git repo (optional)
* `clean`
* `install`
* `uninstall`
* `publish` (to github)

Eventually:

* `.dockerignore`
* `Dockerfile`
* Listing out-of-date dependencies
* Update of out-of-date dependencies
