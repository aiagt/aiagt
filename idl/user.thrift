namespace go usersvc

include './base.thrift'

struct User {
    1: required i64 id
    2: required string username
    3: required string password
    4: required string email
    5: required string phone_number
    6: required string signature
    7: required string homepage
    8: required string description_md
    9: required string github
    10: required i64 created_at
    11: required i64 updated_at
}

struct Secret {
    1: required i64 id
    2: required i64 user_id
    3: required i64 plugin_id
    4: required i64 name
    5: required i64 value
    6: required i64 created_at
    7: required i64 updated_at
}

struct RegisterReq {
    1: required string email
    2: required string password
    3: required i32 captcha
}

struct RegisterResp {
    1: required string token
    2: required i64 id
}

struct LoginReq {
    1: required string email
    2: required string password
    3: required i32 captcha
}

struct LoginResp {
    1: required string token
    2: required string id
}

struct UpdateUserReq {
    1: required i64 id
    2: required string username
    3: required string password
    4: required string email
    5: required string phone_number
    6: required string signature
    7: required string homepage
    8: required string description_md
    9: required string github
}

struct ForgotPasswordReq {
    1: required string email
    2: required i32 captcha
    3: required string new_password
}

struct ForgotPasswordResp {
    1: required string token
    2: required i64 id
}

struct CreateSecretReq {
    1: required i64 user_id
    2: required i64 plugin_id
    3: required i64 name
    4: required i64 value
}

struct UpdateSecretReq {
    1: required i64 id
    2: required i64 user_id
    3: required i64 plugin_id
    4: required i64 name
    5: required i64 value
}

struct ListSecretReq {
    1: required base.PaginationReq pagination
    2: optional i64 plugin_id
}

struct ListSecretResp {
    1: required base.PaginationResp pagination
    2: required list<Secret> secrets;
}

service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(1: LoginReq req)
    ForgotPasswordResp ForgotPassword(1: ForgotPasswordReq req)

    base.Empty UpdateUser(1: UpdateUserReq req)
    User GetUserByID(1: base.IDReq req)

    base.Empty CreateSecret(1: CreateSecretReq req)
    base.Empty UpdateSecret(1: UpdateSecretReq req)
    base.Empty DeleteSecret(1: base.IDReq req)
    ListSecretResp ListSecret(1: ListSecretReq req)
}