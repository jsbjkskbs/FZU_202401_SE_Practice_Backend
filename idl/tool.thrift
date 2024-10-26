namespace go api.tool

include "base.thrift"

struct ToolDeleteReq {
    1: required string access_token (api.header="Access-Token");
    2: required string otype;
    3: required string oid;
}

struct ToolDeleteResp {
    1: i64 code;
    2: string msg;
}

struct AdminToolDeleteReq {
    1: required string access_token (api.header="Access-Token");
    2: required string otype;
    3: required string oid;
}

struct AdminToolDeleteResp {
    1: i64 code;
    2: string msg;
}

struct ToolUploadImageReq {
    1: required string access_token (api.header="Access-Token");
}

struct ToolUploadImageRespData {
    1: string url;
    2: string image_id;
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

struct UserRefreshReq {
    1: required string refresh_token (api.header="Refresh-Token");
}

struct UserRefreshRespData {
    1: string id
    2: string access_token;
}
struct UserRefreshResp {
    1: i64 code;
    2: string msg;
    3: UserRefreshRespData data;
}

service ToolService {
    ToolDeleteResp ToolDelete(1: ToolDeleteReq req) (api.delete="/api/v1/tool/delete");
    AdminToolDeleteResp AdminToolDelete(1: AdminToolDeleteReq req) (api.delete="/api/v1/admin/tool/delete");
    ToolUploadImageResp ToolUploadImage(1: ToolUploadImageReq req) (api.get="/api/v1/tool/upload/image");
    ToolGetImageResp ToolGetImage(1: ToolGetImageReq req) (api.get="/api/v1/tool/get/image");
    UserRefreshResp UserRefresh(1: UserRefreshReq req) (api.get="/api/v1/user/refresh");
}