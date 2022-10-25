export const formatTime = (date: Date) => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return (
    [year, month, day].map(formatNumber).join('/') +
    ' ' +
    [hour, minute, second].map(formatNumber).join(':')
  )
}

const formatNumber = (n: number) => {
  const s = n.toString()
  return s[1] ? s : '0' + s
}



export  function getUserSetting(): Promise<WechatMiniprogram.GetSettingSuccessCallbackResult>{

  return new Promise((resolve,reject)=>{
    wx.getSetting({
      //success: res =>  resolve(res),
      //fail:err => reject(err)
      success: resolve,//调用与外界接连的成功的通路反馈结果
      fail:reject //调用与外界接连的失败的通路反馈结果
    })
  })

}

export  function getUserInfo(): Promise<WechatMiniprogram.GetUserInfoSuccessCallbackResult>{

  return new Promise((resolve,reject)=>{
    wx.getUserInfo({
      //success: res =>  resolve(res),
      //fail:err => reject(err)
      success: resolve,//调用与外界接连的成功的通路反馈结果
      fail:reject //调用与外界接连的失败的通路反馈结果
    })
  })

}