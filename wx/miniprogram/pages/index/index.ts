// index.ts
// 获取应用实例

const app = getApp<IAppOption>()

Page({
    data: {
        motto: 'Hello World',
        userInfo: {},
        hasUserInfo: false,
    },
    // 事件处理函数
    bindViewTap() {
        wx.redirectTo({
            url: '../logs/logs',
        })
    },
    onLoad() {
        app.globalData.userInfo.then(userInfo => {
            this.setData({
                userInfo: userInfo,
                hasUserInfo: true
            })
        })
    },

    getUserProfile() {
        wx.getUserProfile({
            desc: '展示用户信息',
            success: res => {
                app.resolveUserInfo(res.userInfo)
            }
        })
    }
})
