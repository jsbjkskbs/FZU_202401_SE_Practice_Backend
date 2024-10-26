namespace go api.oss

include "base.thrift"

struct OssCallbackReq {
    1: required string authorization (api.header="Authorization");
    2: required string key;
    3: required string bucket;
    4: required string name;
    5: required i64 fsize;
    6: required string hash;
    7: required string otype;
    8: required string oid;
}

struct OssCallbackResp {
}

service OssService {
    OssCallbackResp OssCallbackMethod(1: OssCallbackReq req) (api.post="/api/v1/oss/callback");
}