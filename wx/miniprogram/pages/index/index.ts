// index.ts
// 获取应用实例

const app = getApp<IAppOption>()

Page({
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
            latitude: 31,
            longitude: 120,
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
            },
            {
                iconPath: "/resources/car.png",
                id: 2,
                latitude: 29.47294735,
                longitude: 113.6329432,
                width: 20,
                height: 20
            }
        ]
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
    }
})
