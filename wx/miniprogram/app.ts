// app.ts

import {getSetting, getUserProfile} from "./utils/wxapi";
import {IAppOption} from "./appoption"
import {LoginRequest, LoginResponse} from "./service/proto_gen/auth/auth";

let resolveUserInfo: (value: (WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>)) => void
let rejectUserInfo: (reason?: any) => void

App<IAppOption>({
    globalData: {
        userInfo: new Promise((resolve, reject) => {
            resolveUserInfo = resolve
            rejectUserInfo = reject
        }),
    },
    async onLaunch() {
        // 展示本地存储能力
        const logs = wx.getStorageSync('logs') || []
        logs.unshift(Date.now())
        wx.setStorageSync('logs', logs)

        // 登录
        wx.login({
            success: res => {
                console.log('..................')
                console.log(res)
                console.log(res.code)
                // 发送 res.code 到后台换取 openId, sessionKey, unionId
                wx.request({
                    url: 'http://localhost:8080/v1/auth/login',
                    method: 'POST',
                    data: {
                        code: res.code,
                    } as LoginRequest,
                    success:result => {
                        console.log(result)
                        const res = LoginResponse.fromJson(result.data as string)
                        console.log(res)
                    },
                    fail:console.error,
                })
            },
        })

        try {
            const setting = await getSetting()
            if (setting.authSetting['scope.userInfo']) {
                const userProfile = await getUserProfile()
                resolveUserInfo(userProfile.userInfo)
            }
        } catch (err) {
            rejectUserInfo(err)
        }
    },
    resolveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
        resolveUserInfo(userInfo)
    },
})