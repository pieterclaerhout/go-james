<p align="center"><img align="center" src="https://github.com/pieterclaerhout/go-james/raw/master/resources/gopher-james.png" width="120"/></p>

<h1 align="center">go-james</h1>

<p align="center"><b>This project is stil in alpha, use at your own risk</b></p>

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-james)](https://goreportcard.com/report/github.com/pieterclaerhout/go-james) [![Documentation](https://godoc.org/github.com/pieterclaerhout/go-james?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-james) [![License](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-james/raw/master/LICENSE) [![GitHub Version](https://img.shields.io/github/v/release/pieterclaerhout/go-james)](https://github.com/pieterclaerhout/go-james/releases) [![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-james.svg)](https://github.com/pieterclaerhout/go-james/issues)  [![GitHub Last Commit](https://img.shields.io/github/last-commit/pieterclaerhout/go-james.svg)](https://github.com/pieterclaerhout/go-james/)

James is your butler and helps you to create, build, debug, test and run your [Go](https://golang.org) projects.

When you often create new apps using [Go](https://golang.org), it quickly becomes annoying when you realize all the steps it takes to configure the basics. You need to manually create the source files, version info requires more steps to be injected into the executable, using [Visual Studio Code](https://code.visualstudio.com) requires you to manually setup the tasks you want to run…

Using the `go-james` tool, you can automate and streamline this process. The tool will take care of initializing your project, running your project, debugging it, building it and running the tests.

---

<!-- TOC depthFrom:2 -->

- [Requirements](#requirements)
- [Installation](#installation)
- [Starting a new project](#starting-a-new-project)
- [Initializing an existing project](#initializing-an-existing-project)
- [Building a project](#building-a-project)
- [Pre-build and post-build hooks](#pre-build-and-post-build-hooks)
- [Packaging a project](#packaging-a-project)
- [Debugging a project](#debugging-a-project)
- [Running a project](#running-a-project)
- [Testing a project](#testing-a-project)
- [Installing the executable](#installing-the-executable)
- [Uninstalling the executable](#uninstalling-the-executable)
- [The config file `go-james.json`](#the-config-file-go-jamesjson)
    - [Project Config](#project-config)
    - [Build Config](#build-config)
    - [Test Config](#test-config)
- [Bootstrapping `go-james`](#bootstrapping-go-james)
- [Roadmap](#roadmap)
- [Resources](#resources)

<!-- /TOC -->

---
 
## Requirements

* [Go](https://golang.org) 1.13 or newer
* [Go Modules](https://github.com/golang/go/wiki/Modules) (the de-facto standard)

## Installation

The best way to install this tool is to run `go install` as follows:

```
go get -u github.com/pieterclaerhout/go-james/cmd/go-james
```

This will create the `go-james` command in your `$GOPATH/bin` folder. The tool is self-contained and doesn't have any external dependencies.

## Starting a new project

To start a new project, you can use the `new` subcommand as follows:

```
go-james new --path=<target path> \
             --package=<package> \
             --name=<name of your project> \
             --description=<description of your project> \
             --copyright=<copyright of your project>
             [--create-git-repo] \
             [--overwrite] \
```

When you run it, you'll get the following output:

```
➜ go-james new --path go-example --package github.com/pieterclaerhout/go-example
Creating: go-example/go-james.json
Creating: go-example/.vscode/launch.json
Creating: go-example/.vscode/tasks.json
Creating: go-example/LICENSE
Creating: go-example/.gitignore
Creating: go-example/README.md
Creating: go-example/library.go
Creating: go-example/cmd/go-example/main.go
Creating: go-example/versioninfo/versioninfo.go
go: creating new go.mod: module github.com/pieterclaerhout/go-example
```

It will automatically create the following folder and file structure:

```
go-example
├── .git
├── .gitignore
├── .vscode
│   ├── launch.json
│   └── tasks.json
├── LICENSE
├── README.md
├── cmd
│   └── go-example
│       ├── main.go
│       └── main_test.go
├── go-james.json
├── go.mod
├── library.go
├── library_test.go
├── scripts
│   ├── post_build
│   │   └── post_build.example.go
│   └── pre_build
│       └── pre_build.example.go
└── versioninfo
    └── versioninfo.go
```

An important file which is generated and can be used to further customize the project and it's settings is the [`go-james.json`](#the-config-file-go-jamesjson) file which sits next to the `go.mod` file.

You can specify the following options:

* `--path`: the path where the new project should be created, e.g. `/home/username/go-example` (if not specified, it will create a directory with the name of the prject in the current path)
* `--package`: the main package for the new project, e.g. `github.com/pieterclaerhout/go-example` (defaults to the project name if specified)
* `--name`: the name of the project, if not specified, the last part of the path is used
* `--description`: the description of the project, used for the readme
* `--copyright`: the copyright of the project, used for the readme
* `--create-git-repo`: if specified, a local Git repository will be created for the project and the source files will automatically be committed.
* `--overwrite`: if the destination path already exists, overwrite it (be careful, the original folder will be replaced)

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

By default, the output is put in the `build` subdirectory but can be customized in [the configuration file]((#the-config-file-go-jamesjson)).

You can specify the following options:

* `-v`: the packages which are built will be listed.
* `--output`: you can override the default output path as specified in [the configuration file]((#the-config-file-go-jamesjson)).
* `--goos`: you can override the `GOOS` environment variable which indicates for which OS you are compiling.
* `--goarch`: you can override the `GOARCH` environment variable which indicates for which architecture you are compiling.

You can read more about the `GOOS` and `GOARCH` environment variables [here](https://golang.org/doc/install/source#environment).

As part of the build process, the `versioninfo` package will be filled with the following details:

* `versioninfo.ProjectName`: the name of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.ProjectDescription`: the description of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.ProjectCopyright`: the copyright of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.Version`: the version of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.Revision`: the current Git commit hash
* `versioninfo.Branch`: the current Git branch name

With every build, these variables are automatically updated.

## Pre-build and post-build hooks

_TODO: needs to be written…_

## Packaging a project

From within the project root, run the `package` command to build the executable for windows / darwin / linux in the 386 and amd64 variants and compresses the result as a `.zip` (windows) or `.tgz` (linux / mac):

```
go-james package [-v] [--concurrency=4]
```

By default, the output is put in the `build` subdirectory but can be customized in [the configuration file]((#the-config-file-go-jamesjson)).

The filenames which are constructed use the following convention:

```
build/<project.name>_<goos>-<goarch>_v<project.version>.[zip,tgz]
```

The executable will be compressed and, if present in the project, the project's `README.md` file as well.

You can specify the following options:

* `-v`: the packages which are built will be listed.
* `--concurrency`: how many package processes should run in parallel, defaults to the number of CPUs.

As part of the build process, the `versioninfo` package will be filled with the following details:

* `versioninfo.ProjectName`: the name of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.ProjectDescription`: the description of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.Version`: the version of the project from [the configuration file]((#the-config-file-go-jamesjson))
* `versioninfo.Revision`: the current Git commit hash
* `versioninfo.Branch`: the current Git branch name

With every build, these variables are automatically updated.

## Debugging a project

From within the project root, run:

```
go-james debug
```

This will build the project and run it's main target through the [Delve debugger](https://github.com/go-delve). If the `dlv` command is not yet present in your `$GOPATH/bin` folder, it will automaticall be installed the first time you run it.

When creating a new project or performing `init` on an existing project, it also configures debugging from within [Visual Studio Code](https://code.visualstudio.com). It's a simple as setting one or more breakpoints and choose "Start" > "Debug" from the menu. It creates a launch configuration called `Debug`.

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

## Installing the executable

To install the main executable of your project in `$GOPATH/bin`, simply run the `install` command.

This will build the project and install it in the `$GOPATH/bin` folder. The name of the executable is the basename of build output path (as specified in [the configuration file]((#the-config-file-go-jamesjson)).

```
go-james uninstall
```

## Uninstalling the executable

Similar to the `install` command, there is also an `uninstall` command which removes the executable from `$GOPATH/bin`.

```
go-james uninstall
```

## The config file `go-james.json`

When you create a new project or init an existing one, a `go-james.json` file will be created in the root of your project. This file can be used to configure the project. The full config file is as follows:

```json
{
    "project": {
        "name": "go-example",
        "version": "1.0.0",
        "description": "",
        "copyright": "",
        "package": "github.com/pieterclaerhout/go-example",
        "main_package": "github.com/pieterclaerhout/go-example/cmd/go-example"
    },
    "build": {
        "ouput_path": "build/go-example",
        "ld_flags": [
            "-s",
            "-w"
        ],
        "extra_args": [
            "-trimpath"
        ]
    },
    "test": {
        "extra_args": []
    }
}
```

### Project Config

* `name`: the name of your project (will be availabme under `<package>/versioninfo.ProjectName`)
* `version`: the version of your project (will be availabme under `<package>/versioninfo.Version`)
* `description`: the description of your project (will be availabme under `<package>/versioninfo.ProjectDescription`)
* `copyright`: the description of your project (will be availabme under `<package>/versioninfo.ProjectCopyright`)
* `package`: the root package of your project
* `main_package`: the full path to the main package of your app, defaults to `<package>/cmd/<project-name>`

### Build Config

* `output_path`: the path where the built executable should be placed. Defaults to `build/<project-name>`
* `ld_flags`: the linker flags you want to use for building. You can find more info about these flags [here](https://golang.org/cmd/link/).
* `extra_args`: contains any extra command-line parameters you want to add to the `go build` command when you run `go-james build`.

### Test Config

* `extra_args`: contains any extra command-line parameters you want to add to the `go test` command when you run `go-james test`.

## Bootstrapping `go-james`

If you want to build `go-james` from scratch, you can use the following command (or use the "bootstrap" build task in Visual Studio Code):

```
go build -v -o build/go-james github.com/pieterclaerhout/go-james/cmd/go-james
```

If you have a version of `go-james` installed, you can use it to build itself.

## Roadmap

To get an idea on what's coming, you can check the [GitHub Milestones](https://github.com/pieterclaerhout/go-james/milestones).

## Resources

Follow my [weblog about Go & Kubernetes](https://www.yellowduck.be) :-)