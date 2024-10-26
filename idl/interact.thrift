namespace go api.interact

include "base.thrift"

struct InteractLikeActionReq {
    1: required string access_token (api.header="Access-Token");
    2: required string otype;
    3: required string oid;
    4: required i64 action_type;
}

struct InteractLikeActionResp {
    1: i64 code;
    2: string msg;
}

struct InteractLikeListReq {
    1: required string user_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct InteractLikeListRespData {
    1: list<base.Video> items;
}
struct InteractLikeListResp {
    1: i64 code;
    2: string msg;
    3: InteractLikeListRespData data;
}

struct InteractCommentPublishReq {
    1: required string access_token (api.header="Access-Token");
    2: required string otype;
    3: required string oid;
    4: required string content;
}

struct InteractCommentPublishResp {
    1: i64 code;
    2: string msg;
}

struct InteractCommentListReq {
    1: required string otype;
    2: required string oid;
    3: required i64 page_num;
    4: required i64 page_size;
}

struct InteractCommentListRespData {
    1: list<base.Comment> items;
}
struct InteractCommentListResp {
    1: i64 code;
    2: string msg;
    3: InteractCommentListRespData data;
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
    InteractLikeActionResp InteractLikeActionMethod(1: InteractLikeActionReq req) (api.post="/api/v1/interact/like");
    InteractLikeListResp InteractLikeListMethod(1: InteractLikeListReq req) (api.get="/api/v1/interact/like/list");
    InteractCommentPublishResp InteractCommentPublishMethod(1: InteractCommentPublishReq req) (api.post="/api/v1/interact/comment/publish");
    InteractCommentListResp InteractCommentListMethod(1: InteractCommentListReq req) (api.get="/api/v1/interact/comment/list");
    InteractMessageSendResp InteractMessageSendMethod(1: InteractMessageSendReq req) (api.post="/api/v1/interact/message/send");
}