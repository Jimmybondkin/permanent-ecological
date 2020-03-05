package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeNameDoesNotExist sdk.CodeType = 101
	CodeGenesisReadError sdk.CodeType = 102
	CodeNotTheAdminError sdk.CodeType = 103
	CodeCoinAddError     sdk.CodeType = 104
)

// ErrNameDoesNotExist is the error for name not existing
func ErrNameDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNameDoesNotExist, "Name does not exist")
}

// ErrNameDoesNotExist is the error for name not existing
func ErrGenesisReadError(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeGenesisReadError, "Genesis read failed")
}

// ErrNameDoesNotExist is the error for name not existing
func ErrNotTheAdminError(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNotTheAdminError, "You are not admin")
}

// ErrNameDoesNotExist is the error for name not existing
func ErrCoinAddError(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeCoinAddError, "Coin add error")
}
