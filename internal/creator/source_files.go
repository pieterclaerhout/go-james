package creator

const mainLibTemplate = `package {{.Project.ShortPackageName}}
`

const mainLibTestingTemplate = `package {{.Project.ShortPackageName}}_test
`

const mainCmdTemplate = `package main

import(
	"fmt"

	"{{.Project.Package}}/versioninfo"
)

func main() {
	fmt.Println("Project: "+ versioninfo.ProjectName)
	fmt.Println("Description: "+ versioninfo.ProjectDescription)
	fmt.Println("Copyright: "+ versioninfo.ProjectCopyright)
	fmt.Println("Version: "+ versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)
}
`

const mainCmdTestingTemplate = `package main_test
`

const versionInfoTemplate = `package versioninfo

// ProjectName contains the name of the project
var ProjectName string = {{quote .Config.Project.Name}}

// ProjectDescription contains the description of the project
var ProjectDescription string = {{quote .Config.Project.Name}}

// ProjectCopyright contains the copyright for the project
var ProjectCopyright string = ""

// Version contains the version of the app
var Version string = {{quote .Config.Project.Version}}

// Revision will be injected with the current commit hash
var Revision string = "head"

// Branch will be injected with the current branch name
var Branch string = "master"
`

const preBuildScript = `package main

import (
	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	args, err := james.ParseBuildArgs()
	log.CheckError(err)

	log.InfoDump(args, "pre_build arguments")

}
`

const postBuildScript = `package main

import (
	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	args, err := james.ParseBuildArgs()
	log.CheckError(err)

	log.InfoDump(args, "post_build arguments")

}
`
