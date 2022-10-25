// index.ts

import { IAppOption } from "../../appoption"
import { CarService } from "../../service/car"
import { ProfileService } from "../../service/profile"
//import { car } from "../../service/proto_gen/car/car_pb"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

// 获取应用实例
const app = getApp<IAppOption>()

interface Marker {
      iconPath:string,
      id:number,
      latitude:number,
      longitude:number,
      width:number,
      height:number
}

const defaultAvatar = "/resources/car.png"
const  initialLat = 30
const  initialLng = 120


Page({
  socket:undefined as WechatMiniprogram.SocketTask |undefined,
  ifPageShowing:false,
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
      latitude:initialLat,
      longitude:initialLng,
    },
    scale:17,
    // markers:[{
    //   iconPath:"/resources/car.png",
    //   id:0,
    //   latitude:23.099995,
    //   longitude:113.324520,
    //   width:50,
    //   height:50
    // },{
    //   iconPath:"/resources/car.png",
    //   id:1,
    //   latitude:23.099995,
    //   longitude:114.324520,
    //   width:50,
    //   height:50
    // }],
    markers:[] as Marker[],
  },
  onMyLocationTap(){
    wx.getLocation({
      type:"gcj02",
      success:res=>{
        this.setData({
          location:{
            latitude:res.latitude,
            longitude:res.longitude
          },
        })
      },
      fail:()=>{
        wx.showToast({
          title:'请前往设置页授权',
          icon:'none'
        })
      }
    })
  },
  async onLoad(){
    
    // this.socket = wx.connectSocket({
    //   url:'ws://localhost:9090/ws',
    // })


    // let msgCount = 0
    // this.socket = CarService.subscribe(function (msg: car.v1.ICarEntity) {
    
    //     msgCount++
    //     console.log(msg)
    // })

    // this.socket.onMessage(function (msg: WechatMiniprogram.SocketTaskOnMessageCallbackResult) {
    
    //     msgCount++
    //     console.log(msg)
    // })


    // setInterval(()=>{
    //   this.socket?.send({
    //     data:JSON.stringify({
    //       msg_received:msgCount
    //     })
    //   })
    // },3000)

    // wx.request({
    //   url:"http://localhost:8080/trip/123",
    //   method:"GET",
    //   success:console.log,
    //   fail:console.error
    // })
    const userInfo = await app.globalData.userInfo
    this.setData({
      avatarURL:userInfo?.avatarUrl
    })
  },
  onShow(){
    this.ifPageShowing = true
    if (!this.socket) {
      this.setData({
        markers:[]
      },()=>{
        this.setupCarPosUpdater()
      })
    }
  },
  onHide(){
    this.ifPageShowing = false
    if (this.socket) {
      this.socket.close({
        success:()=>{
          this.socket = undefined
        }
      })
    }
  },
  setupCarPosUpdater(){
    const map = wx.createMapContext("map")
    const markersByCarID = new Map<string,Marker>()
    const  endTranslation =  ()=>{
      translationInProgress = false
    }
    //避免设置信息时间过长,多条信息碰撞
    let translationInProgress = false
    this.socket = CarService.subscribe(car=>{
      if (!car.id || translationInProgress || !this.ifPageShowing) {
        console.log("dropped")
        return
      }
      const  marker = markersByCarID.get(car.id)
      if (!marker) {
        //new create car and  add markers
        const newMarker: Marker = {
          id:this.data.markers.length,
          iconPath:car.car?.driver?.avatarUrl || defaultAvatar,
          latitude: car.car?.position?.latitude || initialLat,
          longitude: car.car?.position?.longitude || initialLng,
          height:20,
          width:20,
        }
        markersByCarID.set(car.id,newMarker)
        this.data.markers.push(newMarker)
        translationInProgress  = true
        this.setData({
          markers:this.data.markers,
        },endTranslation)
        return
      }

      const newAvatar = car.car?.driver?.avatarUrl || defaultAvatar
      const newLat = car.car?.position?.latitude || initialLat
      const newLng = car.car?.position?.longitude || initialLng
      if (marker.iconPath != newAvatar) {
        //chekc information is deff and reset data
        marker.iconPath = newAvatar
        marker.latitude = newLat
        marker.longitude = newLng
        translationInProgress  = true
        this.setData({
          markers:this.data.markers
        },endTranslation)
      }


      if (marker.latitude  !== newLat ||  marker.longitude !== newLng) {
        //Move marker
        translationInProgress  = true
        map.translateMarker({
          markerId:marker.id,
          destination:{
            latitude:newLat,
            longitude:newLng,
          },
          autoRotate: false,
          rotate:0,
          duration:90,
          animationEnd:endTranslation
        })
      }

    })
  },
  moveCars(){
    const map = wx.createMapContext("map")
    const dest = {
      latitude:23.099995,
      longitude:113.324520
    }

    const moveCars = () => {
      dest.latitude += 0.1
      dest.longitude += 0.1
      map.translateMarker({
        destination: {
          latitude:dest.latitude,
          longitude:dest.longitude,
        },
        markerId:0,
        autoRotate:false,
        rotate:0,
        duration:5000,
        animationEnd: ()=>{
          if(this.ifPageShowing) {
            moveCars()
          }
        },
      })
    }
    moveCars()
  },
  async onScanTap(){
    const trips = await TripService.GetTrips(rental.v1.TripStatus.IN_PROGRESS)
    if((trips.trips?.length || 0) > 0){
      await this.selectComponent('#tripModal').showModal()
      wx.navigateTo({
        url:routing.driving({
          trip_id: trips.trips![0].id!,
        }),
      })
      return
    } 

    wx.scanCode({
      success:async ()=>{
        
        const carID = '62b6e446d7ed5c8e33f64ba2'
        //const redirectURL = `/pages/lock/lock?car_id=${carID}`
        const lockURL = routing.lock({
          car_id:carID
        })
        const prof = await ProfileService.getProfile()

        if(prof.identityStatus === rental.v1.IdentityStatus.VERIFIED){
          console.log(1232132141)
          wx.navigateTo({
            url:lockURL,
          })
        }else{
          console.log(678967867867)
          await this.selectComponent('#licModal').showModal()
          wx.navigateTo({
            //url:`/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`,
            url:routing.redirect({
              redirectURL:lockURL
            })
          })
        }
      },
      fail:console.error
      // fail:() => {
      //   const carID = 'car123'
      //   //const redirectURL = `/pages/lock/lock?car_id=${carID}`
      //   const redirectURL = routing.lock({
      //     car_id:carID
      //   })
      //   console.log(83021983012038102)
      //   wx.navigateTo({
      //     // url:`/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`,
      //     url:routing.redirect({
      //       redirectURL:redirectURL
      //     }),
      //   })
      // }
    })
  },
  onMyTripsTap() {
    // this.socket?.close({
    //   code:1006,
    //   complete:res=>{
    //     console.log(res)
    //   }
    // })
    wx.navigateTo({
      url: routing.mytrips(),
    })
  },
})
