//nolint
package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	// Gov errors reserve 200 ~ 299.
	CodeInvalidAddress    CodeType = 201
	CodeUnknownAddress    CodeType = 104
	CodeInsufficientCoins CodeType = 105
	CodeInvalidCoins      CodeType = 106
)

// NOTE: Don't stringer this, we'll put better messages in later.
func codeToDefaultMsg(code CodeType) string {
	switch code {
	case CodeInvalidAddress:
		return "Invalid address"
	case CodeUnknownAddress:
		return "Unknown address"
	case CodeInsufficientCoins:
		return "Insufficient coins"
	case CodeInvalidCoins:
		return "Invalid coins"
	case CodeUnknownRequest:
		return "Unknown request"
	default:
		return sdk.CodeToDefaultMsg(code)
	}
}

//----------------------------------------
// Error constructors

func ErrInvalidAddress(msg string) sdk.Error {
	return newError(CodeInvalidAddress, msg)
}

func ErrUnknownAddress(msg string) sdk.Error {
	return newError(CodeUnknownAddress, msg)
}

func ErrInsufficientCoins(msg string) sdk.Error {
	return newError(CodeInsufficientCoins, msg)
}

func ErrInvalidCoins(msg string) sdk.Error {
	return newError(CodeInvalidCoins, msg)
}

func ErrUnknownRequest(msg string) sdk.Error {
	return newError(CodeUnknownRequest, msg)
}

//----------------------------------------

func msgOrDefaultMsg(msg string, code CodeType) string {
	if msg != "" {
		return msg
	} else {
		return codeToDefaultMsg(code)
	}
}

func newError(code CodeType, msg string) sdk.Error {
	msg = msgOrDefaultMsg(msg, code)
	return sdk.NewError(code, msg)
}
