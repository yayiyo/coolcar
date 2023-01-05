// app.ts

import {getSetting, getUserProfile} from "./utils/wxapi";
import {IAppOption} from "./appoption"
import {CoolCar} from "./service/request";

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
        CoolCar.login()

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