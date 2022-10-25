// 获取应用实例
const app = getApp<IAppOption>()


Page({
  isPageShowing: false,
  data:{
    avatarURL:"",
    setting:{
      skew:0,
      rotate:0,
      showLocation:true,
      showScale:true,
      subKey:'',
      layerStyle:-1,
      enableZoom:true,
      enableScroll:true,
      enableRotate:false,
      showCompass:false,
      enable3D:false,
      enableOverlooking:false,
      enanleStatellite:false,
      enableTraffic:false,
    },
    location:{
      latitude:31,
      longitude:120,
    },
    scale:10,
    markers:[{
      iconPath:"/resources/car.png",
      id:0,
      latitude:23.099995,
      longitude:113.324520,
      width:50,
      height:50
    },{
      iconPath:"/resources/car.png",
      id:1,
      latitude:23.099995,
      longitude:114.324520,
      width:50,
      height:50
    }],
  },
  onLoad(){
    //this.moveCars()
  },
  onMyLocationTap(){
    wx.getLocation({
      type: 'gcj02',
      success:res => {
        this.setData({
          location:{
            latitude: res.latitude,
            longitude: res.longitude
          }
        })
      },
      fail: () => {
        wx.showToast({
          title:"请前往设置页进行授权",
          icon:"none",
          duration: 3000
        })
      }
    })
  },
  onShow(){
    this.isPageShowing = true
  },
  onHide(){
    this.isPageShowing = false
  },
  moveCars(){
    const map  =  wx.createMapContext("map")
    const pos  = {
      latitude:23.099995,
      longitude:114.324520,
    }

    const  moveCar = ()=>{
      pos.latitude += 0.1
      pos.longitude += 0.1
      map.translateMarker({
        markerId:0,
        autoRotate: false,
        destination: {
          latitude: pos.latitude,
          longitude: pos.longitude,
        },
        rotate: 0,
        duration: 3000,
        animationEnd: ()=>{
          if (this.isPageShowing) {          
            moveCar() 
          }
        }
      })
   }
   moveCar()
  },
  onScanTap(){
    wx.scanCode({
      success: () => {
        wx.redirectTo({
          url: '/pages/register/register'
        })
      },
      fail:console.error
    })
  }
})