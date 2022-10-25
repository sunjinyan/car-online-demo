import { IAppOption } from "../../appoption"
import { CarService } from "../../service/car"
import { car } from "../../service/proto_gen/car/car_pb"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

// pages/lock/lock.ts
const shareLocationKey = "share_location"
Page({
    carID:'',
    carRefresher:0,
    /**
     * 页面的初始数据
     */
    data: {
        userInfo:{},
        avatarURL:'',
        shareLocationKey:false,
        shareLocation:""
    },

    /**
     * 生命周期函数--监听页面加载
     */
    async onLoad(opt:Record<'car_id',string>) {
        const o: routing.LockOpts = opt
        
        this.carID = o.car_id
        console.log(o.car_id)
        
        const userInfo = await getApp<IAppOption>().globalData.userInfo
        wx.getStorageSync(shareLocationKey)
        this.setData({
            avatarURL:userInfo?.avatarUrl,
            shareLocation:wx.getStorageSync(shareLocationKey)  || false
        })
    },
    onUnlockTap(){
        wx.getLocation({
            type:'gcj02',
            success: async loc=>{
                console.log('starting a trip',{
                    location:{
                        latitude: loc.latitude,
                        longitude:loc.longitude
                    },
                    //todo: 需要双向绑定
                    avatarURL:this.data.shareLocationKey? this.data.avatarURL:'',
                    carID:''
                })
                
                if (!this.carID) {
                    console.error("no carID specified")
                    return
                }

                //console.log("============================",loc)
                let trip: rental.v1.ITripEntity
                try {
                    trip = await TripService.CreateTrip({
                        //start:"abc"
                        start:{
                            latitude:loc.latitude,
                            longitude:loc.longitude
                        },
                        carId:this.carID
                    })

                    if (!trip.id) {
                        console.error("no trip id specified")
                        return
                    }
                } catch (err) {
                    wx.showToast({
                        title:'创建行程失败',
                        icon:'none',
                    })
                    return
                }

                
                //return
                

                //const tripID = 'trip456'
                wx.showLoading({
                    title:"开锁中",
                    mask:true,
                })
                

                //检查车的状态并开锁

                this.carRefresher = setInterval(async ()=>{
                    const c = await CarService.getCar(this.carID)
                    if (c.status === car.v1.CarStatus.UNLOCKED){
                        this.clearCarRefresher()
                        wx.redirectTo({
                            //url: `/pages/driving/driving?trip_id=${tripID}`,
                            url: routing.driving({
                                trip_id:trip.id!
                            }),
                            complete:()=>{
                                wx.hideLoading()
                            }
                        })
                    }
                },2000)

                // setTimeout(() => {
                //     wx.redirectTo({
                //         //url: `/pages/driving/driving?trip_id=${tripID}`,
                //         url: routing.driving({
                //             trip_id:trip.id!
                //         }),
                //         complete:()=>{
                //             wx.hideLoading()
                //         }
                //     })
                // },2000)
            },
            fail:()=>{
                wx.showToast({
                    icon:'none',
                    title:'请前往设置页授权位置信息'
                })
            }
        })
    },
    clearCarRefresher(){
        if(this.carRefresher){
            clearInterval(this.carRefresher)
            this.carRefresher = 0
        }
    },
    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {
        this.clearCarRefresher()
        wx.hideLoading()
    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    },
    onGetUserInfo(e:any){
        console.log(e)
        const userInfo: WechatMiniprogram.UserInfo = e.detail.value.userInfo
        if(e.detail.value.userInfo){
            getApp<IAppOption>().resolveUserInfo(userInfo)
        }
    },
    onShareLocation(e:any){
        this.data.shareLocation = e.detail.value
        //const shareLocation:boolean = e.detail.value
        wx.setStorageSync(shareLocationKey,this.data.shareLocation)
    }
})