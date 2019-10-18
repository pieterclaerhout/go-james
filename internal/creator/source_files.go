package creator

const mainLibTemplate = `package {{.ShortPackageName}}
`

const mainLibTestingTemplate = `package {{.ShortPackageName}}_test
`

const mainCmdTemplate = `package main

import(
	"fmt"

	"{{.Project.Package}}/versioninfo"
)

func main() {
	fmt.Println("Project: "+ versioninfo.ProjectName)
	fmt.Println("Description: "+ versioninfo.ProjectDescription)
	fmt.Println("Version: "+ versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)
}
`

const mainCmdTestingTemplate = `package main_test
`

const versionInfoTemplate = `package versioninfo

// ProjectName contains the name of the project
var ProjectName string

// ProjectDescription contains the description of the project
var ProjectDescription string

// Version contains the version of the app
var Version string

// Revision will be injected with the current commit hash
var Revision string

// Branch will be injected with the current branch name
var Branch string
`

const postBuildScript = `package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Running post buid: %v\n", os.Args)
}
`
