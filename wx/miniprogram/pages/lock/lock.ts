Page({
  data: {
    avatarUrl: '',
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

        wx.showLoading({
          title: '开锁中',
        })

        setTimeout(function () {
          wx.redirectTo({
            url: '/pages/driving/driving',
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