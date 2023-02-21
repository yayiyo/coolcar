// index.ts
// 获取应用实例

import {IAppOption} from "../../appoption"
import {routing} from "../../utils/routing"
import {TripService} from "../../service/trip";
import {IdentityStatus, TripStatus} from "../../service/proto_gen/rental/rental";
import {ProfileService} from "../../service/profile";
import {CarService} from "../../service/car";

const app = getApp<IAppOption>()

interface Marker {
    iconPath: string
    id: number
    latitude: number
    longitude: number
    width: number
    height: number
}

const defaultAvatar = '/resources/car.png'
const initialLat = 30
const initialLng = 120

Page({
    isPageShowing: false,
    socket: undefined as WechatMiniprogram.SocketTask | undefined,
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
            latitude: initialLat,
            longitude: initialLng,
        },
        scale: 10,
        markers: [] as Marker[],
    },
    onLoad() {

        this.socket = CarService.subscribe(msg => {
            console.log(msg)
        })

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
    async onScanTap() {
        const trips = await TripService.getTrips(TripStatus.IN_PROGRESS)
        if (trips.trips?.length || 0 > 0) {
            await this.selectComponent('#tripModal').showModal().then((res: 'ok' | 'cancel' | 'close') => {
                console.log(res)
                if (res === 'ok') {
                    wx.navigateTo({
                        url: routing.driving({
                            trip_id: trips.trips[0].id
                        })
                    })
                }
            })
            return
        }

        wx.scanCode({
            success: async () => {
                const car_id = '63f1e13c288896196b07f276'
                const lockUrl = routing.lock({car_id})
                const profile = await ProfileService.getProfile()
                if (profile.identityStatus === IdentityStatus.VERIFIED) {
                    wx.navigateTo({
                        url: lockUrl,
                    })
                } else {
                    await this.selectComponent('#licModal').showModal().then((res: 'ok' | 'cancel' | 'close') => {
                        console.log(res)
                        if (res === 'ok') {
                            wx.navigateTo({
                                url: routing.register({
                                    redirectURL: lockUrl,
                                }),
                            })
                        }
                    })

                }
            },
            fail: console.error,
        })
    },
    onHide() {
        this.isPageShowing = false
        if (this.socket) {
            this.socket.close({
                success: () => {
                    this.socket = undefined
                }
            })
        }
    },

    onShow() {
        this.isPageShowing = true
        const avatarURL = wx.getStorageSync('avatar')
        this.setData({
            avatarURL,
        })

        if (!this.socket) {
            this.setData({
                markers: []
            }, () => {
                this.setupCarPositionUpdater()
            })
        }
    },
    setupCarPositionUpdater() {
        const map = wx.createMapContext("map")
        const markersByCarID = new Map<string, Marker>()
        const endTrans = () => {
            translationInProgress = false
        }
        let translationInProgress = false
        this.socket = CarService.subscribe(car => {
            if (!car.id || translationInProgress || !this.isPageShowing) {
                return
            }
            const marker = markersByCarID.get(car.id)
            if (!marker) {
                // new marker
                const newMarker: Marker = {
                    id: this.data.markers.length,
                    iconPath: car.car?.driver?.avatarUrl || defaultAvatar,
                    latitude: car.car?.position?.latitude || initialLat,
                    longitude: car.car?.position?.longitude || initialLng,
                    height: 20,
                    width: 20,
                }
                markersByCarID.set(car.id, newMarker)
                this.data.markers.push(newMarker)
                translationInProgress = true
                this.setData({
                    markers: this.data.markers,
                }, endTrans)
                return
            }

            const newAvatar = car.car?.driver?.avatarUrl || defaultAvatar
            const newLat = car.car?.position?.latitude || initialLat
            const newLng = car.car?.position?.longitude || initialLng
            if (marker.iconPath !== newAvatar) {
                marker.iconPath = newAvatar
                marker.latitude = newLat
                marker.longitude = newLng
                translationInProgress = true
                this.setData({
                    markers: this.data.markers,
                }, endTrans)
            }

            if (marker.latitude !== newLat || marker.longitude !== newLng) {
                translationInProgress = true
                map.translateMarker({
                    destination: {
                        latitude: newLat,
                        longitude: newLng,
                    },
                    markerId: 0,
                    autoRotate: false,
                    rotate: 0,
                    duration: 90,
                    animationEnd: endTrans,
                })
            }
        })
    },
    onMyTripsTap() {
        wx.navigateTo({
            url: routing.mytrips(),
        })
    },
})
