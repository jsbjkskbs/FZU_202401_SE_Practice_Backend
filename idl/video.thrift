namespace go api.video

include "base.thrift"

struct VideoFeedReq {
    1: optional string category;
    2: optional string access_token (api.header="Access-Token");
}

struct VideoFeedRespData {
    1: list<base.Video> items;
}
struct VideoFeedResp {
    1: i64 code;
    2: string msg;
    3: VideoFeedRespData data;
}

struct VideoInfoReq {
    1: required string video_id;
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
    1: string upload_video_url;
    2: string upload_cover_url;
}
struct VideoPublishResp {
    1: i64 code;
    2: string msg;
    3: VideoPublishRespData data;
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
}

struct VideoListRespData {
    1: list<base.Video> items;
}
struct VideoListResp {
    1: i64 code;
    2: string msg;
    3: VideoListRespData data;
}

struct VideoPopularReq {
    1: required i64 page_num;
    2: required i64 page_size;
}

struct VideoPopularRespData {
    1: list<base.Video> items;
}
struct VideoPopularResp {
    1: i64 code;
    2: string msg;
    3: VideoPopularRespData data;
}

struct VideoSubmitAllReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct VideoSubmitAllRespData {
    1: list<base.Video> items;
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
}

struct VideoSearchRespData {
    1: list<base.Video> items;
}
struct VideoSearchResp {
    1: i64 code;
    2: string msg;
    3: VideoSearchRespData data;
}

service VideoService {
    VideoFeedResp VideoFeedMethod(1: VideoFeedReq req) (api.get="/api/v1/video/feed");
    VideoInfoResp VideoInfoMethod(1: VideoInfoReq req) (api.get="/api/v1/video/info");
    VideoPublishResp VideoPublishMethod(1: VideoPublishReq req) (api.post="/api/v1/video/publish");
    VideoCategoriesResp VideoCategoriesMethod(1: VideoCategoriesReq req) (api.get="/api/v1/video/categories");
    VideoListResp VideoListMethod(1: VideoListReq req) (api.get="/api/v1/video/list");
    VideoPopularResp VideoPopularMethod(1: VideoPopularReq req) (api.get="/api/v1/video/popular");
    VideoSubmitAllResp VideoSubmitAllMethod(1: VideoSubmitAllReq req) (api.get="/api/v1/video/submit/all");
    VideoSubmitReviewResp VideoSubmitReviewMethod(1: VideoSubmitReviewReq req) (api.get="/api/v1/video/submit/review");
    VideoSubmitLockedResp VideoSubmitLockedMethod(1: VideoSubmitLockedReq req) (api.get="/api/v1/video/submit/locked");
    VideoSubmitPassedResp VideoSubmitPassedMethod(1: VideoSubmitPassedReq req) (api.get="/api/v1/video/submit/passed");
    VideoSearchResp VideoSearchMethod(1: VideoSearchReq req) (api.get="/api/v1/video/search");
}