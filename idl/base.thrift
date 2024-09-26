namespace go base

struct Empty {}

struct PaginationReq {
    1: i32 page = 1 (go.tag='json:"page" query:"page"')
    2: i32 page_size = 20 (go.tag='json:"page_size" query:"page_size"')
}

struct PaginationResp {
    1: required i32 page
    2: required i32 page_size
    3: required i32 total
    4: required i32 page_total
}

struct IDReq {
    1: required i64 id (go.tag='path:"id"')
}

struct IDsReq {
    1: required list<i64> ids
}

// The Time type is encapsulated to encapsulate unified time processing logic
struct Time {
    1: optional i64 timestamp
}

struct Duration {
    1: optional i64 milliseconds
}