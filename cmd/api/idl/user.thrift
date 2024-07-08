namespace go user

include "model.thrift"

//注册
struct RegisterRequest{
    1:required string username (api.form="username"),
    2:required string password (api.form="password"),
}

struct RegisterResponse{
    1:model.BaseResp base,
    2:model.User data,
}
//登录
struct LoginRequest{
    1:required string username (api.form="username"),
    2:required string password (api.form="password"),
    3:optional string code (api.form="code"),
}

struct LoginResponse{
    1:model.BaseResp base,
    2:model.UserInfo data,
}
//用户信息
struct InfoRequest{
    1:required string user_id (api.query="user_id"),
}

struct InfoResponse{
    1:model.BaseResp base,
    2:model.UserInfo data,
}
//上传头像
struct UploadRequest{
//    1:required binary data (api.form="data"),
}

struct UploadResponse{
    1:model.BaseResp base,
    2:model.User data,

}

//获取 MFA qrcode
struct MFAGetRequest{

}

struct MFAGetResponse{
    1:model.BaseResp base,
    2:model.MFA data,

}

//绑定多因素身份认证(MFA)
struct MFABindRequest{
    1:required string code (api.form="code"),
    2:required string secret (api.form="secret"),
}

struct MFABindResponse{
    1:model.BaseResp base,
}
//关闭多因素身份认证(MFA)
struct MFAStatusRequest{
    1:required string code (api.form="code"),
    2:required string action_type (api.form="action_type"),
}

struct MFAStatusResponse{
    1:model.BaseResp base,
}

//上传图片到向量数据库
struct UploadImagesRequest{

}

struct UploadImagesResponse{
    1:model.BaseResp base,
}

//以图搜图
struct SearchImagesRequest{

}

struct SearchImagesResponse{
    1:model.BaseResp base,
    2:string data
}


service UserService{
    RegisterResponse Register(1:RegisterRequest req)(api.post="/user/register"),
    LoginResponse Login(1:LoginRequest req)(api.post="/user/login"),
    InfoResponse Info(1:InfoRequest req)(api.get="/user/info"),
    UploadResponse Upload(1:UploadRequest req)(api.put="/user/avatar/upload")
    MFAGetResponse MFAGet(1:MFAGetRequest req)(api.get="/auth/mfa/qrcode")
    MFABindResponse MFA(1:MFABindRequest req)(api.post="/auth/mfa/bind")
    MFAStatusResponse MFAStatus(1:MFAStatusRequest req)(api.post="/auth/mfa/status")
    UploadImagesResponse UploadImages(1:UploadImagesRequest req)(api.post="/image/upload")
    SearchImagesResponse SearchImages(1:SearchImagesRequest req)(api.post="/image/search")
}