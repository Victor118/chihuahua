package types

// DONTCOVER

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrDenomExists                           = errorsmod.Register(ModuleName, 2, "attempting to create a denom that already exists (has bank metadata)")
	ErrUnauthorized                          = errorsmod.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidDenom                          = errorsmod.Register(ModuleName, 4, "invalid denom")
	ErrInvalidCreator                        = errorsmod.Register(ModuleName, 5, "invalid creator")
	ErrInvalidAuthorityMetadata              = errorsmod.Register(ModuleName, 6, "invalid authority metadata")
	ErrInvalidGenesis                        = errorsmod.Register(ModuleName, 7, "invalid genesis")
	ErrSubdenomTooLong                       = errorsmod.Register(ModuleName, 8, fmt.Sprintf("subdenom too long, max length is %d bytes", MaxSubdenomLength))
	ErrCreatorTooLong                        = errorsmod.Register(ModuleName, 9, fmt.Sprintf("creator too long, max length is %d bytes", MaxCreatorLength))
	ErrDenomDoesNotExist                     = errorsmod.Register(ModuleName, 10, "denom does not exist")
	ErrBurnOrForceTransferFromModuleAccount  = errorsmod.Register(ModuleName, 11, "burning or force transfer from Module Account is not allowed")
	ErrBurnOrForceTransferFromBuilderAccount = errorsmod.Register(ModuleName, 12, "burning or force transfer from Builder Account is not allowed")
	ErrCapabilityNotEnabled                  = errorsmod.Register(ModuleName, 13, "this capability is not enabled on chain")
	ErrBadBlockParameters                    = errorsmod.Register(ModuleName, 14, "Bad block range, end block should be greater than start block")
	ErrBadBlockHeight                        = errorsmod.Register(ModuleName, 15, "Bad block range, chain height is greater than start block")
)
