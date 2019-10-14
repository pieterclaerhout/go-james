package creator

const visualStudioDirName = ".vscode"
const visualStudioCodeTasksFileName = "tasks.json"

type visualStudioCodeTasks struct {
	Version string
	Tasks   []visualStudioCodeTask `json:"tasks"`
}

type visualStudioCodeTask struct {
	Label          string   `json:"label"`
	Command        string   `json:"command"`
	Type           string   `json:"type"`
	Group          string   `json:"group"`
	ProblemMatcher []string `json:"problemMatcher"`
}

func newVisualStudioCodeTaskList(tasks ...visualStudioCodeTask) *visualStudioCodeTasks {
	result := &visualStudioCodeTasks{
		Version: "2.0.0",
		Tasks:   []visualStudioCodeTask{},
	}
	for _, task := range tasks {
		result.Add(task)
	}
	return result
}

func (tasks *visualStudioCodeTasks) Add(task visualStudioCodeTask) {
	if task.Type == "" {
		task.Type = "shell"
	}
	if task.Group == "" {
		task.Group = "build"
	}
	tasks.Tasks = append(tasks.Tasks, task)
}
