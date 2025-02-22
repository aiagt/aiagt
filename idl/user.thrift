namespace go usersvc

include './base.thrift'

struct User {
    1: required i64 id
    2: required string username
    3: required string email
    4: required string phone_number
    5: required string signature
    6: required string homepage
    7: required string description_md
    8: required string github
    9: required string avatar
    10: required base.Time created_at
    11: required base.Time updated_at
}

struct Secret {
    1: required i64 id
    2: required i64 user_id
    3: required i64 plugin_id
    4: required string name
    5: required string value
    6: required base.Time created_at
    7: required base.Time updated_at
    8: optional string plugin_name
    9: optional string plugin_logo
}

struct RegisterReq {
    1: required string email
    2: required string captcha
    3: optional string username
    4: optional string password
}

struct RegisterResp {
    1: required string token
    2: required base.Time expire
    3: required User user
}

struct LoginReq {
    1: required string email
    2: optional string password
    3: optional string captcha
}

struct LoginResp {
    1: required string token
    2: required base.Time expire
    3: required User user
}

struct UpdateUserReq {
    1: optional string username
    2: optional string email
    3: optional string phone_number
    4: optional string signature
    5: optional string homepage
    6: optional string description_md
    7: optional string github
    8: optional string avatar
}

struct ResetPasswordReq {
    1: required string email
    2: required string captcha
    3: required string password
}

struct CreateSecretReq {
    1: required i64 plugin_id
    2: required string name
    3: required string value
}

struct UpdateSecretReq {
    1: required i64 id (go.tag='path:"id"')
    2: optional i64 plugin_id
    3: optional string name
    4: optional string value
}

struct ListSecretReq {
    1: required base.PaginationReq pagination
    2: optional i64 plugin_id (go.tag='query:"plugin_id"')
    3: optional string name (go.tag='query:"name"')
}

struct ListSecretResp {
    1: required base.PaginationResp pagination
    2: required list<Secret> secrets
}

enum CaptchaType {
    AUTH
    RESET
}

struct SendCaptchaReq {
    1: required string email
    2: required CaptchaType type
}

struct SendCaptchaResp {
    1: required bool exists
}

service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(1: LoginReq req)
    i64 ParseToken(1: string token)
    base.Empty ResetPassword(1: ResetPasswordReq req)
    SendCaptchaResp SendCaptcha(1: SendCaptchaReq req)

    base.Empty UpdateUser(1: UpdateUserReq req)
    User GetUser()
    User GetUserByID(1: base.IDReq req)
    list<User> GetUserByIds(1: base.IDsReq req)

    base.Empty CreateSecret(1: CreateSecretReq req)
    base.Empty UpdateSecret(1: UpdateSecretReq req)
    base.Empty DeleteSecret(1: base.IDReq req)
    ListSecretResp ListSecret(1: ListSecretReq req)
}