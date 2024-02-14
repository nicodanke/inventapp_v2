package gapi

import (
	"context"
	"strings"

	db "github.com/nicodanke/bankTutorial/db/sqlc"
	"github.com/nicodanke/bankTutorial/pb"
	"github.com/nicodanke/bankTutorial/utils"
	"github.com/nicodanke/bankTutorial/validators"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %s", err)
	}

	code := strings.ReplaceAll(req.GetCompanyName(), " ", "")

	// TODO: check that account does not exists (if code us available)

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
		return nil, status.Errorf(codes.Internal, "Fail to create account: %s", err)
	}

	rsp := &pb.CreateAccountResponse{
		Account: convertAccount(result.Account),
		User:    convertUser(result.User),
	}
	return rsp, nil
}

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if err := validators.ValidateLastname(req.GetLastname()); err != nil {
		violations = append(violations, fieldViolation("lastname", err))
	}

	if err := validators.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := validators.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := validators.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := validators.ValidateCompanyName(req.GetCompanyName()); err != nil {
		violations = append(violations, fieldViolation("companyName", err))
	}

	if err := validators.ValidateCountry(req.GetCountry()); err != nil {
		violations = append(violations, fieldViolation("country", err))
	}

	return violations
}
