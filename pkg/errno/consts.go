package errno

const (
	NoErrorCode = 0

	// Common error
	InternalServerErrorCode = 1001
	ThirdPartyCallErrorCode = 1002
	DatabaseCallErrorCode   = 1004
	FileSystemErrorCode     = 1008

	// Auth error
	AccessTokenInvalidErrorCode       = 2001
	AccessTokenExpiredErrorCode       = 2002
	AccessTokenForbiddenErrorCode     = 2004
	RefreshTokenInvalidErrorCode      = 2008
	RefreshTokenExpiredErrorCode      = 2016
	RefreshTokenForbiddenErrorCode    = 2032
	AccountOrPasswordInvalidErrorCode = 2064

	// Api error
	ApiForbiddenErrorCode    = 3001
	ApiInvokeFailedErrorCode = 3002

	// Param error
	ParamInvalidErrorCode  = 4001
	ParamVaildateErrorCode = 4002

	// Resource error
	ResourceNotFoundErrorCode = 5001
	ResourceConflictErrorCode = 5002

	// Power error
	PowerNotEnoughErrorCode = 6001

	// MFA error
	MfaAuthFailedErrorCode = 7001

	// Query Limit error
	QueryLimitErrorCode = 8001

	// User Offline error
	UserOfflineErrorCode = 9001

	// Custom error(not cause by system)
	CustomErrorCode = 10000
)

const (
	NoErrorMsg = "ok"

	// Common error
	InternalServerErrorMsg = "Internal server error"
	ThirdPartyCallErrorMsg = "Third party call error"
	DatabaseCallErrorMsg   = "Database call error"
	FileSystemErrorMsg     = "File system error"

	// Auth error
	AccessTokenInvalidErrorMsg    = "Access token invalid"
	AccessTokenExpiredErrorMsg    = "Access token expired"
	AccessTokenForbiddenErrorMsg  = "Access token forbidden"
	RefreshTokenInvalidErrorMsg   = "Refresh token invalid"
	RefreshTokenExpiredErrorMsg   = "Refresh token expired"
	RefreshTokenForbiddenErrorMsg = "Refresh token forbidden"
	AccountOrPasswordInvalidMsg   = "Account or password invalid"

	// Api error
	ApiForbiddenErrorMsg    = "Api forbidden"
	ApiInvokeFailedErrorMsg = "Api invoke failed"

	// Param error
	ParamInvalidErrorMsg = "Param invalid"

	// Resource error
	ResourceNotFoundErrorMsg = "Resource not found"
	ResourceConflictErrorMsg = "Resource conflict"

	// Power error
	PowerNotEnoughErrorMsg = "Power not enough"

	// MFA error
	MfaAuthFailedErrorMsg = "MFA auth failed"

	// Query Limit error
	QueryLimitErrorMsg = "Server receive too many requests at this moment, please wait a moment and try again"

	// User Offline error
	UserOfflineErrorMsg = "User offline"
)

var (
	NoError = NewErrno(NoErrorCode, NoErrorMsg)

	// Common error
	InternalServerError = NewErrno(InternalServerErrorCode, InternalServerErrorMsg)
	ThirdPartyCallError = NewErrno(ThirdPartyCallErrorCode, ThirdPartyCallErrorMsg)
	DatabaseCallError   = NewErrno(DatabaseCallErrorCode, DatabaseCallErrorMsg)
	FileSystemError     = NewErrno(FileSystemErrorCode, FileSystemErrorMsg)

	// Auth error
	AccessTokenInvalid       = NewErrno(AccessTokenInvalidErrorCode, AccessTokenInvalidErrorMsg)
	AccessTokenExpired       = NewErrno(AccessTokenExpiredErrorCode, AccessTokenExpiredErrorMsg)
	AccessTokenForbidden     = NewErrno(AccessTokenForbiddenErrorCode, AccessTokenForbiddenErrorMsg)
	RefreshTokenInvalid      = NewErrno(RefreshTokenInvalidErrorCode, RefreshTokenInvalidErrorMsg)
	RefreshTokenExpired      = NewErrno(RefreshTokenExpiredErrorCode, RefreshTokenExpiredErrorMsg)
	RefreshTokenForbidden    = NewErrno(RefreshTokenForbiddenErrorCode, RefreshTokenForbiddenErrorMsg)
	AccountOrPasswordInvalid = NewErrno(AccountOrPasswordInvalidErrorCode, AccountOrPasswordInvalidMsg)

	// Api error
	ApiForbidden    = NewErrno(ApiForbiddenErrorCode, ApiForbiddenErrorMsg)
	ApiInvokeFailed = NewErrno(ApiInvokeFailedErrorCode, ApiInvokeFailedErrorMsg)

	// Param error
	ParamInvalid = NewErrno(ParamInvalidErrorCode, ParamInvalidErrorMsg)

	// Resource error
	ResourceNotFound = NewErrno(ResourceNotFoundErrorCode, ResourceNotFoundErrorMsg)
	ResourceConflict = NewErrno(ResourceConflictErrorCode, ResourceConflictErrorMsg)

	// Power error
	PowerNotEnough = NewErrno(PowerNotEnoughErrorCode, PowerNotEnoughErrorMsg)

	// MFA error
	MfaAuthFailed = NewErrno(MfaAuthFailedErrorCode, MfaAuthFailedErrorMsg)

	// Query Limit error
	QueryLimit = NewErrno(QueryLimitErrorCode, QueryLimitErrorMsg)

	// User Offline error
	UserOffline = NewErrno(UserOfflineErrorCode, UserOfflineErrorMsg)

	// Custom error(not cause by system)
	CustomError = NewErrno(CustomErrorCode, "")
)
