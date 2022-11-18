// app.ts

import {getSetting, getUserProfile} from "./utils/util";

let resolveUserInfo: (value: (WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>)) => void
let rejectUserInfo: (reason?: any) => void

App<IAppOption>({
    globalData: {
        userInfo: new Promise((resolve, reject) => {
            resolveUserInfo = resolve
            rejectUserInfo = reject
        })
    },
    onLaunch() {
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
            },
        })

        // 获取用户信息
        getSetting().then(res => {
            if (res.authSetting['scope.userInfo']) {
                return getUserProfile()
            }
            return undefined
        }).then(res => {
            if (!res) {
                return
            }
            resolveUserInfo(res.userInfo)
            // }).catch(err => reject(err))
        }).catch(rejectUserInfo)
    },
    resolveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
        resolveUserInfo(userInfo)
    }
})