// pages/driving/driving.ts

import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { formatDuration, formatFee } from "../../utils/format"
import { routing } from "../../utils/routing"

const updateIntervalSec = 5

function durationStr(sec:number) {
    const dur = formatDuration(sec)
    return `${dur.hh}:${dur.mm}:${dur.ss}`
}

// function formatDuration(sec:number){

//     const padString = (n:number)=>n<10?'0'+n.toFixed(0):n.toFixed(0)

//     const h = Math.floor(sec/3600)
//     sec -= 3600*h
//     const m = Math.floor(sec/60)
//     sec -= 60 * m
//     const s = Math.floor(sec)
//     return `${padString(h)}:${padString(m)}:${padString(s)}`
// }

// function formatFee(cents:number) {

//     return (cents/100).toFixed(2)
// }

Page({
    tripID:"",

    timer: undefined as  number | undefined,
    /**
     * 页面的初始数据
     */
    data: {
        location:{
            latitude: 32.92,
            longitude: 118.56,
        },
        scale: 12,
        elapsed: '00:00:00',
        fee: '0.00',
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(opt:Record<'trip_id',string>) {
        const o:routing.DrivingOpts  =  opt
        console.log('current trip',o.trip_id)
        this.tripID = o.trip_id
        TripService.GetTrip(o.trip_id).then(console.log)
        this.setupLocationUpdate()
        this.setupTimer(o.trip_id)
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
        wx.stopLocationUpdate()
        if(this.timer){
            clearInterval(this.timer)
        }
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
    setupLocationUpdate(){
        wx.startLocationUpdate({
            fail:err=>{
                console.log(err)
            }
        })
        wx.onLocationChange(loc=>{
            console.log('location:',loc)
            this.setData({
                location:{
                    latitude:loc.latitude,
                    longitude:loc.longitude,
                }
            })
        })
    },
    async setupTimer(trip_id:string){
        const trip = await TripService.updateTripPos(trip_id)
        if (trip.status !==  rental.v1.TripStatus.IN_PROGRESS) {
            console.error("trip not in progress")
            return
        }
        let secSinceLastUpdate = 0
        let lastUpdateDurationSec = trip.current!.timestampSec!

        this.setData({
            elapsed:durationStr(lastUpdateDurationSec),
            fee:formatFee(trip.current!.feeCent!) 
        })

        //let cents = 0
        this.timer = setInterval(()=>{
            secSinceLastUpdate++
            //cents += centPerSec
            if (secSinceLastUpdate % updateIntervalSec === 0) {
                TripService.updateTripPos(trip_id,{
                    latitude:this.data.location.latitude,
                    longitude:this.data.location.longitude,
                }).then(trip => {
                    lastUpdateDurationSec = trip.current!.timestampSec! - trip.start!.timestampSec!
                    secSinceLastUpdate = 0
                    this.setData({
                        fee:formatFee(trip.current!.feeCent!)//如果地道点，可以使用?替代!然后在前边做出判断
                    })
                }).catch(console.error)
            }
            this.setData({
                elapsed:durationStr(lastUpdateDurationSec + secSinceLastUpdate),
                //fee:formatFee(cents)
            })
        },1000)
    },
    onEndTripTap(){
        TripService.finishTrip(this.tripID).then(()=>{
            wx.redirectTo({
                url:routing.mytrips(),
            })
        }).catch(err=>{
            console.log(err)
            wx.showToast({
                title:"结束失败",
                icon:"none",
            })
        })
    }
})