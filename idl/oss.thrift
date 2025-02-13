namespace go api.oss

include "base.thrift"

struct OssCallbackAvatarReq {
    1: required string authorization (api.header="Authorization");
    2: string key;
    3: string bucket;
    4: string name;
    5: i64 fsize;
    6: string hash;
    7: string otype;
    8: string oid;
}

struct OssCallbackAvatarResp {
}

struct OssCallbackVideoReq {
    1: required string authorization (api.header="Authorization");
    2: string key;
    3: string bucket;
    4: string name;
    5: i64 fsize;
    6: string hash;
    7: string otype;
    8: string oid;
}

struct OssCallbackVideoCoverReq {
    1: required string authorization (api.header="Authorization");
    2: string key;
    3: string bucket;
    4: string name;
    5: i64 fsize;
    6: string hash;
    7: string otype;
    8: string oid;
}

struct OssCallbackVideoResp {
}

struct OssCallbackImageReq {
    1: required string authorization (api.header="Authorization");
    2: string key;
    3: string bucket;
    4: string name;
    5: i64 fsize;
    6: string hash;
    7: string otype;
    8: string oid;
    9: string user_id;
}

struct OssCallbackImageResp {
}

struct OssCallbackFopReq {
    1: required string authorization (api.header="Authorization");
}

struct OssCallbackFopResp {
}

service OssService {
    OssCallbackAvatarResp OssCallbackAvatarMethod(1: OssCallbackAvatarReq req) (api.post="/api/v1/oss/callback/avatar");
    OssCallbackFopResp OssCallbackVideoMethod(1: OssCallbackVideoReq req) (api.post="/api/v1/oss/callback/video");
    OssCallbackFopResp OssCallbackVideoCoverMethod(1: OssCallbackVideoCoverReq req) (api.post="/api/v1/oss/callback/cover");
    OssCallbackImageResp OssCallbackImageMethod(1: OssCallbackImageReq req) (api.post="/api/v1/oss/callback/image");
    OssCallbackFopResp OssCallbackFopMethod(1: OssCallbackFopReq req) (api.post="/api/v1/oss/callback/fop");
}