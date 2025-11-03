package todo

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTodo(title, description string) (*Todo, error) {
	todo := &Todo{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *Service) GetAllTodos() ([]Todo, error) {
	return s.repo.GetAll()
}

func (s *Service) GetTodoByID(id uint) (*Todo, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateTodo(id uint, title, description string) (*Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Description = description
	err = s.repo.Update(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *Service) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) ToggleTodoComplete(id uint) error {
	return s.repo.ToggleComplete(id)
}

func (s *Service) GetCompletedTodos() ([]Todo, error) {
	return s.repo.GetCompleted()
}

func (s *Service) GetPendingTodos() ([]Todo, error) {
	return s.repo.GetPending()
}

func (s *Service) GetTodosByTitle(text string) ([]Todo, error) {
	return s.repo.GetByTitleContaining(text)
}
