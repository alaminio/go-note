package note

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNote(title, content string) (*Note, error) {
	note := &Note{
		Title:   title,
		Content: content,
	}
	err := s.repo.Create(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *Service) GetAllNotes() ([]Note, error) {
	return s.repo.GetAll()
}

func (s *Service) GetNoteByID(id uint) (*Note, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateNote(id uint, title, content string) (*Note, error) {
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	note.Title = title
	note.Content = content
	err = s.repo.Update(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *Service) DeleteNote(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) SearchNotes(title, content string) ([]Note, error) {
	return s.repo.SearchByTitleOrContent(title, content)
}
