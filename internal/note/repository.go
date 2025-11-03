package note

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(note *Note) error {
	return r.db.Create(note).Error
}

func (r *Repository) GetAll() ([]Note, error) {
	var notes []Note
	err := r.db.Find(&notes).Error
	return notes, err
}

func (r *Repository) GetByID(id uint) (*Note, error) {
	var note Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *Repository) Update(note *Note) error {
	return r.db.Save(note).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Note{}, id).Error
}

func (r *Repository) SearchByTitleOrContent(title, content string) ([]Note, error) {
	var notes []Note
	query := r.db

	if title != "" && content != "" {
		// Both title and content provided - OR condition
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+title+"%", "%"+content+"%")
	} else if title != "" {
		// Only title provided
		query = query.Where("title LIKE ?", "%"+title+"%")
	} else if content != "" {
		// Only content provided
		query = query.Where("content LIKE ?", "%"+content+"%")
	} else {
		// No search parameters - return all notes
		return r.GetAll()
	}

	err := query.Find(&notes).Error
	return notes, err
}
