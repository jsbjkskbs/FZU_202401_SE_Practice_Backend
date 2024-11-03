namespace go api.report

include "base.thrift"

struct ReportVideoReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
    3: required string content;
    4: required string label;
}

struct ReportVideoResp {
    1: i64 code;
    2: string msg;
}

struct ReportActivityReq {
    1: required string access_token (api.header="Access-Token");
    2: required string activity_id;
    3: required string content;
    4: required string label;
}

struct ReportActivityResp {
    1: i64 code;
    2: string msg;
}

struct ReportCommentReq {
    1: required string access_token (api.header="Access-Token");
    2: required string comment_type;
    3: required string from_media_id;
    4: required string comment_id;
    5: required string content;
    6: required string label;
}

struct ReportCommentResp {
    1: i64 code;
    2: string msg;
}

struct AdminVideoReportListReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
    4: optional string status;
    5: optional string keyword;
    6: optional string user_id;
    7: optional string label;
}

struct AdminVideoReportListRespData {
    1: list<base.VideoReport> items;
    2: i64 total;
    3: i64 page_num;
    4: i64 page_size;
    5: bool is_end;
}
struct AdminVideoReportListResp {
    1: i64 code;
    2: string msg;
    3: AdminVideoReportListRespData data;
}

struct AdminActivityReportListReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
    4: optional string status;
    5: optional string keyword;
    6: optional string user_id;
    7: optional string label;
}

struct AdminActivityReportListRespData {
    1: list<base.ActivityReport> items;
    2: i64 total;
    3: i64 page_num;
    4: i64 page_size;
    5: bool is_end;
}
struct AdminActivityReportListResp {
    1: i64 code;
    2: string msg;
    3: AdminActivityReportListRespData data;
}

struct AdminCommentReportListReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
    4: required string comment_type;
    5: optional string status;
    6: optional string keyword;
    7: optional string user_id;
    8: optional string label;
}

struct AdminCommentReportListRespData {
    1: list<base.CommentReport> items;
    2: i64 total;
    3: i64 page_num;
    4: i64 page_size;
    5: bool is_end;
}
struct AdminCommentReportListResp {
    1: i64 code;
    2: string msg;
    3: AdminCommentReportListRespData data;
}

struct AdminVideoReportHandleReq {
    1: required string access_token (api.header="Access-Token");
    2: required string report_id;
    3: required i64 action_type;
}

struct AdminVideoReportHandleResp {
    1: i64 code;
    2: string msg;
}

struct AdminActivityReportHandleReq {
    1: required string access_token (api.header="Access-Token");
    2: required string report_id;
    3: required i64 action_type;
}

struct AdminActivityReportHandleResp {
    1: i64 code;
    2: string msg;
}

struct AdminCommentReportHandleReq {
    1: required string access_token (api.header="Access-Token");
    2: required string comment_type;
    3: required string report_id;
    4: required i64 action_type;
}

struct AdminCommentReportHandleResp {
    1: i64 code;
    2: string msg;
}

struct AdminVideoListReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
    4: optional string category;
}

struct AdminVideoListRespData {
    1: list<base.Video> items;
    2: i64 total;
    3: i64 page_num;
    4: i64 page_size;
    5: bool is_end;
}
struct AdminVideoListResp {
    1: i64 code;
    2: string msg;
    3: AdminVideoListRespData data;
}

struct AdminVideoHandleReq {
    1: required string access_token (api.header="Access-Token");
    2: required string video_id;
    3: required i64 action_type;
}

struct AdminVideoHandleResp {
    1: i64 code;
    2: string msg;
}

service ReportService {
    ReportVideoResp ReportVideo(1: ReportVideoReq req) (api.post="/api/v1/report/video");
    ReportActivityResp ReportActivity(1: ReportActivityReq req) (api.post="/api/v1/report/activity");
    ReportCommentResp ReportComment(1: ReportCommentReq req) (api.post="/api/v1/report/comment");
    AdminVideoReportListResp AdminVideoReportList(1: AdminVideoReportListReq req) (api.get="/api/v1/admin/video/report/list");
    AdminActivityReportListResp AdminActivityReportList(1: AdminActivityReportListReq req) (api.get="/api/v1/admin/activity/report/list");
    AdminCommentReportListResp AdminCommentReportList(1: AdminCommentReportListReq req) (api.get="/api/v1/admin/comment/report/list");
    AdminVideoReportHandleResp AdminVideoReportHandle(1: AdminVideoReportHandleReq req) (api.post="/api/v1/admin/video/report/handle");
    AdminActivityReportHandleResp AdminActivityReportHandle(1: AdminActivityReportHandleReq req) (api.post="/api/v1/admin/activity/report/handle");
    AdminCommentReportHandleResp AdminCommentHandle(1: AdminCommentReportHandleReq req) (api.post="/api/v1/admin/comment/report/handle");
    AdminVideoListResp AdminVideoList(1: AdminVideoListReq req) (api.get="/api/v1/admin/video/list");
    AdminVideoHandleResp AdminVideoHandle(1: AdminVideoHandleReq req) (api.post="/api/v1/admin/video/handle");
}