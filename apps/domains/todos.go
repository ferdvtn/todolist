package domains

type Todos struct {
	ID          int
	Title       string
	Description string
	Priority    int
}

func (d Todos) IsValid() error {
	return nil
}
