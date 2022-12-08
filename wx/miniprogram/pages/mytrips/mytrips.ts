import { routing } from "../../utils/routing"

Page({
  data: {
    promotionItems: [
      {
        img: 'https://img.mukewang.com/5f7301d80001fdee18720764.jpg',
        promotionID: 1,
      },
      {
        img: 'https://img.mukewang.com/5f6805710001326c18720764.jpg',
        promotionID: 2,
      },
      {
        img: 'https://img.mukewang.com/5f6173b400013d4718720764.jpg',
        promotionID: 3,
      },
      {
        img: 'https://img.mukewang.com/5f7141ad0001b36418720764.jpg',
        promotionID: 4,
      },
    ],
    avatarURL: '',
  },

  onLoad() {
    const avatarURL = wx.getStorageSync('avatar')
    this.setData({
      avatarURL,
    })
  },

  onPromotionItemTap(e: any) {
    const promotionID: number = e.currentTarget.dataset.promotionId
    if (promotionID) {
      console.log('claiming promotion', promotionID)
    }
  },
  onChooseAvatar(e: any) {
    console.log(e)
    const { avatarUrl } = e.detail
    this.setData({
      avatarURL: avatarUrl,
    })
    wx.setStorageSync('avatar', avatarUrl)
  },
  onRegisterTap() {
    wx.navigateTo({
      url: routing.register(),
    })
  }
})
