namespace go api.user

include "base.thrift"

struct UserRegisterReq {
    1: required string username;
    2: required string password;
    3: required string email(api.vd="email($); msg:'邮箱格式不正确'");
    4: required string code;
}

struct UserRegisterResp {
    1: i64 code;
    2: string msg;
}

struct UserSecurityEmailCodeReq {
    1: required string email(api.vd="email($); msg:'邮箱格式不正确'");
}

struct UserSecurityEmailCodeResp {
    1: i64 code;
    2: string msg;
}

struct UserLoginReq {
    1: required string username;
    2: required string password;
    3: optional string mfa_code;
}

struct UserLoginResp {
    1: i64 code;
    2: string msg;
    3: base.UserWithToken data;
}

struct UserInfoReq {
    1: required string user_id;
}

struct UserInfoResp {
    1: i64 code;
    2: string msg;
    3: base.User data;
}

struct UserFollowerCountReq {
    1: required string user_id;
}

struct UserFollowerCountRespData {
    1: string id;
    2: i64 follower_count;
}
struct UserFollowerCountResp {
    1: i64 code;
    2: string msg;
    3: UserFollowerCountRespData data;
}

struct UserFollowingCountReq {
    1: required string user_id;
}

struct UserFollowingCountRespData {
    1: string id;
    2: i64 following_count;
}
struct UserFollowingCountResp {
    1: i64 code;
    2: string msg;
    3: UserFollowingCountRespData data;
}

struct UserLikeCountReq {
    1: required string user_id;
}

struct UserLikeCountRespData {
    1: string id;
    2: i64 like_count;
}
struct UserLikeCountResp {
    1: i64 code;
    2: string msg;
    3: UserLikeCountRespData data;
}

struct UserAvatarUploadReq {
    1: required string access_token (api.header="Access-Token");
}

struct UserAvatarUploadData {
    1: string upload_url;
    2: string uptoken;
}
struct UserAvatarUploadResp {
    1: i64 code;
    2: string msg;
    3: UserAvatarUploadData data;
}

struct UserMfaQrcodeReq {
    1: required string access_token (api.header="Access-Token");
}

struct UserMfaQrcodeData {
    1: string secret;
    2: string qrcode;
}
struct UserMfaQrcodeResp {
    1: i64 code;
    2: string msg;
    3: UserMfaQrcodeData data;
}

struct UserMfaBindReq {
    1: required string access_token (api.header="Access-Token");
    2: required string code;
    3: required string secret;
}

struct UserMfaBindResp {
    1: i64 code;
    2: string msg;
}

struct UserSearchReq {
    1: required string keyword;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct UserSearchRespData {
    1: list<base.User> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
}
struct UserSearchResp {
    1: i64 code;
    2: string msg;
    3: UserSearchRespData data;
}

struct UserPasswordRetrieveReq {
    1: required string otype (api.vd="$ == 'email' || $ == 'phone'; msg:'类型不正确'");
    2: required string oid (api.vd="((Otype)$ == 'phone' && phone($)) || ((Otype)$ == 'email' && email($)); msg:'邮箱或电话格式不正确'");
}

struct UserPasswordRetrieveResp {
    1: i64 code;
    2: string msg;
}

struct UserPasswordResetReq {
    1: required string otype (api.vd="$ == 'email' || $ == 'phone'; msg:'类型不正确'");
    2: required string oid (api.vd="((Otype)$ == 'phone' && phone($)) || ((Otype)$ == 'email' && email($)); msg:'邮箱或电话格式不正确'");
    3: required string password;
    4: required string code;
}

struct UserPasswordResetResp {
    1: i64 code;
    2: string msg;
}

service UserService {
    UserRegisterResp RegisterMethod(1: UserRegisterReq req) (api.post="/api/v1/user/register");
    UserSecurityEmailCodeResp SecurityEmailCodeMethod(1: UserSecurityEmailCodeReq req) (api.post="/api/v1/user/security/email/code");
    UserLoginResp LoginMethod(1: UserLoginReq req) (api.post="/api/v1/user/login");
    UserInfoResp InfoMethod(1: UserInfoReq req) (api.get="/api/v1/user/info");
    UserFollowerCountResp FollowerCountMethod(1: UserFollowerCountReq req) (api.get="/api/v1/user/follower_count");
    UserFollowingCountResp FollowingCountMethod(1: UserFollowingCountReq req) (api.get="/api/v1/user/following_count");
    UserLikeCountResp LikeCountMethod(1: UserLikeCountReq req) (api.get="/api/v1/user/like_count");
    UserAvatarUploadResp AvatarUploadMethod(1: UserAvatarUploadReq req) (api.get="/api/v1/user/avatar/upload");
    UserMfaQrcodeResp MfaQrcodeMethod(1: UserMfaQrcodeReq req) (api.get="/api/v1/user/mfa/qrcode");
    UserMfaBindResp MfaBindMethod(1: UserMfaBindReq req) (api.post="/api/v1/user/mfa/bind");
    UserSearchResp SearchMethod(1: UserSearchReq req) (api.get="/api/v1/user/search");
    UserPasswordRetrieveResp PasswordRetrieveMethod(1: UserPasswordRetrieveReq req) (api.post="/api/v1/user/security/password/retrieve");
    UserPasswordResetResp PasswordResetMethod(1: UserPasswordResetReq req) (api.post="/api/v1/user/security/password/reset");
}