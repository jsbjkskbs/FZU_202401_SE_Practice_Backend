namespace go base

struct User {
    1: string id;
    2: string username;
    3: string avatar_url;
    4: i64 created_at;
    5: i64 updated_at;
    6: i64 deleted_at;
    7: bool is_followed;
}

struct UserWithToken {
    1: string id;
    2: string username;
    3: string avatar_url;
    4: i64 created_at;
    5: i64 updated_at;
    6: i64 deleted_at;
    7: string access_token;
    8: string refresh_token;
    9: string role;
}

struct Video {
    1: string id;
    2: User user;
    3: string video_url;
    4: string cover_url;
    5: string title;
    6: string description;
    7: i64 visit_count;
    8: i64 like_count;
    9: i64 comment_count;
    10: string category;
    11: list<string> labels;
    12: string status;
    13: i64 created_at;
    14: i64 updated_at;
    15: i64 deleted_at;
    16: bool is_liked;
}

struct Comment {
    1: string id;
    2: User user;
    3: string otype;
    4: string oid;
    5: string root_id;
    6: string parent_id;
    7: i64 like_count;
    8: i64 child_count;
    9: string content;
    10: i64 created_at;
    11: i64 updated_at;
    12: i64 deleted_at;
    13: bool is_liked;
}

struct Activity {
    1: string id;
    2: User user;
    3: string content;
    4: list<string> image;
    5: string ref_video;
    6: string ref_activity;
    7: i64 like_count;
    8: i64 comment_count;
    9: i64 created_at;
    10: i64 updated_at;
    11: i64 deleted_at;
    12: bool is_liked;
}

struct VideoReport {
    1: string id;
    2: string video_id;
    3: string user_id;
    4: string reason;
    5: string label;
    6: string status;
    7: i64 created_at;
    8: i64 resolved_at;
    9: string admin_id;
}

struct ActivityReport {
    1: string id;
    2: string activity_id;
    3: string user_id;
    4: string reason;
    5: string label;
    6: string status;
    7: i64 created_at;
    8: i64 resolved_at;
    9: string admin_id;
}

struct CommentReport {
    1: string id;
    2: string comment_type;
    3: string comment_id;
    4: string user_id;
    5: string reason;
    6: string label;
    7: string status;
    8: i64 created_at;
    9: i64 resolved_at;
    10: string admin_id;
}