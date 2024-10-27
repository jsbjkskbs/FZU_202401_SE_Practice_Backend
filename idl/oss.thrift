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

struct OssCallbackFopReq {
    1: required string authorization (api.header="Authorization");
}

struct OssCallbackFopResp {
}

service OssService {
    OssCallbackAvatarResp OssCallbackAvatarMethod(1: OssCallbackAvatarReq req) (api.post="/api/v1/oss/callback/avatar");
    OssCallbackFopResp OssCallbackFopMethod(1: OssCallbackFopReq req) (api.post="/api/v1/oss/callback/fop");
}