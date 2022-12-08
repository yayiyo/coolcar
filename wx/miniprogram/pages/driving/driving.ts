// 每秒一分钱
const centPerSec = 0.8

function formatTime(second: number) {
  const parseString = (n: number) => {
    return n < 10 ? '0' + n.toFixed(0) : n.toFixed(0)
  }
  const h = Math.floor(second / 3600)
  second -= h * 3600
  const m = Math.floor(second / 60)
  second -= m * 60
  const s = Math.floor(second)
  return `${parseString(h)}:${parseString(m)}:${parseString(s)}`
}

function formatFee(cent: number) {
  return (cent / 100).toFixed(2)
}

Page({
  timer: undefined as number | undefined,
  data: {
    location: {
      latitude: 40.22077,
      longitude: 116.23128,
    },
    scale: 12,
    elapsed: '00:00:00',
    fee: '0.00',
  },

  onLoad(opt) {
    console.log('current trip', opt.trip_id)
    this.setupLocationUpdator()
    this.setupTimer()
  },

  onUnload() {
    wx.stopLocationUpdate()
    if (this.timer) {
      clearInterval(this.timer)
    }
  },

  setupLocationUpdator() {
    wx.startLocationUpdate({
      fail: console.error,
    })
    wx.onLocationChange(loc => {
      console.log(loc)
      this.setData({
        location: {
          latitude: loc.latitude,
          longitude: loc.longitude,
        }
      })
    })
  },
  setupTimer() {
    let elapsedSec = 0
    let cents = 0
    this.timer = setInterval(() => {
      cents += centPerSec
      elapsedSec++
      this.setData({
        elapsed: formatTime(elapsedSec),
        fee: formatFee(cents),
      })
    }, 1000)
  }

  // TODO: 结束行程
})