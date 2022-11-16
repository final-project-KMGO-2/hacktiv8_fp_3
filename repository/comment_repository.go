package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	GetComment(ctx context.Context, userID uint64) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error)
	UpdateCommentByID(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	DeleteCommentByID(ctx context.Context, commentID uint64) error
}

type CommentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentConnection{
		connection: db,
	}
}

// CreateComment implements CommentRepository
func (db *CommentConnection) CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	tx := db.connection.Create(&comment)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}
	return comment, nil
}

// GetComment implements CommentRepository
func (db *CommentConnection) GetComment(ctx context.Context, userID uint64) ([]entity.Comment, error) {
	var comment []entity.Comment
	tx := db.connection.Debug().Where(("user_id = ?"), userID).Preload("User").Preload("Photo").Find(&comment)
	if tx.Error != nil {
		return []entity.Comment{}, tx.Error
	}
	return comment, nil

}

// GetCommentByID implements CommentRepository
func (db *CommentConnection) GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error) {
	var comment entity.Comment
	tx := db.connection.Where(("id = ?"), commentID).Take(&comment)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	return comment, nil
}

// UpdateCommentByID implements CommentRepository
func (db *CommentConnection) UpdateCommentByID(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	tx := db.connection.Model(&entity.Comment{}).Where(("id = ?"), comment.ID).Update("message", comment.Message)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}
	return comment, nil
}

// DeleteCommentByID implements CommentRepository
func (db *CommentConnection) DeleteCommentByID(ctx context.Context, commentID uint64) error {
	var comment entity.Comment
	tx := db.connection.Where(("id = ?"), commentID).Delete(&comment)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
