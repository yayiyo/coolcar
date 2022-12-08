Page({
  data: {
    avatarUrl: '',
  },

  onLoad(opt) {
    console.log(opt)
    console.log('unloack car', opt.car_id)
  },
  onChooseAvatar(e: any) {
    console.log(e)
    const { avatarUrl } = e.detail
    this.setData({
      avatarUrl,
    })
    wx.setStorageSync('avatar', avatarUrl)
  },
  onUnlockTap() {
    wx.getLocation({
      success: loc => {
        console.log('startting a trip', {
          location: {
            latitude: loc.latitude,
            longitude: loc.longitude,
          },
          avatarURL: this.data.avatarUrl,
        })

        const tripID = '2324'

        wx.showLoading({
          title: '开锁中',
        })

        setTimeout(function () {
          wx.redirectTo({
            url: `/pages/driving/driving?trip_id=${tripID}`,
          })
        }, 3000)
      },
      fail: () => {
        wx.showToast({
          title: '没有位置权限',
          icon: 'error'
        })
      }
    })
  }
})