import {routing} from "../../utils/routing"
import {TripService} from "../../service/trip";
import {CarService} from "../../service/car";
import {CarStatus} from "../../service/proto_gen/car/car";

Page({
    carId: '',
    carRefresher: 0,
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
                    avatarUrl: this.data.avatarUrl,
                })

                if (!trip.id) {
                    console.error('trip_id is required')
                    return
                }

                wx.showLoading({
                    title: '开锁中',
                })

                this.carRefresher = setInterval(async () => {
                    const car = await CarService.getCar(this.carId)
                    if (car.status === CarStatus.UNLOCKED) {
                        this.clearRefresher()
                        wx.redirectTo({
                            url: routing.driving({
                                trip_id: trip.id,
                            }),
                        })
                    }
                }, 2000)
            },
            fail: () => {
                wx.showToast({
                    title: '没有位置权限',
                    icon: 'error'
                })
            }
        })
    },

    onUnload() {
        this.clearRefresher()
        wx.hideLoading()
    },
    clearRefresher() {
        if (this.carRefresher) {
            clearInterval(this.carRefresher)
            this.carRefresher = 0
        }
    },
})