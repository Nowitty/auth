package user

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}