import camelcaseKeys from "camelcase-keys"
import { auth } from "../service/proto_gen/auth/auth_pb"

export  namespace Coolcar {
    const  url: string = "http://coolcar.dreaminglifes.com"
    let   token: string = ""
    let   expire: number = 0
    //统一处理请求
    // wx.request({
    //     url: 'http://localhost:8080/v1/auth/login',
    //     method: 'POST',
    //     data: {
    //       code: res.code
    //     } as auth.v1.ILoginRequest,
    //     success: res => {
    //        const loginResp: auth.v1.LoginResponse = 
    //        auth.v1.LoginResponse.fromObject(camelcaseKeys(res.data as object,{deep:true}))
    //        console.log(loginResp)
    //        wx.request({
    //         url: 'http://localhost:8080/v1/trip',
    //         method: 'POST',
    //         header: {
    //           authorization: 'Bearer ' + loginResp.accessToken
    //         },
    //         data: {
    //           start: "123789"
    //         } as rental.v1.ICreateTripRequest,
    //         success: res => {
    //           const tripResp: rental.v1.CreateTripResponse = 
    //           rental.v1.CreateTripResponse.fromObject(camelcaseKeys(res.data as object,{deep:true}))
    //           console.log(tripResp)
    //         }
    //        })
    //     },
    //     fail: console.error
    //   })
    //首先定义一个request的统一参数

    interface RequestOptions <REQ,RES> {
        path:string
        method: "GET" | "POST" | "DELETE" | "PUT" | "OPTIONS" | "HEAD" | "TRACE" | "CONNECT"
        data?: REQ,
        anyfromObjectFunc:(r: object) => RES
    }


    const AUTH_ERROR = "AUTH_ERROR"

    interface headerOpt {
        needAttachTokenHeader: boolean
    }


    interface retryOpts {
        needRetry: boolean
        times: number
    }


    function wxLogin(): Promise<WechatMiniprogram.LoginSuccessCallbackResult> {
        return  new Promise((resolve, reject) => {
            wx.login({
                success:resolve,
                fail:reject
            })
        })
    }



    export async function  login(){
        if (token != "" && expire >= Date.now()){
            return
        }
        const wl = await wxLogin()
        sendRequest<auth.v1.ILoginRequest,auth.v1.ILoginResponse>({
            path: "/v1/auth/login",
            method: "POST",
            data: {
                code: wl.code
            } as auth.v1.ILoginRequest,
            anyfromObjectFunc: auth.v1.LoginResponse.fromObject
        },{
            needAttachTokenHeader: false
        }).then(res => {
            token = res.accessToken as string
            expire = res.expiresIn as number
        })
    }

    export async  function  sendRequestWithRetry<REQ,RES>(o:RequestOptions<REQ,RES>,ho: headerOpt,ro: retryOpts): Promise<RES>{
        const r = ro || {
            needRetry:true,
            times:0
        }
        await login()
        try {
            return sendRequest(o,ho)
        } catch (error) {
            if (error === AUTH_ERROR && ro.needRetry && ro.times <= 0) {
                token = ""
                expire = 0
                r.needRetry = false
                r.times+=1//只重试一次
                return sendRequestWithRetry(o,ho,r)
            }else{
                throw error
            }
        }
    }

    function sendRequest<REQ,RES>(o:RequestOptions<REQ,RES>,ho: headerOpt): Promise<RES> {

        const h = ho || {
            needAttachTokenHeader: true
        }

        return new Promise<RES>((resolve, reject) => {
            const header: Record<string,any> = {}
            if (ho.needAttachTokenHeader) {//非默认情况
                if (token !== "" && expire >= Date.now()) {
                    header.authorization = "Bearer " + token
                }else{
                    reject(AUTH_ERROR)
                    return
                }
            }
            if (h.needAttachTokenHeader && token !== "" && expire >= Date.now())  {
                header.authorization = "Bearer " + token //默认情况
            }

            wx.request({
                url: url + o.path,
                method: o.method,
                data: o.data,
                header,
                success: res => {
                    if (res.statusCode === 401) {
                        reject(AUTH_ERROR)
                    }else if (res.statusCode >= 400){
                        reject("Unknow")
                    }else{
                     resolve(o.anyfromObjectFunc(
                        camelcaseKeys(res.data as object, {
                            deep: true,
                        })))
                    }
                },
                fail: reject
            })

        })
    }


}