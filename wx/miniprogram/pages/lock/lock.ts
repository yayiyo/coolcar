import {routing} from "../../utils/routing"
import {TripService} from "../../service/trip";

Page({
    carId: '',
    data: {
        avatarUrl: '',
    },

    onLoad(opt: Record<'car_id', string>) {
        const o: routing.LockOpts = opt
        this.carId = o.car_id
        console.log(o)
        console.log('unlock car', o.car_id)
        const avatarUrl = wx.getStorageSync('avatar')
        this.setData({
            avatarUrl,
        })
    },
    onChooseAvatar(e: any) {
        console.log(e)
        const {avatarUrl} = e.detail
        this.setData({
            avatarUrl,
        })
        wx.setStorageSync('avatar', avatarUrl)
    },
    onUnlockTap() {
        wx.getLocation({
            success: async loc => {
                console.log('starting a trip', {
                    location: {
                        latitude: loc.latitude,
                        longitude: loc.longitude,
                    },
                    avatarURL: this.data.avatarUrl,
                })

                if (!this.carId) {
                    console.error('car_id is required')
                    return
                }

                const trip = await TripService.createTrip({
                    start: {
                        latitude: loc.latitude,
                        longitude: loc.longitude,
                    },
                    carId: this.carId,
                })

                if (!trip.id) {
                    console.error('trip_id is required')
                    return
                }

                wx.showLoading({
                    title: '开锁中',
                })

                setTimeout(function () {
                    wx.redirectTo({
                        url: routing.driving({
                            trip_id: trip.id,
                        }),
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