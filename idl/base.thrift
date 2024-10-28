namespace go base

struct User {
    1: string id;
    2: string username;
    3: string avatar_url;
    4: i64 created_at;
    5: i64 updated_at;
    6: i64 deleted_at;
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
    2: string user_id;
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
}

struct Comment {
    1: string id;
    2: string user_id;
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
}

struct Activity {
    1: string id;
    2: string user_id;
    3: string title;
    4: string text;
    5: list<string> image;
    6: string ref_video;
    7: string ref_activity;
    8: i64 created_at;
    9: i64 updated_at;
    10: i64 deleted_at;
}

struct Report {
    1: string id;
    2: string otype;
    3: string oid;
    4: string user_id;
    5: string content;
    6: string labels;
    7: i64 created_at;
    8: i64 updated_at;
    9: i64 deleted_at;
}