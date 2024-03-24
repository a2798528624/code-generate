package dao

import "go-code-generate/quick-start-demo/example/models"

type AuthorsDAO struct {
}

var authorsDAO = AuthorsDAO{}

func NewAuthorsDAO() *AuthorsDAO {
	return &authorsDAO
}

func (AuthorsDAO) CreateAuthors(info *models.Authors) error {
	return db.Create(info).Error
}

func (AuthorsDAO) CreateAllAuthors(infos *[]models.Authors) error {
	return db.Create(infos).Error
}

func (AuthorsDAO) QueryAuthors(cond, result *models.Authors) error {
	return db.Model(cond).First(result).Error
}

func (AuthorsDAO) QueryAllAuthors(cond, results *[]models.Authors) error {
	return db.Model(cond).Find(results).Error
}

func (AuthorsDAO) UpdateAuthors(old, new *models.Authors) error {
	return db.Model(old).Updates(*new).Error
}

func (AuthorsDAO) DeleteAuthors(cond *models.Authors) error {
	return db.Delete(cond).Error
}
