<p align="center"><img align="center" src="https://github.com/pieterclaerhout/go-james/raw/master/resources/gopher-james.png" width="120"/></p>

<h1 align="center">go-james</h1>

<p align="center"><a href="https://goreportcard.com/report/github.com/pieterclaerhout/go-james" rel="nofollow"><img src="https://camo.githubusercontent.com/4870b0a18ef6cc3f3d10e5c468a8da7f17aee89d/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f706965746572636c616572686f75742f676f2d6a616d6573" alt="Go Report Card" data-canonical-src="https://goreportcard.com/badge/github.com/pieterclaerhout/go-james" style="max-width:100%;"></a> <a href="http://godoc.org/github.com/pieterclaerhout/go-james" rel="nofollow"><img src="https://camo.githubusercontent.com/c1cccc4298cc381799b2ac901183e30d76f24033/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f706965746572636c616572686f75742f676f2d6a616d65733f7374617475732e737667" alt="Documentation" data-canonical-src="https://godoc.org/github.com/pieterclaerhout/go-james?status.svg" style="max-width:100%;"></a> <a href="https://github.com/pieterclaerhout/go-james/raw/master/LICENSE"><img src="https://camo.githubusercontent.com/b65a7a2b7a579e10dad81e1ef2ef8945c8970806/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d41706163686525323076322d6f72616e67652e737667" alt="License" data-canonical-src="https://img.shields.io/badge/license-Apache%20v2-orange.svg" style="max-width:100%;"></a> <a href="https://github.com/pieterclaerhout/go-james/releases"><img src="https://camo.githubusercontent.com/113e1eb8bce0a2c952c6d12bf679b2ada756a4f5/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f762f72656c656173652f706965746572636c616572686f75742f676f2d6a616d6573" alt="GitHub Version" data-canonical-src="https://img.shields.io/github/v/release/pieterclaerhout/go-james" style="max-width:100%;"></a> <a href="https://github.com/pieterclaerhout/go-james/issues"><img src="https://camo.githubusercontent.com/e53684485fb0405984e3c51d9fcfcd895a2340ce/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6973737565732f706965746572636c616572686f75742f676f2d6a616d65732e737667" alt="GitHub issues" data-canonical-src="https://img.shields.io/github/issues/pieterclaerhout/go-james.svg" style="max-width:100%;"></a>  <a href="https://github.com/pieterclaerhout/go-james/"><img src="https://camo.githubusercontent.com/3fe96add896b4c526262790e6ba86472517b6993/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6173742d636f6d6d69742f706965746572636c616572686f75742f676f2d6a616d65732e737667" alt="GitHub Last Commit" data-canonical-src="https://img.shields.io/github/last-commit/pieterclaerhout/go-james.svg" style="max-width:100%;"></a></p>

