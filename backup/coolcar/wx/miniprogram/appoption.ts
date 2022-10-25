export interface IAppOption {
    globalData: {
      userInfo: Promise<WechatMiniprogram.UserInfo>,//为了保持网络信息同步返回，页面与所需数据顺序加载，所以需要将结果使用Promise来进行操作
    }
    userInfoReadyCallback?: WechatMiniprogram.GetUserInfoSuccessCallback,
    resolveUserInfo(userInfo: WechatMiniprogram.UserInfo):void//将此处接口与app里实现的保持统一，app里需要这种，所以在这里定义
    rejectUserInfo(userInfo: WechatMiniprogram.UserInfo):void
  }