namespace go api.relation

include "base.thrift"

struct RelationFollowActionReq {
    1: required string access_token (api.header="Access-Token");
    2: required string to_user_id;
    3: required i64 action_type;
}

struct RelationFollowActionResp {
    1: i64 code;
    2: string msg;
}

struct RelationFollowListReq {
    1: required string user_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct RelationFollowListRespData {
    1: list<base.User> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct RelationFollowListResp {
    1: i64 code;
    2: string msg;
    3: RelationFollowListRespData data;
}

struct RelationFollowerListReq {
    1: required string user_id;
    2: required i64 page_num;
    3: required i64 page_size;
}

struct RelationFollowerListRespData {
    1: list<base.User> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct RelationFollowerListResp {
    1: i64 code;
    2: string msg;
    3: RelationFollowerListRespData data;
}

struct RelationFriendListReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct RelationFriendListRespData {
    1: list<base.User> items;
    2: bool is_end;
    3: i64 page_num;
    4: i64 page_size;
    5: i64 total;
}
struct RelationFriendListResp {
    1: i64 code;
    2: string msg;
    3: RelationFriendListRespData data;
}

service RelationService {
    RelationFollowActionResp FollowActionMethod(1: RelationFollowActionReq req) (api.post="/api/v1/relation/follow/action");
    RelationFollowListResp FollowListMethod(1: RelationFollowListReq req) (api.get="/api/v1/relation/follow/list");
    RelationFollowerListResp FollowerListMethod(1: RelationFollowerListReq req) (api.get="/api/v1/relation/follower/list");
    RelationFriendListResp FriendListMethod(1: RelationFriendListReq req) (api.get="/api/v1/relation/friend/list");
}