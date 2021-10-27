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
	CreateComment(data Comment) (Comment, error)
	DeleteComment(id uint) error
	UpdateComment(id uint, data Comment) (Comment, error)
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetComment(id uint) (Comment, error) {
	var comment Comment
	if result := s.db.First(&comment, id); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.db.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

func (s *Service) CreateComment(newComment Comment) (Comment, error) {
	if result := s.db.Save(&newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return newComment, nil
}

func (s *Service) UpdateComment(id uint, commentData Comment) (Comment, error) {
	comment, err := s.GetComment(id)
	if err != nil {
		return Comment{}, err
	}
	if result := s.db.Model(&comment).Updates(commentData); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) DeleteComment(id uint) error {
	if result := s.db.Delete(Comment{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.db.Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}
