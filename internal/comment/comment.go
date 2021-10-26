package comment

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type CommentService interface {
	GetCommentById(id uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(data Comment) (Comment, error)
	DeleteComment(id uint) error
	UpdateComment(id uint, data Comment)
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}
