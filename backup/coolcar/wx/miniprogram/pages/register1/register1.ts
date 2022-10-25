// pages/register/register.ts
Page({

    /**
     * 页面的初始数据
     */
    data: {

    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad() {
        interface PageInfo {
          title: string;
        }
          
        type Page = "home" | "about" | "contact";
        //x为Record<Page, PageInfo>，而该类型是将Page中的每一种类型对应的值的类型都变成PageInfo类型，并生成一个对象
        //也就是使用Page中的所有类型的值为key，类似上述枚举中的类型，那么就是以上几种类型为key值，PageInfo类型为value生成一个对象，该对象就是Record<Page, PageInfo>类型的值
        const x: Record<Page, PageInfo> = {
          about: { title: "about" },
          contact: { title: "contact" },
          home: { title: "home" }
        };
        console.log(x) 

         
        type  carId =  "car_id";
        const opt: Record<carId,string> = {//泛型的第一个参数只能是string | number | symbol类型之一的值，第二个参数可以使任意内置类型或自定义类型
            "car_id": "aaaa",
        }
        function onLoad(optRecord: Record<carId,string>){
            console.log(optRecord)    
        }   
        onLoad(opt)
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
        wx.getLocation({
            success: res => {
                console.log(res)
            }
        })
        
        
        //微信告诉我位置变动了，能实时更新，而不是隔几秒我主动去申请
        wx.startLocationUpdate({
            
        })//前台更新
        wx.startLocationUpdateBackground({

        })//后台更新

        //上述的两个位置更新变动调用不需要操作什么，而是需要其他的一个函数
        wx.onLocationChange(loc => {
            this.setData({
                a : loc.latitude,
                b : loc.longitude
            })
        })
        wx.stopLocationUpdate({})
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