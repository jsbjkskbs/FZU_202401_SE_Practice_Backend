namespace go api.interact

include "base.thrift"

struct InteractLikeVideoActionReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
    3: required i64 action_type;
}

struct InteractLikeVideoActionResp {
    1: i64 code;
    2: string msg;
}

struct InteractLikeActivityActionReq {
    1: required string access_token (api.header="Access-Token");
    2: required string activity_id;
    3: required i64 action_type;
}

struct InteractLikeActivityActionResp {
    1: i64 code;
    2: string msg;
}

struct InteractLikeCommentActionReq {
    1: required string access_token (api.header="Access-Token");
    2: required string comment_type;
    3: required string comment_id;
    4: required i64 action_type;
}

struct InteractLikeCommentActionResp {
    1: i64 code;
    2: string msg;
}

struct InteractLikeVideoListReq {
    1: required string user_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct InteractLikeVideoListRespData {
    1: list<base.Video> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct InteractLikeVideoListResp {
    1: i64 code;
    2: string msg;
    3: InteractLikeVideoListRespData data;
}

struct InteractCommentVideoPublishReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
    3: required string content;
    4: optional string root_id;
    5: optional string parent_id;
}

struct InteractCommentVideoPublishResp {
    1: i64 code;
    2: string msg;
}

struct InteractCommentActivityPublishReq {
    1: required string access_token (api.header="Access-Token");
    2: required string activity_id;
    3: required string content;
    4: optional string root_id;
    5: optional string parent_id;
}

struct InteractCommentActivityPublishResp {
    1: i64 code;
    2: string msg;
}

struct InteractCommentVideoListReq {
    1: required string video_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct InteractCommentVideoListRespData {
    1: list<base.Comment> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct InteractCommentVideoListResp {
    1: i64 code;
    2: string msg;
    3: InteractCommentVideoListRespData data;
}

struct InteractCommentActivityListReq {
    1: required string activity_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct InteractCommentActivityListRespData {
    1: list<base.Comment> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}

struct InteractCommentActivityListResp {
    1: i64 code;
    2: string msg;
    3: InteractCommentActivityListRespData data;
}

struct InteractVideoChildCommentListReq {
    1: required string comment_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct InteractVideoChildCommentListRespData {
    1: list<base.Comment> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct InteractVideoChildCommentListResp {
    1: i64 code;
    2: string msg;
    3: InteractVideoChildCommentListRespData data;
}

struct InteractActivityChildCommentListReq {
    1: required string comment_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct InteractActivityChildCommentListRespData {
    1: list<base.Comment> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct InteractActivityChildCommentListResp {
    1: i64 code;
    2: string msg;
    3: InteractActivityChildCommentListRespData data;
}

struct InteractMessageSendReq {
    1: required string access_token (api.header="Access-Token");
    2: required string to_user_id;
    3: required string content;
}

struct InteractMessageSendResp {
    1: i64 code;
    2: string msg;
}

service InteractService {
    InteractLikeVideoActionResp InteractLikeVideoActionMethod(1: InteractLikeVideoActionReq req) (api.post="/api/v1/interact/like/video/action");
    InteractLikeActivityActionResp InteractLikeActivityActionMethod(1: InteractLikeActivityActionReq req) (api.post="/api/v1/interact/like/activity/action");
    InteractLikeCommentActionResp InteractLikeCommentActionMethod(1: InteractLikeCommentActionReq req) (api.post="/api/v1/interact/like/comment/action");
    InteractLikeVideoListResp InteractLikeVideoListMethod(1: InteractLikeVideoListReq req) (api.get="/api/v1/interact/like/video/list");
    InteractCommentVideoPublishResp InteractCommentVideoPublishMethod(1: InteractCommentVideoPublishReq req) (api.post="/api/v1/interact/comment/video/publish");
    InteractCommentActivityPublishResp InteractCommentActivityPublishMethod(1: InteractCommentActivityPublishReq req) (api.post="/api/v1/interact/comment/activity/publish");
    InteractCommentVideoListResp InteractCommentVideoListMethod(1: InteractCommentVideoListReq req) (api.get="/api/v1/interact/comment/video/list");
    InteractCommentActivityListResp InteractCommentActivityListMethod(1: InteractCommentActivityListReq req) (api.get="/api/v1/interact/comment/activity/list");
    InteractVideoChildCommentListResp InteractVideoChildCommentListMethod(1: InteractVideoChildCommentListReq req) (api.get="/api/v1/interact/video/child_comment/list");
    InteractActivityChildCommentListResp InteractActivityChildCommentListMethod(1: InteractActivityChildCommentListReq req) (api.get="/api/v1/interact/activity/child_comment/list");
    InteractMessageSendResp InteractMessageSendMethod(1: InteractMessageSendReq req) (api.post="/api/v1/interact/message/send");
}