package james

type Project struct {
	Path string
}

func NewProject(path string) Project {
	return Project{
		Path: path,
	}
}
