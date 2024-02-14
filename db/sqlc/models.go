// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID          int64       `json:"id"`
	Code        string      `json:"code"`
	CompanyName string      `json:"company_name"`
	Phone       pgtype.Text `json:"phone"`
	Email       string      `json:"email"`
	WebUrl      pgtype.Text `json:"web_url"`
	Active      bool        `json:"active"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type AccountAddress struct {
	AccountID  int64         `json:"account_id"`
	Country    string        `json:"country"`
	State      pgtype.Text   `json:"state"`
	SubState   pgtype.Text   `json:"sub_state"`
	Street     pgtype.Text   `json:"street"`
	Number     pgtype.Text   `json:"number"`
	Unit       pgtype.Text   `json:"unit"`
	PostalCode pgtype.Text   `json:"postal_code"`
	Lat        pgtype.Float8 `json:"lat"`
	Lng        pgtype.Float8 `json:"lng"`
}

type AccountModule struct {
	ID        int64              `json:"id"`
	ModuleID  int64              `json:"module_id"`
	AccountID int64              `json:"account_id"`
	StartedAt time.Time          `json:"started_at"`
	EndedAt   pgtype.Timestamptz `json:"ended_at"`
	Price     float64            `json:"price"`
}

type AccountPlan struct {
	ID        int64              `json:"id"`
	PlanID    int64              `json:"plan_id"`
	AccountID int64              `json:"account_id"`
	StartedAt time.Time          `json:"started_at"`
	EndedAt   pgtype.Timestamptz `json:"ended_at"`
	Price     float64            `json:"price"`
}

type AccountPromotion struct {
	ID          int64              `json:"id"`
	PromotionID int64              `json:"promotion_id"`
	AccountID   int64              `json:"account_id"`
	StartedAt   time.Time          `json:"started_at"`
	EndedAt     pgtype.Timestamptz `json:"ended_at"`
}

type Feature struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Module struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type ModuleCountry struct {
	ModuleID int64   `json:"module_id"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
}

type ModuleFeature struct {
	ModuleID  int64 `json:"module_id"`
	FeatureID int64 `json:"feature_id"`
}

// Common table shared by all accounts
type Permission struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	ParentID pgtype.Int8 `json:"parent_id"`
}

type Plan struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PlanCountry struct {
	PlanID  int64   `json:"plan_id"`
	Country string  `json:"country"`
	Price   float64 `json:"price"`
}

type PlanModule struct {
	PlanID   int64 `json:"plan_id"`
	ModuleID int64 `json:"module_id"`
}

type Promotion struct {
	ID              int64         `json:"id"`
	Name            string        `json:"name"`
	DiscountPercent pgtype.Float8 `json:"discount_percent"`
	Amount          pgtype.Float8 `json:"amount"`
}

type Role struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

type RolePermission struct {
	RoleID       int64 `json:"role_id"`
	PermissionID int64 `json:"permission_id"`
}

type User struct {
	ID                int64       `json:"id"`
	Username          string      `json:"username"`
	Password          string      `json:"password"`
	Name              string      `json:"name"`
	Lastname          string      `json:"lastname"`
	Email             string      `json:"email"`
	Phone             pgtype.Text `json:"phone"`
	Active            bool        `json:"active"`
	IsAdmin           bool        `json:"is_admin"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	RoleID            int64       `json:"role_id"`
	AccountID         int64       `json:"account_id"`
}