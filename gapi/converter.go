package gapi

import (
	db "github.com/nicodanke/bankTutorial/db/sqlc"
	"github.com/nicodanke/bankTutorial/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
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

func convertAccount(user db.Account) *pb.Account {
	return &pb.Account{
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