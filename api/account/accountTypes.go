package account

import (
	db "github.com/guilleamutio/go4money/db/sqlc"
)

type Server struct {
	store *db.Store
}

type updateAccountRequest struct {
	ID      int64 `json:"ID" binding:"required"`
	Balance int64 `json:"balance" binding:"required,min=0"`
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

type deleteAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type deleteAccountResponse struct {
	Deleted bool
	Account db.Account
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
