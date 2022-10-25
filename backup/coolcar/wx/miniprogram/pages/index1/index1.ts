// index.ts
// 获取应用实例
//const app = getApp<IAppOption>()
// Component({
//   properties:{
//     showModal: Boolean,
//     showCancel: Boolean,
//     title: String,
//     contents: String,
//   },
//   options:{addGlobalClass:true,},
// })
Page({
  data: {
    motto: 'Hello World',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    canIUseGetUserProfile: false,
    canIUseOpenData: wx.canIUse('open-data.type.userAvatarUrl') && wx.canIUse('open-data.type.userNickName') // 如需尝试获取用户信息可改为false
  },
  // 事件处理函数
  bindViewTap() {
    wx.redirectTo({
      url: '../logs/logs',
    })
  },
  async onLoad() {
    await app.globalData.userInfo.then(userInfo => {
      this.setData({
        userInfo,
        hasUserInfo: true
      })
      // @ts-ignore
      if (wx.getUserProfile) {
        this.setData({
          canIUseGetUserProfile: true
        })
      }
    })
  },
  getUserProfile() {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },
  getUserInfo(e: any) {
    // 不推荐使用getUserInfo获取用户信息，预计自2021年4月13日起，getUserInfo将不再弹出弹窗，并直接返回匿名的用户个人信息
    console.log(e)
    const  userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo //确认e.detail.userInfo类型为具体的微信个人类型信息，不是any，起到了保护作用，这样就不会误将
    //userInfo 获取的e.detail.userInfo的值通过等号进行赋值给app.globalData.userInfo，app.globalData.userInfo = userInfo这种错误形式

    // app.globalData.userInfo = userInfo  //错误。不同的类型赋值

    //希望可以通过app.globalData.userInfo.resolve(userInfo)方式，那么就要修改app.ts里边的globalData的返回结果，可以将resolve存在app.ts中的全局变量中，供给外部结果调用
    app.resolveUserInfo(userInfo)
      this.setData({
      //userInfo: e.detail.userInfo,//此处的userInfo会与OnLoad中设置的userInfo不同，该设置类型为值类型，而OnLoad设置的是Promise类型，所以需要统一
      hasUserInfo: true
    })



    // app.globalData.userInfo.then(res => {
    //   this.setData({
    //     userInfo: e.detail.userInfo,//此处的userInfo会与OnLoad中设置的userInfo不同，该设置类型为值类型，而OnLoad设置的是Promise类型，所以需要统一
    //     hasUserInfo: true
    //   })
    // })
    // this.setData({
    //   userInfo: e.detail.userInfo,//此处的userInfo会与OnLoad中设置的userInfo不同，该设置类型为值类型，而OnLoad设置的是Promise类型，所以需要统一
    //   hasUserInfo: true
    // })
  }
})
