package domains

type Todolist struct {
	ID          int
	Title       string
	Description string
	Priority    int
}

func (d Todolist) IsValid() error {
	return nil
}
