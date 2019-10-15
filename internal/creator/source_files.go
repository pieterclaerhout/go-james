package creator

const mainLibTemplate = `package {{.ShortPackageName}}
`

const mainCmdTemplate = `package {{.ShortPackageName}}

import(
	"fmt"
)

func main() {
	fmt.Println("{{.Project.Name}}")
}
`
