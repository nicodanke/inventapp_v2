package gapi

import (
	"context"
	"errors"
	"strings"

	db "github.com/nicodanke/inventapp_v2/db/sqlc"
	"github.com/nicodanke/inventapp_v2/pb/requests/v1/login"
	"github.com/nicodanke/inventapp_v2/utils"
	userValidator "github.com/nicodanke/inventapp_v2/validators/user"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	fields := strings.Split(req.GetUsername(), "@")
	account, err := server.store.GetAccountByCode(ctx, fields[1])
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Account does not exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get account: %s", err)
	}

	if !account.Active {
		return nil, status.Errorf(codes.Unauthenticated, "Account not active")
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "User does not exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get user: %s", err)
	}

	err = utils.CheckPassword(req.GetPassword(), user.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid password: %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating accessToken: %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating refreshToken: %s", err)
	}

	metadata := server.extractMetadata(ctx)

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    metadata.UserAgent,
		ClientIp:     metadata.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating session: %s", err)
	}

	rsp := &login.LoginResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}

func validateLoginRequest(req *login.LoginRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := userValidator.ValidateFullUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := userValidator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
