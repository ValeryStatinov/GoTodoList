package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    uint8  `json:"priority"`
	ProjectId   int    `json:"projectId,omitempty"`
	Completed   bool   `json:"completed"`
}

func (t *Task) Validate() bool {
	if t.Priority < 1 || t.Priority > 3 {
		return false
	}

	if t.Name == "" {
		return false
	}

	return true
}
