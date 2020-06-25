package repository

import (
	"../model"
	"github.com/jinzhu/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db : db}
}

func (r *UrlRepository) Save(url *model.Url) RepositoryResult {
	err := r.db.Save(url).Error

	if err != nil {
		return RepositoryResult{Error : err}
	}

	return RepositoryResult{Result : url}
}

func (r *UrlRepository) FindAll() RepositoryResult {
	var url model.Url

	err := r.db.Find(&url).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &url}
}

func (r *UrlRepository) FindOneById(slug string) RepositoryResult {
	var url model.Url

	err := r.db.Where(&model.Url{Slug: slug}).Take(&url).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &url}
}

func (r *UrlRepository) DeleteOneById(slug string) RepositoryResult {
	err := r.db.Delete(&model.Url{Slug: slug}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}