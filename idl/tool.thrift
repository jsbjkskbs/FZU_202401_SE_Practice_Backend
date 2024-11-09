namespace go api.tool

include "base.thrift"

struct ToolDeleteVideoReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
}

struct ToolDeleteVideoResp {
    1: i64 code;
    2: string msg;
}

struct ToolDeleteActivityReq {
    1: required string access_token (api.header="Access-Token");
    2: required string activity_id;
}

struct ToolDeleteActivityResp {
    1: i64 code;
    2: string msg;
}

struct ToolDeleteCommentReq {
    1: required string access_token (api.header="Access-Token");
    2: required string comment_type;
    3: required string from_media_id;
    4: required string comment_id;
}

struct ToolDeleteCommentResp {
    1: i64 code;
    2: string msg;
}

struct AdminToolDeleteVideoReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
}

struct AdminToolDeleteVideoResp {
    1: i64 code;
    2: string msg;
}

struct AdminToolDeleteActivityReq {
    1: required string access_token (api.header="Access-Token");
    2: required string activity_id;
}

struct AdminToolDeleteActivityResp {
    1: i64 code;
    2: string msg;
}

struct AdminToolDeleteCommentReq {
    1: required string access_token (api.header="Access-Token");
    2: required string comment_type;
    3: required string from_media_id;
    4: required string comment_id;
}

struct AdminToolDeleteCommentResp {
    1: i64 code;
    2: string msg;
}

struct AdminToolDeleteResp {
    1: i64 code;
    2: string msg;
}

struct ToolUploadImageReq {
    1: required string access_token (api.header="Access-Token");
}

struct ToolUploadImageRespData {
    1: string image_id;
    2: string upload_url;
    3: string upload_key;
    4: string uptoken;
}
struct ToolUploadImageResp {
    1: i64 code;
    2: string msg;
    3: ToolUploadImageRespData data;
}

struct ToolGetImageReq {
    1: required string image_id;
}

struct ToolGetImageRespData {
    1: string url;
}
struct ToolGetImageResp {
    1: i64 code;
    2: string msg;
    3: ToolGetImageRespData data;
}

struct ToolTokenRefreshReq {
    1: required string refresh_token (api.header="Refresh-Token");
}

struct ToolTokenRefreshRespData {
    1: string id
    2: string access_token;
}
struct ToolTokenRefreshResp {
    1: i64 code;
    2: string msg;
    3: ToolTokenRefreshRespData data;
}

struct ToolRefreshTokenRefreshReq {
    1: required string refresh_token (api.header="Refresh-Token");
}

struct ToolRefreshTokenRefreshRespData {
    1: string id
    2: string refresh_token;
}
struct ToolRefreshTokenRefreshResp {
    1: i64 code;
    2: string msg;
    3: ToolRefreshTokenRefreshRespData data;
}

service ToolService {
    ToolDeleteVideoResp ToolDeleteVideo(1: ToolDeleteVideoReq req) (api.delete="/api/v1/tool/delete/video");
    ToolDeleteActivityResp ToolDeleteActivity(1: ToolDeleteActivityReq req) (api.delete="/api/v1/tool/delete/activity");
    ToolDeleteCommentResp ToolDeleteComment(1: ToolDeleteCommentReq req) (api.delete="/api/v1/tool/delete/comment");
    AdminToolDeleteVideoResp AdminToolDeleteVideo(1: AdminToolDeleteVideoReq req) (api.delete="/api/v1/admin/tool/delete/video");
    AdminToolDeleteActivityResp AdminToolDeleteActivity(1: AdminToolDeleteActivityReq req) (api.delete="/api/v1/admin/tool/delete/activity");
    AdminToolDeleteCommentResp AdminToolDeleteComment(1: AdminToolDeleteCommentReq req) (api.delete="/api/v1/admin/tool/delete/comment");
    ToolUploadImageResp ToolUploadImage(1: ToolUploadImageReq req) (api.get="/api/v1/tool/upload/image");
    ToolGetImageResp ToolGetImage(1: ToolGetImageReq req) (api.get="/api/v1/tool/get/image");
    ToolTokenRefreshResp ToolTokenRefresh(1: ToolTokenRefreshReq req) (api.get="/api/v1/tool/token/refresh");
    ToolRefreshTokenRefreshResp ToolRefreshTokenRefresh(1: ToolRefreshTokenRefreshReq req) (api.get="/api/v1/tool/refresh_token/refresh");
}