import { routing } from "../../utils/routing"

interface Trip {
  id: string
  start: string
  end: string
  duration: string
  fee: string
  distance: string
  status: string
}

interface MainItem {
  id: string
  navId: string
  data: Trip
}

interface NavItem {
  id: string
  mainId: string
  label: string
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
    navCount: 0,
    mainItems: [] as MainItem[],
    mainScroll: '',
    navItems: [] as NavItem[],
    navSel: '',
    navScr: '',
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
        const height = wx.getSystemInfoSync().windowHeight - res.height
        this.setData({
          tripsHeight: height,
          navCount: Math.round(height / 50),
        })
      }).exec()
  },

  populateTrips() {
    const mainItems: MainItem[] = []
    const navItems: NavItem[] = []
    let navSel = ''
    for (let i = 0; i < 100; i++) {
      const mainId = 'main-' + i
      const navId = 'nav-' + i
      const trip_id = (10001 + i).toString()
      mainItems.push({
        id: mainId,
        navId: navId,
        data: {
          id: trip_id,
          start: '北京',
          end: '上海',
          duration: '72时44分55秒',
          fee: '28888元',
          distance: '2000KM',
          status: '已结束',
        },
      })

      navItems.push({
        id: navId,
        mainId: mainId,
        label: trip_id
      })

      if (i === 0) {
        navSel = navId
      }
    }
    this.setData({
      mainItems,
      navItems,
      navSel,
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
  },

  onNavItemTap(e: any) {
    const mainId: string = e.currentTarget?.dataset?.mainId
    const navId: string = e.currentTarget?.id
    if (mainId && navId) {
      this.setData({
        mainScroll: mainId,
        navSel: navId,
      })
    }
  },

  onMainItemTap(e: any) {
    console.log(e)
    const navId = e.currentTarget?.dataset?.navId
    if (navId) {
      this.setData({
        navScr: navId,
      })
    }
  }
})
