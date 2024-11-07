package vikingdb

type Task struct {
	CollectionName string
	CreateTime     string
	ProcessInfo    map[string]interface{}
	TaskID         string
	TaskParams     map[string]interface{}
	TaskStatus     string
	TaskType       string
	UpdatePerson   string
	UpdateTime     string
}
