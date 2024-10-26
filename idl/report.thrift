namespace go api.report

include "base.thrift"

struct ReportReq {
    1: required string access_token (api.header="Access-Token");
    2: required string otype;
    3: required string oid;
    4: required string content;
    5: required list<string> labels;
}

struct ReportResp {
    1: i64 code;
    2: string msg;
}

struct AdminReportListReq {
    1: required string access_token (api.header="Access-Token");
    2: required i64 page_num;
    3: required i64 page_size;
}

struct AdminReportListRespData {
    1: list<base.Report> items;
}
struct AdminReportListResp {
    1: i64 code;
    2: string msg;
    3: AdminReportListRespData data;
}

struct AdminReportHandleReq {
    1: required string access_token (api.header="Access-Token");
    2: required string report_id;
    3: required i64 action_type;
    4: required string content;
}

struct AdminReportHandleResp {
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
    4: required string content;
}

struct AdminVideoHandleResp {
    1: i64 code;
    2: string msg;
}

service ReportService {
    ReportResp Report(1: ReportReq req) (api.post="/api/v1/report");
    AdminReportListResp AdminReportList(1: AdminReportListReq req) (api.get="/api/v1/admin/report/list");
    AdminReportHandleResp AdminReportHandle(1: AdminReportHandleReq req) (api.post="/api/v1/admin/report/handle");
    AdminVideoListResp AdminVideoList(1: AdminVideoListReq req) (api.get="/api/v1/admin/video/list");
    AdminVideoHandleResp AdminVideoHandle(1: AdminVideoHandleReq req) (api.post="/api/v1/admin/video/handle");
}