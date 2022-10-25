import { ProfileService } from "../../service/profile"
import { rental } from "../../service/proto_gen/rental/rental_pb"
import { Coolcar } from "../../service/request"
import { padString } from "../../utils/format"
import { routing } from "../../utils/routing"
function formatDate(millis:number) {
    const dt = new Date(millis)
    const y = dt.getFullYear()
    const m = dt.getMonth() + 1
    const d = dt.getDate()

    return `${padString(y)}-${padString(m)}-${padString(d)}`
}
// pages/register/register.ts
Page({
    redirectURL:'',
    profileRefresher:0,
    /**
     * 页面的初始数据
     */
    data: {
        genderIndex:0,
        //licImgURL: "/resources/sedan.png" as  string | undefined,
        licImgURL: '',
        genders:['未知','男','女','其他'],
        birthday:'1990-01-01',
        licNo:'112321',
        name:'',
        //state:'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED'
        state: rental.v1.IdentityStatus[rental.v1.IdentityStatus.UNSUBMITTED]
    },
    onGenderChange(opt:any){
        this.setData({
            genderIndex: parseInt(opt.detail.value)
        })
    },
    onBirthday(e:any){
        this.setData({
            birthday:e.detail.value
        })
    },

    renderProfile(p:rental.v1.IProfile){
        this.renderIdentity(p.identity!)
        this.setData({
            // licNo:p.identity?.licNumber||"",
            // name:p.identity?.name||"",
            // genderIndex:p.identity?.gender||0,
            // birthDate:formatDate(p.identity?.birthDateMillis || 0),
            state: rental.v1.IdentityStatus[p.identityStatus||0]
        })
    },
    renderIdentity(i:rental.v1.IIdentity){
        this.setData({
            licNo:i?.licNumber||"",
            name:i?.name||"",
            genderIndex:i?.gender||0,
            birthDate:formatDate(i?.birthDateMillis || 0),
        })
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(opt:Record<'redirect',string>) {
        const o: routing.RegisterOpts = opt
        if (o.redirect){
            this.redirectURL = decodeURIComponent(o.redirect)
        }
        ProfileService.getProfile().then(p =>{
            this.renderProfile(p)
        })
        ProfileService.getProfilePhoto().then(p => {
            this.setData({
                licImgURL: p.url || ""
            })
        })
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
        this.clearProfileRefreshere()
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
    onUploadLic(){
        wx.chooseImage({
            success:async res => {
                if (res.tempFilePaths.length === 0) {
                    return
                }

                this.setData({
                    licImgURL:res.tempFilePaths[0]
                })

                const uploadOpt = {
                    filePath: res.tempFilePaths[0],
                    url:"/v1/profile/photo/upload",
                    name:"testing_upload",
                    formData:{
                        fileName:"jsa"
                    },
                    respMarshaller:rental.v1.UploadFilePhotoResponse.fromObject
                }
                //上传图片
                const uploadFile = await Coolcar.wxUploadFile<rental.v1.IUploadFilePhotoRequest, rental.v1.IUploadFilePhotoResponse>(uploadOpt,{
                    attachAuthHeader: false,
                    retryOnAuthError: false,
                })

                // const uploadFile = await ProfileService.uploadProfilePhoto()
                console.log(uploadFile.url)
                return

                // const photoRes = await ProfileService.createProfilePhoto()

                // if (!photoRes.uploadUrl) {
                //     return 
                // }

                // await Coolcar.uploadfile({
                //     localPath: res.tempFilePaths[0],
                //     url: photoRes.uploadUrl
                // })

                // const identity = await ProfileService.completeProfilePhoto()

                // this.renderIdentity(identity)
                // if(res.tempFilePaths[0].length > 0){    
                //     this.setData({
                //         licImgURL:res.tempFilePaths[0]
                //     })
                //     // wx.uploadFile({
                //     //     filePath: res.tempFilePaths[0],
                //     //     name: 'abc.jpg',
                //     //     url:'https://sunjinyan-testing.oss-cn-beijing.aliyuncs.com/my-object?Expires=1654686088&OSSAccessKeyId=TMP.3KdUPwwmRStfpuSKdiw6hHgEndUDzcqwoiN5jXBxjqUsHVWkLiNu1JUsCcysP7pB8HdBsK5kBpUUQ7oZbpobm2Hi8QKSDN&Signature=aOh3FdpMfu8eyUICyn1g6QpK3uo%3D',
                //     //     success:console.log,
                //     //     fail:console.error,
                //     // })

                //     const data = wx.getFileSystemManager().readFileSync(res.tempFilePaths[0])

                //     wx.request({
                //         method:"PUT",
                //         data,
                //         url:'https://sunjinyan-testing.oss-cn-beijing.aliyuncs.com',
                //         success:console.log,
                //         fail:console.error,
                //     })
                // }
                setTimeout(()=>{
                    this.setData({
                        licNo:'12312312312',
                        name:'法外狂徒张三',
                        genderIndex:1,
                        birthday:'1990-05-23'
                    })
                })
            }
        })
    },
    onSubmit(){
        // this.setData({
        //     state:'PENDING'
        // })
        // setTimeout(() => {
        //     this.onLicVerified()
        // }, 3000);
        ProfileService.submitProfile({
            licNumber:this.data.licNo,
            name:this.data.name,
            gender:this.data.genderIndex,
            birthDateMillis:Date.parse(this.data.birthday)
        }).then(p=>{
            this.renderProfile(p)
            this.scheduleProfileRefresher()
        })
    },
    scheduleProfileRefresher(){
        this.profileRefresher = setInterval(()=>{
            ProfileService.getProfile().then(p => {
                this.renderProfile(p)
                console.log('=========================',p)
                if (p.identityStatus !== rental.v1.IdentityStatus.PENDING) {
                    this.clearProfileRefreshere()
                }
                if (p.identityStatus === rental.v1.IdentityStatus.VERIFIED) {
                    this.onLicVerified()
                }
            })
        },1000)
    },
    clearProfileRefreshere(){
        if (this.profileRefresher) {
            clearInterval(this.profileRefresher)
            this.profileRefresher = 0;
        }
    },
    onResubmit(){
        // this.setData({
        //     state:'UNSUBMITTED',
        //     licImgURL:''
        // })
        ProfileService.clearProfile().then(p => this.renderProfile(p))
        
    },
    onLicVerified(){
        // this.setData({
        //     state:'VERIFIED'
        // })

        if(this.redirectURL){
            wx.redirectTo({
                url:this.redirectURL
            })
        }
    }
})