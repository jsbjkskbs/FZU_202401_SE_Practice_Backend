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
	MfaAuthFailedErrorCode     = 7001
	MfaGenerateFailedErrorCode = 7002

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
	InternalServerErrorMsg = "内部服务器错误"
	ThirdPartyCallErrorMsg = "第三方调用错误"
	DatabaseCallErrorMsg   = "数据访问错误"
	FileSystemErrorMsg     = "文件系统错误"

	// Auth error
	AccessTokenInvalidErrorMsg    = "Access token 无效"
	AccessTokenExpiredErrorMsg    = "Access token 过期"
	AccessTokenForbiddenErrorMsg  = "Access token 被禁止"
	RefreshTokenInvalidErrorMsg   = "Refresh token 无效"
	RefreshTokenExpiredErrorMsg   = "Refresh token 过期"
	RefreshTokenForbiddenErrorMsg = "Refresh token 被禁止"
	AccountOrPasswordInvalidMsg   = "账号或密码错误"

	// Api error
	ApiForbiddenErrorMsg    = "API 被废弃"
	ApiInvokeFailedErrorMsg = "API 调用失败"

	// Param error
	ParamInvalidErrorMsg = "参数错误"

	// Resource error
	ResourceNotFoundErrorMsg = "资源不存在"
	ResourceConflictErrorMsg = "资源冲突"

	// Power error
	PowerNotEnoughErrorMsg = "权限不足"

	// MFA error
	MfaAuthFailedErrorMsg     = "MFA 认证失败"
	MfaGenerateFailedErrorMsg = "MFA 生成失败"

	// Query Limit error
	QueryLimitErrorMsg = "服务器繁忙，请稍后再试"

	// User Offline error
	UserOfflineErrorMsg = "用户已离线"
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
	MfaAuthFailed     = NewErrno(MfaAuthFailedErrorCode, MfaAuthFailedErrorMsg)
	MfaGenerateFailed = NewErrno(MfaGenerateFailedErrorCode, MfaGenerateFailedErrorMsg)

	// Query Limit error
	QueryLimit = NewErrno(QueryLimitErrorCode, QueryLimitErrorMsg)

	// User Offline error
	UserOffline = NewErrno(UserOfflineErrorCode, UserOfflineErrorMsg)

	// Custom error(not cause by system)
	CustomError = NewErrno(CustomErrorCode, "")
)
