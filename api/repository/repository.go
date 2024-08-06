package repository

import (
	"grpc-prj/api/models"
	"log"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Insert(*models.Article) error
	Read(id uint) (*models.Article, error)
	Update(*models.Article) error
	Delete(id uint) error
}

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{DB: db}
}

func (rep *articleRepository) Insert(c *models.Article) error {
	if err := rep.DB.Create(c).Error; err != nil {
		log.Printf("Insert contact %+v err %v\n", c, err)
	}
	log.Printf("Insert %v successfully", c)
	return nil
}
func (rep *articleRepository) Read(id uint) (*models.Article, error) {
	var article models.Article
	if err := rep.DB.First(&article, id).Error; err != nil {
		log.Printf("Read article err %+v \n", err)
		return nil, err
	}
	return &article, nil
}
func (rep *articleRepository) Update(article *models.Article) error {
	err := rep.DB.Model(&models.Article{}).Where("id = ?", article.ID).Updates(article).Error
	if err != nil {
		return err
	}
	return nil
}
func (rep *articleRepository) Delete(id uint) error {
	if err := rep.DB.Delete(&models.Article{}, id).Error; err != nil {
		return err
	}
	return nil
}
