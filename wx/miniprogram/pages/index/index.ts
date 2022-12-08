// index.ts
// 获取应用实例

const app = getApp<IAppOption>()

Page({
    isPageShowing: false,
    avatarURL: '',
    data: {
        setting: {
            skew: 0,
            rotate: 0,
            showLocation: true,
            showScale: true,
            subKey: '',
            layerStyle: 1,
            enableZoom: true,
            enableScroll: true,
            enableRotate: false,
            showCompass: false,
            enable3D: false,
            enableOverlooking: false,
            enableSatellite: false,
            enableTraffic: false,
        },
        location: {
            latitude: 23.0009343,
            longitude: 113.36299,
        },
        scale: 10,
        markers: [
            {
                iconPath: "/resources/car.png",
                id: 0,
                latitude: 23.0009343,
                longitude: 113.36299,
                width: 20,
                height: 20
            },
            {
                iconPath: "/resources/car.png",
                id: 1,
                latitude: 23.6329420,
                longitude: 114.692635,
                width: 20,
                height: 20
            }
        ]
    },
    onLoad() {
        const avatarURL = wx.getStorageSync('avatar')
        this.setData({
            avatarURL,
        })
    },
    onMyLocationTap() {
        wx.getLocation({
            success: res => {
                this.setData({
                    location: {
                        latitude: res.latitude,
                        longitude: res.longitude,
                    }
                })
            },
            fail: () => {
                wx.showToast({
                    title: '没有位置权限',
                    icon: 'error'
                })
            }
        })
    },
    onScanTap() {
        wx.scanCode({
            success: res => {
                console.log(res)
                const car_id = '88888'
                const redirectURL = `/pages/lock/lock?car_id=${car_id}`
                wx.navigateTo({
                    url: `/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`,
                })
            },
            fail: console.error,
        })
    },
    onHide() {
        this.isPageShowing = false
    },

    onShow() {
        this.isPageShowing = true
    },
    moveCars() {
        const map = wx.createMapContext("map")
        const dest = {
            latitude: this.data.markers[0].latitude,
            longitude: this.data.markers[0].longitude,
        }

        const moveCar = () => {
            dest.latitude += 0.1
            dest.longitude += 0.1
            map.translateMarker({
                destination: {
                    latitude: dest.latitude,
                    longitude: dest.longitude,
                },
                markerId: 0,
                autoRotate: false,
                rotate: 0,
                duration: 5000,
                animationEnd: () => {
                    this.data.markers[0].latitude = dest.latitude
                    this.data.markers[0].longitude = dest.longitude
                    if (this.isPageShowing) {
                        moveCar()
                    }
                }
            })
        }
        moveCar()
    },
    onMyTripsTap() {
        wx.navigateTo({
            url: '/pages/mytrips/mytrips',
        })
    }
})
