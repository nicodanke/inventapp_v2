package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/inventapp_v2/db/sqlc"
	"github.com/nicodanke/inventapp_v2/pb/requests/v1/role"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*emptypb.Empty, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.DeleteRoleParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteRole(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintf("failed to delete role with id: %d", req.GetId()), err)
	}

	return &emptypb.Empty{}, nil
}