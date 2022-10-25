import { IAppOption } from "../../appoption"
import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/trip"
import { formatDuration, formatFee } from "../../utils/format"
import { routing } from "../../utils/routing"
interface Trip{
    id: string
    shortId:string
    start: string
    end:string
    duration:string
    fee:string
    distance:string
    status:string
}

interface MainItem {
    id: string
    navId: string
    navScrollId:string
    data: Trip
}

interface NavItem{
    id:string
    mainId:string
    label:string
}

interface MainItemQueryResult {
    id: string
    top: number
    dataset: {
        navId: string
        navScrollId: string
    }
}

const tripStatusMap = new Map([
    [rental.v1.TripStatus.IN_PROGRESS,'进行中'],
    [rental.v1.TripStatus.FINISHED,'已完成'],
])
const licStatusMap = new Map([
    [rental.v1.IdentityStatus.UNSUBMITTED,'未认证'],
    [rental.v1.IdentityStatus.PENDING,'认证中'],
    [rental.v1.IdentityStatus.VERIFIED,'已认证'],
])

// pages/mytrips/mytrips.ts
Page({
    scrollStates:{
        mainItems:[] as MainItemQueryResult[]
    },
    layoutResolver: undefined as ((value: unknown) => void)|undefined,
    // layoutReady:{
    //     promise:undefined as Promise<void>|undefined,
    //     resolver:undefined as (()=>void)|undefined
    // },
    /**
     * 页面的初始数据
     */
    data: {
        licStatus:licStatusMap.get(rental.v1.IdentityStatus.UNSUBMITTED),
        avatarURL:"",
        indicatorDots:true,
        autoPlay:true,
        interval:3000,
        duration:500,
        circular:true,
        multiItemCount:1,
        prevMargin:'',
        nextMargin:'',
        vertical:false,
        current:0,
        promotionItems:[
            {
                img:'https://img1.mukewang.com/626603a70001b02417920764.jpg',
                promotionID:1
            },
            {
                img:'https://img2.mukewang.com/626791df000107ec17920764.jpg',
                promotionID:2
            },
            {
                img:'https://img.mukewang.com/625f680f0001000917920764.jpg',
                promotionID:3
            },
            {
                img:'https://img4.mukewang.com/626605620001e2e617920764.jpg',
                promotionID:4
            }
        ],
        mainItems:[] as MainItem[],
        navItems: [] as NavItem[],
        tripsHeight:0,
        //scrollTop:0,
        //scrollIntoView:''
        mainScroll:'',
        navCount:1,
        navSel:'',
        navScroll:'',
    },
    onRegisterTap(){
        wx.navigateTo({
            url:routing.redirect(),//'/pages/register/register'
        })
    },
    //onSwiperChange(e:any){
    onSwiperChange(){
        //console.log(e)
    },
    onPromotionItemTap(e:any){
        console.log(e)
        const promotionID = e.currentTarget.dataset.promotionId
        if(promotionID){
            console.log(promotionID)
        }
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {
       const layoutReady = new Promise((resolve)=>{
            this.layoutResolver = resolve
        })
        Promise.all([TripService.GetTrips(),layoutReady]).then(([trips])=>{
            this.populateTrips(trips.trips!)
        })
        //const [res] = await Promise.all([TripService.GetTrips(),layoutReady])
        //console.log(res)
        //this.populateTrips(res.trips!)
        //const userInfo = await getApp<IAppOption>().globalData.userInfo
        getApp<IAppOption>().globalData.userInfo?.then(userInfo=>{
            this.setData({
                avatarURL:userInfo?.avatarUrl
            })
        })
    },
    populateTrips(trips:rental.v1.ITripEntity[]){
        const mainItems: MainItem[] = []
        const navItems: NavItem[] = []
        let navSel = ''
        let prevNav = ''
        for(let i = 0; i < trips.length; i++){
            const trip = trips[i]
            const mainId = 'main-' + i
            const navId = 'nav-' + i
            //const tripId = (10001+i).toString()
            const shortId = trip.id?.substring(trip.id.length-6)
            
            if (!prevNav){
                prevNav = navId
            }
            // distance:'27.0公里',
            //         duration:'0时44分',
            //         fee:'128.00元',
            //         status:'已完成'
            const tripData:Trip = {
                id:trip.id!,
                    shortId:'***'+shortId!,
                    start:trip.trip?.start?.poiName||'未知',
                    end:'',
                    distance:'',
                    duration:'',
                    fee:'',
                    status:tripStatusMap.get(trip.trip?.status!)||'未知'
            }
            const end = trip.trip?.end
            if (end) {
                tripData.end = end.poiName||'未知'
                tripData.distance = end.kmDriven?.toFixed(1)+'公里'
                tripData.fee = formatFee(end.feeCent || 0)
                const dur = formatDuration((end.timestampSec || 0) - (trip.trip?.start?.timestampSec! || 0))
                tripData.duration = `${dur.hh}时${dur.mm}分`
            }
            mainItems.push({
                id:mainId,
                navId: navId,
                navScrollId:prevNav,
                // data:{
                //     id:trip.id!,
                //     shortId:'***'+shortId!,
                //     start:trip.trip?.start?.poiName||'未知',
                //     end:'迪士尼',
                //     distance:'27.0公里',
                //     duration:'0时44分',
                //     fee:'128.00元',
                //     status:'已完成'
                // }
                data:tripData
            })
            navItems.push({
                id: navId,
                mainId:mainId,
                label:shortId||''
            })
            if(i === 0){
                navSel = navId
            }
            prevNav = navId
        }
        console.log('nav count:',this.data.navCount)
        for (let i = 0; i < this.data.navCount-1; i++) {
            navItems.push({
                id:'',
                mainId:'',
                label:'',
            })
        }
        this.setData({
            mainItems,
            navItems,
            navSel
        },()=>{
            this.prepareScrollStates()
        })
    },
    prepareScrollStates(){
        wx.createSelectorQuery().selectAll('.main-item').fields({
            id:true,
            dataset:true,
            rect:true,
        }).exec(res => {
            this.scrollStates.mainItems = res[0]
        })
    },
    onMainScroll(e: any){
        //console.log(e)
        const top: number = e.currentTarget?.offsetTop+e.detail?.scrollTop
        
        if(top === undefined){
            return
        }

       const  selItem =  this.scrollStates.mainItems.find(v=>v.top>=top)
        if (!selItem){
            return
        }
        this.setData({
            navSel:selItem.dataset.navId,
            navScroll:selItem.dataset.navScrollId
        })
    },
    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {
        wx.createSelectorQuery().select('#heading').boundingClientRect(rect => {
            const height =  wx.getSystemInfoSync().windowHeight - rect.height
            this.setData({
                tripsHeight: height,
                navCount:Math.round(height/50)
            },()=>{
                if (this.layoutResolver) {
                    this.layoutResolver('')
                }
            })
        }).exec()
    },
    onNavItemTap(e:any){
        const mainId: string = e.currentTarget?.dataset?.mainId
        const navId: string = e.currentTarget?.id
        if(mainId){
            this.setData({
                mainScroll:mainId,
                navSel:navId
            })
        }
    },
    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        ProfileService.getProfile().then(p=>{
            this.setData({
                licStatus: licStatusMap.get(p.identityStatus || 0)
            })
        })
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

    }
})