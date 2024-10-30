package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"auth/internal/model"
	desc "auth/pkg/user_v1"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:   info.Name,
		Email: info.Email,
		Role: desc.Role(info.Role),
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:   info.Name,
		Email: info.Email,
		Role: model.Role(info.Role),
	}
}
