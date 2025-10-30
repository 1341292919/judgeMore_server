namespace go user
include "./model.thrift"

struct RegisterRequest {
    1: required string username,
    2: required string password,
    3: required string email,
}

struct RegisterResponse {
    1: model.BaseResp base,
    2: optional i64 UserId,
}

//  login
struct LoginRequest{
    1: required string username,
    2: required string password,
    3: optional string captcha,
}
struct LoginResponse{
    1: model.BaseResp base,
    2: optional model.UserInfo data,
}
// logout
struct LogoutReq {
}
struct LogoutResp {
    model.BaseResp base,
}
// QueryUserInfo
struct QueryUserInfoRequest {
    1 :required i64 UserId,
}

struct QueryUserInfoResponse {
    1: required model.UserInfo data,
    2: optional model.BaseResp base,
}
// Verify Email
struct VerifyEmailRequest{
      1: required string email,
      2: required string code,
      3: required i64 id,
}
struct VerifyEmailResponse{
    1: required model.BaseResp base,
}
// UpdateUserInfo
struct UpdateUserInfoRequest{
    1: optional string college,
    2: optional string grade,
    3: optional string major,
    4: required i64 id,
}
struct UpdateUserInfoResponse{
    1: optional model.UserInfo data,
    2: required model.BaseResp base,
}
service UserService {
        RegisterResponse Register(1: RegisterRequest req),
        LoginResponse Login(1: LoginRequest req),
        LogoutResp Logout(1: LogoutReq req),
        QueryUserInfoResponse QueryUserInfo(1: QueryUserInfoRequest req),
         VerifyEmailResponse VerifyEmail(1: VerifyEmailRequest req),
         UpdateUserInfoResponse UpdateUserInfo(1: UpdateUserInfoRequest req),
}
