package gapi

import (
	"context"

	db "github.com/nicodanke/inventapp_v2/db/sqlc"
	"github.com/nicodanke/inventapp_v2/pb/requests/v1/role"
	userValidator "github.com/nicodanke/inventapp_v2/validators/user"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateRoleRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateRoleParams{
		Name:      req.GetName(),
		AccountID: authPayload.AccountID,
	}

	result, err := server.store.CreateRole(ctx, arg)
	if err != nil {
		return nil, internalError("failed to create role", err)
	}

	rsp := &role.CreateRoleResponse{
		Role: convertRole(result),
	}
	return rsp, nil
}

func validateCreateRoleRequest(req *role.CreateRoleRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := userValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}
