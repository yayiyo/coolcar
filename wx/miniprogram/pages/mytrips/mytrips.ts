import { routing } from "../../utils/routing"

interface Trip {
  id: string
  start: string
  end: string
  duration: string
  fee: string
  distance: string
}

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
    tripsHeight: 0,
    trips: [] as Trip[],
  },

  onLoad() {
    this.populateTrips()
    const avatarURL = wx.getStorageSync('avatar')
    this.setData({
      avatarURL,
    })
  },

  onReady() {
    wx.createSelectorQuery().select('#heading')
      .boundingClientRect(res => {
        this.setData({
          tripsHeight: wx.getSystemInfoSync().windowHeight - res.height,
        })
      }).exec()
  },

  populateTrips() {
    const trips: Trip[] = []
    for (let i = 0; i < 100; i++) {
      trips.push({
        id: (10001 + i).toString(),
        start: '北京',
        end: '上海',
        duration: '72时44分55秒',
        fee: '28888元',
        distance: '2000KM',
      })
    }
    this.setData({
      trips,
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
