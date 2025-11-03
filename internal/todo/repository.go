package todo

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(todo *Todo) error {
	return r.db.Create(todo).Error
}

func (r *Repository) GetAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *Repository) GetByID(id uint) (*Todo, error) {
	var todo Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *Repository) Update(todo *Todo) error {
	return r.db.Save(todo).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Todo{}, id).Error
}

func (r *Repository) ToggleComplete(id uint) error {
	var todo Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return err
	}
	todo.Completed = !todo.Completed
	return r.db.Save(&todo).Error
}
