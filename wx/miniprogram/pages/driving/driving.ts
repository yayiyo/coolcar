import {routing} from "../../utils/routing"
import {TripService} from "../../service/trip";
import {formatDuration, formatFee, parseString} from "../../utils/format";
import {Trip, TripStatus} from "../../service/proto_gen/rental/rental";


const updateIntervalSec = 5
const initialLat = 30
const initialLng = 120

function durationStr(sec: number) {
    const dur = formatDuration(sec)
    return `${dur.hh}:${dur.mm}:${dur.ss}`
}

function formatTime(second: number) {
    const h = Math.floor(second / 3600)
    second -= h * 3600
    const m = Math.floor(second / 60)
    second -= m * 60
    const s = Math.floor(second)
    return `${parseString(h)}:${parseString(m)}:${parseString(s)}`
}

Page({
    timer: undefined as number | undefined,
    tripID: '',
    data: {
        location: {
            latitude: initialLat,
            longitude: initialLng,
        },
        scale: 12,
        elapsed: '00:00:00',
        fee: '0.00',
        markers: [
            {
                iconPath: "/resources/car.png",
                id: 0,
                latitude: initialLat,
                longitude: initialLng,
                width: 20,
                height: 20,
            },
        ],
    },

    onLoad(opt: Record<'trip_id', string>) {
        const o: routing.DriveOpts = opt
        console.log('current trip', o.trip_id)
        this.tripID = o.trip_id
        this.setupLocationUpdator()
        this.setupTimer(o.trip_id)
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
    async setupTimer(tripID: string) {
        const trip = await TripService.getTrip(tripID)
        if (trip.status !== TripStatus.IN_PROGRESS) {
            console.error('trip not in progress')
            return
        }
        let secSinceLastUpdate = 0
        let lastUpdateDurationSec = trip.current!.timestampSec! - trip.start!.timestampSec!
        const toLocation = (trip: Trip) => ({
            latitude: trip.current?.location?.latitude || initialLat,
            longitude: trip.current?.location?.longitude || initialLng,
        })
        const location = toLocation(trip)
        this.data.markers[0].latitude = location.latitude
        this.data.markers[0].longitude = location.longitude
        this.setData({
            elapsed: durationStr(Number(lastUpdateDurationSec)),
            fee: formatFee(trip.current?.feeCent || 0),
            location,
            markers: this.data.markers,
        })

        this.timer = setInterval(() => {
            secSinceLastUpdate++
            if (secSinceLastUpdate % updateIntervalSec === 0) {
                TripService.getTrip(tripID).then(trip => {
                    console.log(trip)
                    lastUpdateDurationSec = trip.current!.timestampSec! - trip.start!.timestampSec!
                    secSinceLastUpdate = 0
                    const location = toLocation(trip)
                    this.data.markers[0].latitude = location.latitude
                    this.data.markers[0].longitude = location.longitude
                    this.setData({
                        fee: formatFee(trip.current?.feeCent || 0),
                        location,
                        markers: this.data.markers,
                    })
                }).catch(console.error)
            }
            this.setData({
                elapsed: durationStr(Number(lastUpdateDurationSec) + secSinceLastUpdate),
            })
        }, 1000)
    },

    onEndTripTap() {
        TripService.finishTrip(this.tripID).then(() => {
            wx.redirectTo({
                url: routing.mytrips(),
            })
        }).catch(err => {
            console.error(err)
            wx.showToast({
                title: '结束行程失败',
                icon: 'none',
            })
        })
    }
})