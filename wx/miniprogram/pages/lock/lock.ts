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
  }
})