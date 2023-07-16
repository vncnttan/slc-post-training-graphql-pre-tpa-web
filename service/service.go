package service

import (
	"context"

	"github.com/vncnttan/TrainingGraphQL/database"
	"github.com/vncnttan/TrainingGraphQL/graph/model"
	"gorm.io/gorm"
)

var db *gorm.DB = database.GetInstance()

func GetUser(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	return user, db.First(&user, "id = ?", id).Error
}
