package gapi

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/inventapp_v2/db/sqlc"
	"github.com/nicodanke/inventapp_v2/pb"
	"github.com/nicodanke/inventapp_v2/utils"
	"github.com/nicodanke/inventapp_v2/validators"
	accountValidator "github.com/nicodanke/inventapp_v2/validators/account"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	violations := validateUpdateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateAccountParams{
		ID: req.GetId(),
		CompanyName: pgtype.Text{
			String: req.GetCompanyName(),
			Valid:  req.CompanyName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		WebUrl: pgtype.Text{
			String: req.GetWebUrl(),
			Valid:  req.WebUrl != nil,
		},
		Active: pgtype.Bool{
			Bool:  req.GetActive(),
			Valid: req.Active != nil,
		},
		Country: pgtype.Text{
			String: req.GetCompanyName(),
			Valid:  req.Country != nil,
		},
	}

	result, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to update account: %s", err)
	}

	rsp := &pb.UpdateAccountResponse{
		Account: convertAccount(result),
	}
	return rsp, nil
}

func (server *Server) ActivateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %s", err)
	}

	code := strings.ReplaceAll(req.GetCompanyName(), " ", "")

	arg := db.CreateAccountTxParams{
		Code:           code,
		CompanyName:    req.GetCompanyName(),
		Email:          req.GetEmail(),
		Name:           req.GetName(),
		Lastname:       req.GetLastname(),
		Username:       req.GetUsername(),
		Country:        req.GetCountry(),
		HashedPassword: hashedPassword,
	}

	result, err := server.store.CreateAccountTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			return nil, status.Error(codes.Internal, "Failed to create account: code already in use")
		}
		return nil, status.Errorf(codes.Internal, "Fail to create account: %s", err)
	}

	rsp := &pb.CreateAccountResponse{
		Account: convertAccount(result.Account),
		User:    convertUser(result.User),
	}
	return rsp, nil
}

func validateUpdateAccountRequest(req *pb.UpdateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.CompanyName != nil {
		if err := accountValidator.ValidateCompanyName(req.GetCompanyName()); err != nil {
			violations = append(violations, fieldViolation("companyName", err))
		}
	}

	if req.Country != nil {
		if err := accountValidator.ValidateCountry(req.GetCountry()); err != nil {
			violations = append(violations, fieldViolation("country", err))
		}
	}

	if req.Email != nil {
		if err := validators.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.WebUrl != nil {
		if err := accountValidator.ValidateWebUrl(req.GetWebUrl()); err != nil {
			violations = append(violations, fieldViolation("webUrl", err))
		}
	}

	if req.Phone != nil {
		if err := accountValidator.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}

	return violations
}
