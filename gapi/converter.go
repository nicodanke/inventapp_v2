package gapi

import (
	db "github.com/nicodanke/inventapp_v2/db/sqlc"
	"github.com/nicodanke/inventapp_v2/pb/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *models.User {
	return &models.User{
		Id:                user.ID,
		Username:          user.Username,
		Name:              user.Name,
		Lastname:          user.Lastname,
		Email:             user.Email,
		Phone:             user.Phone.String,
		Active:            user.Active,
		IsAdmin:           user.IsAdmin,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func convertAccount(user db.Account) *models.Account {
	return &models.Account{
		Id:          user.ID,
		Code:        user.Code,
		CompanyName: user.CompanyName,
		Phone:       user.Phone.String,
		Email:       user.Email,
		WebUrl:      user.WebUrl.String,
		Active:      user.Active,
		CreatedAt:   timestamppb.New(user.CreatedAt),
	}
}

func convertRole(role db.Role) *models.Role {
	return &models.Role{
		Id:                role.ID,
		Name:              role.Name,
	}
}
