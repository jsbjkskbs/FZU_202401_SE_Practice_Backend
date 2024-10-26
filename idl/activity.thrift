namespace go api.activity

include "base.thrift"

struct ActivityInfoReq {
    1: required string activity_id;
}

struct ActivityInfoResp {
    1: i64 code;
    2: string msg;
    3: base.Activity data;
}

struct ActivityFeedReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct ActivityFeedRespData {
    1: list<base.Activity> items;
}
struct ActivityFeedResp {
    1: i64 code;
    2: string msg;
    3: ActivityFeedRespData data;
}

struct ActivityPublishReq {
    1: required string access_token (api.header="Access-Token");
    2: required string title;
    3: required string text;
    4: optional list<string> image;
    5: optional string ref_video;
    6: optional string ref_activity;
}

struct ActivityPublishResp {
    1: i64 code;
    2: string msg;
}

struct ActivityListReq {
    1: required string user_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct ActivityListRespData {
    1: list<base.Activity> items;
}
struct ActivityListResp {
    1: i64 code;
    2: string msg;
    3: ActivityListRespData data;
}

service ActivityService {
    ActivityInfoResp ActivityInfoMethod(1: ActivityInfoReq req) (api.get="/api/v1/activity/info");
    ActivityFeedResp ActivityFeedMethod(1: ActivityFeedReq req) (api.get="/api/v1/activity/feed");
    ActivityPublishResp ActivityPublishMethod(1: ActivityPublishReq req) (api.post="/api/v1/activity/publish");
    ActivityListResp ActivityListMethod(1: ActivityListReq req) (api.get="/api/v1/activity/list");
}