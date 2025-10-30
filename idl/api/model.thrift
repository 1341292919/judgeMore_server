namespace go api.model

struct BaseResp{
    1: i64 code,
    2: string msg,
}

struct UserInfo{
    1: i64 username,  //姓名
    2: i64 userId,   // 学号
    4: string Major // 专业
    5: string college, //学院
    6: string grade,  // 年级
}