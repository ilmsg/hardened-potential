package api

import (
	"github.com/ilmsg/hardened-potential/model"
	"github.com/ilmsg/hardened-potential/slug"
	"gorm.io/gorm"
)

type IUserHandler interface {
	NewUser(name string) (*model.User, error)
	FindAll() ([]model.User, error)
}

type UserHandler struct {
	db *gorm.DB
}

func (h *UserHandler) FindAll() ([]model.User, error) {
	var user []model.User
	if tx := h.db.Find(&user); tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (h *UserHandler) NewUser(name string) (*model.User, error) {
	newUser := model.User{
		ID:   slug.GenerateSlug(),
		Name: name,
	}
	if tx := h.db.Create(&newUser); tx.Error != nil {
		return nil, tx.Error
	}

	var user model.User
	if tx := h.db.First(&user, "id = ?", newUser.ID); tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func NewUserHandler(db *gorm.DB) IUserHandler {
	return &UserHandler{db}
}
