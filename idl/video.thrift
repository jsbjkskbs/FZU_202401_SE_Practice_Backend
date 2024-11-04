namespace go api.video

include "base.thrift"

struct VideoFeedReq {
    1: optional string category;
}

struct VideoFeedRespData {
    1: list<base.Video> items;
}
struct VideoFeedResp {
    1: i64 code;
    2: string msg;
    3: VideoFeedRespData data;
}

struct VideoCustomFeedReq {
    1: optional string category;
    2: required string access_token (api.header="Access-Token");
}

struct VideoCustomFeedRespData {
    1: list<base.Video> items;
}
struct VideoCustomFeedResp {
    1: i64 code;
    2: string msg;
    3: VideoCustomFeedRespData data;
}

struct VideoInfoReq {
    1: required string video_id;
    2: optional string access_token (api.header="Access-Token");
}

struct VideoInfoResp {
    1: i64 code;
    2: string msg;
    3: base.Video data;
}

struct VideoPublishReq {
    1: required string access_token (api.header="Access-Token");
    2: required string title;
    3: required string description;
    4: required string category;
    5: required list<string> labels;
}

struct VideoPublishRespData {
    1: string upload_url;
    2: string upload_key;
    3: string uptoken;
}
struct VideoPublishResp {
    1: i64 code;
    2: string msg;
    3: VideoPublishRespData data;
}

struct VideoCoverUploadReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
}

struct VideoCoverUploadRespData {
    1: string upload_url;
    2: string upload_key;
    3: string uptoken;
}
struct VideoCoverUploadResp {
    1: i64 code;
    2: string msg;
    3: VideoCoverUploadRespData data;
}

struct VideoCategoriesReq {
}

struct VideoCategoriesRespData {
    1: list<string> items;
}
struct VideoCategoriesResp {
    1: i64 code;
    2: string msg;
    3: VideoCategoriesRespData data;
}

struct VideoListReq {
    1: required string user_id;
    2: required i64 page_num;
    3: required i64 page_size;
    4: optional string access_token (api.header="Access-Token");
}

struct VideoListRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct VideoListResp {
    1: i64 code;
    2: string msg;
    3: VideoListRespData data;
}

struct VideoSubmitAllReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct VideoSubmitAllRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct VideoSubmitAllResp {
    1: i64 code;
    2: string msg;
    3: VideoSubmitAllRespData data;
}

struct VideoSubmitReviewReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct VideoSubmitReviewRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct VideoSubmitReviewResp {
    1: i64 code;
    2: string msg;
    3: VideoSubmitReviewRespData data;
}

struct VideoSubmitLockedReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct VideoSubmitLockedRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct VideoSubmitLockedResp {
    1: i64 code;
    2: string msg;
    3: VideoSubmitLockedRespData data;
}

struct VideoSubmitPassedReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct VideoSubmitPassedRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct VideoSubmitPassedResp {
    1: i64 code;
    2: string msg;
    3: VideoSubmitPassedRespData data;
}

struct VideoSearchReq {
    1: required string keyword;
    2: required i64 page_num;
    3: required i64 page_size;
    4: optional i64 from_date;
    5: optional i64 to_date;
    6: optional string access_token (api.header="Access-Token");
}

struct VideoSearchRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct VideoSearchResp {
    1: i64 code;
    2: string msg;
    3: VideoSearchRespData data;
}

service VideoService {
    VideoFeedResp VideoFeedMethod(1: VideoFeedReq req) (api.get="/api/v1/video/feed");
    VideoCustomFeedResp VideoCustomFeedMethod(1: VideoCustomFeedReq req) (api.get="/api/v1/video/custom/feed");
    VideoInfoResp VideoInfoMethod(1: VideoInfoReq req) (api.get="/api/v1/video/info");
    VideoPublishResp VideoPublishMethod(1: VideoPublishReq req) (api.post="/api/v1/video/publish");
    VideoCoverUploadResp VideoCoverUploadMethod(1: VideoCoverUploadReq req) (api.post="/api/v1/video/cover/upload");
    VideoCategoriesResp VideoCategoriesMethod(1: VideoCategoriesReq req) (api.get="/api/v1/video/categories");
    VideoListResp VideoListMethod(1: VideoListReq req) (api.get="/api/v1/video/list");
    VideoSubmitAllResp VideoSubmitAllMethod(1: VideoSubmitAllReq req) (api.get="/api/v1/video/submit/all");
    VideoSubmitReviewResp VideoSubmitReviewMethod(1: VideoSubmitReviewReq req) (api.get="/api/v1/video/submit/review");
    VideoSubmitLockedResp VideoSubmitLockedMethod(1: VideoSubmitLockedReq req) (api.get="/api/v1/video/submit/locked");
    VideoSubmitPassedResp VideoSubmitPassedMethod(1: VideoSubmitPassedReq req) (api.get="/api/v1/video/submit/passed");
    VideoSearchResp VideoSearchMethod(1: VideoSearchReq req) (api.get="/api/v1/video/search");
}