James is your butler and helps you to create, build, debug, test and run your [Go](https://golang.org) projects.

When you often create new apps using [Go](https://golang.org), it quickly becomes annoying when you realize all the steps it takes to configure the basics. You need to manually create the source files, version info requires more steps to be injected into the executable, using [Visual Studio Code](https://code.visualstudio.com) requires you to manually setup the tasks you want to run…

Using the `go-james` tool, you can automate and streamline this process. The tool will take care of initializing your project, running your project, debugging it, building it and running the tests.

You should be using `go-james` if:

* you're tired of setting up your projects manually
* you don't want to have to specify the main package every time you want build or run your project
* you keep on forgetting how to setup debugging [Visual Studio Code](https://code.visualstudio.com)
* you don't want to setup your tasks file manually in [Visual Studio Code](https://code.visualstudio.com)
* you want to have the Git revision, branch name and version number automatically in your project
* you're tired of manually writing Makefiles to build, test and run your project
* you want a better way to do tasks before and after the build than writing non-portable shell scripts
* you want an easy way to cross-compile for all common GOOS/GOARCH combinations
* you want a way to build, run and test your project in a cross-platform manner

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
- [Bootstrapping `go-james`](#bootstrapping-go-james)
- [Roadmap](#roadmap)
- [Resources](#resources)

<!-- /TOC -->

---

## Requirements

* [Go](https://golang.org) 1.14 or newer
* [Go Modules](https://github.com/golang/go/wiki/Modules) (the de-facto standard)

## Installation

### Using `go get`

You can run the following command to install `go-james`:

```
go get -u github.com/pieterclaerhout/go-james/cmd/go-james
```

This will create the `go-james` command in your `$GOPATH/bin` folder.

The tool is self-contained and doesn't have any external dependencies.

### Homebrew

To install via [homebrew](https://brew.sh), run the following commands:

```
$ brew tap pieterclaerhout/go-james
==> Tapping pieterclaerhout/go-james
Cloning into '/usr/local/Homebrew/Library/Taps/pieterclaerhout/homebrew-go-james'...
remote: Enumerating objects: 4, done.
remote: Counting objects: 100% (4/4), done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 4 (delta 0), reused 4 (delta 0), pack-reused 0
Receiving objects: 100% (4/4), done.
Tapped 1 formula (27 files, 26.5KB).

$ brew install go-james
==> Installing go-james from pieterclaerhout/go-james
==> Downloading https://github.com/pieterclaerhout/go-james/releases/download/v1.6.0/go-james_darwin_amd64.tar.gz
######################################################################## 100.0%
🍺  /usr/local/Cellar/go-james/1.6.0: 4 files, 11.5MB, built in 3 seconds
```

### Manually

To install it manually, download the `go-james` executable from  the [releases](https://github.com/pieterclaerhout/go-james/releases) and place it in `$GOPATH/bin`.

## Starting a new project

To start a new project, you can use the `new` subcommand as follows:

```
go-james new --path=<target path> \
             --package=<package> \
             --name=<name of your project> \
             --description=<description of your project> \
             --copyright=<copyright of your project> \
             [--with-git] \
             [--with-docker] \
             [--with-github-action] \
             [--with-gitlab-ci] \
             [--overwrite]
```

When you run it, you'll get the following output:

```
➜ go-james new --path go-example --package github.com/pieterclaerhout/go-example
Creating package: github.com/pieterclaerhout/go-example
Project path: /Users/pclaerhout/Desktop/go-example
Writing: go-example
Writing: go-example/go-james.json
Writing: go-example/.vscode
Writing: go-example/.vscode/tasks.json
Writing: go-example/.vscode/settings.json
Writing: go-example/.vscode/launch.json
Writing: go-example/LICENSE
Writing: go-example/.gitignore
Writing: go-example/.dockerignore
Writing: go-example/Dockerfile
Writing: go-example/README.md
Writing: go-example/scripts/post_build
Writing: go-example/scripts/post_build/post_build.example.go
Writing: go-example/library.go
Writing: go-example/library_test.go
Writing: go-example/cmd/go-example
Writing: go-example/cmd/go-example/main.go
Writing: go-example/versioninfo
Writing: go-example/versioninfo/versioninfo.go
Writing: go-example/scripts/pre_build
Writing: go-example/scripts/pre_build/pre_build.example.go
Writing: go-example/go.mod
```

It will automatically create the following folder and file structure:

```
go-example
.
├── Dockerfile
├── Downloads
├── DropBox
├── Twixl\ Publisher
└── go-example
    ├── Dockerfile
    ├── LICENSE
    ├── README.md
    ├── cmd
    │   └── go-example
    │       └── main.go
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
* `--with-git`: if specified, a local Git repository will be created for the project and the source files will automatically be committed.
* `--with-docker`: if specified, a sample Dockerfile and .dockerignore file will be created.
* `--with-github-action`: if specified, a sample Github Actions file will be created.
* `--with-gitlab-ci`: if specified, a sample Gitlab-CI file will be created.
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

Just before the build, if a file called `<project_root>/scripts/pre_build/pre_build.go` is present, it will be executed and will get a lot of info about the build injected. It's a plain Go file, so use whatever trick or tool you know. A sample pre-build script looks as follows:

```go
package main

import (
	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	args, err := james.ParseBuildArgs()
	log.CheckError(err)

	log.InfoDump(args, "pre_build arguments")

}
```

You can also execute a script after the build. To do so, create a file `<project_root>/scripts/post_build/post_build.go` with contents similar to:

```go
package main

import (
	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	args, err := james.ParseBuildArgs()
	log.CheckError(err)

	log.InfoDump(args, "post_build arguments")

}
```

### Parsing arguments in pre/post build scripts.

To parse the arguments, you can use [`james.ParseBuildArgs()`](https://godoc.org/github.com/pieterclaerhout/go-james#ParseBuildArgs).

The parameters it gets are are struct of the type [`james.BuildArgs`](https://godoc.org/github.com/pieterclaerhout/go-james#BuildArgs):

```
james.BuildArgs{
  ProjectPath: "/home/user/go-james",
  OutputPath: "/home/user/go-james/build/go-james",
  GOOS: "darwin",
  GOARCH: "amd64",
  ProjectName: "go-james",
  ProjectDescription: "James is your butler and helps you to create, build, test and run your Go projects",
  ProjectCopyright: "© 2019-2020 Copyright Pieter Claerhout",
  Version: "0.7.0",
  Revision: "2065b13",
  Branch: "master",
  RawBuildCommand: []string{
    "go",
    "build",
    "-o",
    "build/go-james",
    "-ldflags",
    "-s -w -X github.com/pieterclaerhout/go-james/versioninfo.ProjectName=go-james -X 'github.com/pieterclaerhout/go-james/versioninfo.ProjectDescription=James is your butler and helps you to create, build, test and run your Go projects' -X 'github.com/pieterclaerhout/go-james/versioninfo.ProjectCopyright=© 2019 Copyright Pieter Claerhout' -X github.com/pieterclaerhout/go-james/versioninfo.Version=0.7.0 -X github.com/pieterclaerhout/go-james/versioninfo.Revision=2065b13 -X github.com/pieterclaerhout/go-james/versioninfo.Branch=master",
    "-trimpath",
    "github.com/pieterclaerhout/go-james/cmd/go-james",
  },
}
```

The same information is also available in the following environment variables:

```
GO_JAMES_BRANCH="master"
GO_JAMES_GOARCH="amd64"
GO_JAMES_GOOS="darwin"
GO_JAMES_OUTPUT_PATH="/home/user/go-james/build/go-james"
GO_JAMES_PROJECT_COPYRIGHT="© 2019-2020 Copyright Pieter Claerhout"
GO_JAMES_PROJECT_DESCRIPTION="James is your butler and helps you to create, build, test and run your Go projects"
GO_JAMES_PROJECT_NAME="go-james"
GO_JAMES_PROJECT_PATH="/home/user/go-james"
GO_JAMES_REVISION="2065b13"
GO_JAMES_VERSION="0.7.0"
```

### Using Bash scripts for pre/post build actions.

If you prefer to use a bash script instead for the pre/post build actions, you can create a file called:

```
<project_root>/scripts/post_build/pre_build.sh
```

or

```
<project_root>/scripts/post_build/post_build.sh
```

It should be marked as executable.


### Using Batch files for pre/post build actions.

If you prefer to use a batch file on Windows instead for the pre/post build actions, you can create a file called:

```
<project_root>/scripts/post_build/pre_build.bat
```

or

```
<project_root>/scripts/post_build/post_build.bat
```

### How go-james defines which pre/post build script to run

go-james will search for pre/post build scripts in the following order:

* `pre_build.go` / `post_build.go`
* `pre_build.sh` / `post_build.sh`
* `pre_build.bat` / `post_build.bat`

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

## Running a static analysis

You can use the `staticcheck` command to run the [staticcheck](https://staticcheck.io/docs/) static analyzer. The binary required to run staticcheck is automatically installed if needed.

```
go-james staticcheck
```

### Building a Docker image

You can use the `docker-image` command to build a Docker image using the Dockerfile in the project folder. When you create a new project, a starter Dockerfile will be created automatically.

```
go-james docker-container
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
        "output_path": "build/go-example",
        "ld_flags": [
            "-s",
            "-w"
        ],
        "ld_flags_windows": [
            "-s",
            "-w",
            "-H",
            "windowsgui"
        ],
        "ld_flags_darwin": [],
        "ld_flags_linux": [],
        "extra_args": [
            "-trimpath"
        ],
        "use_gotip": false
    },
    "run": {
        "environ": {
            "var": "val"
        }
    },
    "package": {
        "include_readme": true
    },
    "test": {
        "extra_args": []
    },
    "staticcheck": {
        "checks": ["all", "-ST1005", "-ST1000"]
    },
    "docker-image": {
        "name": "go-james",
        "repository": "pieterclaerhout/go-james",
        "tag": "revision",
        "prune_images_after_build": false
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
* `ld_flags`: the linker flags you want to use for building. You can find more info about these flags [here](https://golang.org/cmd/link/). These are only used if you don't specify specific parameters for a specifc `GOOS`.
* `ld_flags_darwin`: the linker flags you want to use for building `darwin`. You can find more info about these flags [here](https://golang.org/cmd/link/).
* `ld_flags_linux`: the linker flags you want to use for building for `linux`. You can find more info about these flags [here](https://golang.org/cmd/link/).
* `ld_flags_windows`: the linker flags you want to use for building for `windows`. You can find more info about these flags [here](https://golang.org/cmd/link/).
* `extra_args`: contains any extra command-line parameters you want to add to the `go build` command when you run `go-james build`.
* `use_gotip`: setting this to true uses `gotip` to compile instead of the regular `go`command. Make sure you have [`gotip`](https://pkg.go.dev/golang.org/dl/gotip?tab=doc) installed.

### Run Config

* `environ`: the environment variables to use when running the app

### Package Config

* `include_readme`: boolean indicating if the README.md file should be included in the package or not

### Test Config

* `extra_args`: contains any extra command-line parameters you want to add to the `go test` command when you run `go-james test`.

### Staticcheck Config

* `checks`: the checks for [staticcheck](https://staticcheck.io/docs/) you want to run

### Docker Image Config

* `name`: the name of the docker image you want to create. Defaults to the project name.
* `repository`: the repository to which you want to push the image. If left empty, the image will only be created locally.
* `tag`: can be either `revision` or `version` (the default) and indicates what value should be used for the tag.
* `prune_images_after_build`: if set to true, a `docker image prune -f` will be executed after the docker build step.

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