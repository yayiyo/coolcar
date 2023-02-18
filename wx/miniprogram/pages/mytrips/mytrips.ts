import {routing} from "../../utils/routing"
import {TripService} from "../../service/trip";
import {IdentityStatus, TripEntity, TripStatus} from "../../service/proto_gen/rental/rental";
import {ProfileService} from "../../service/profile";
import {formatDuration, formatFee} from "../../utils/format";

interface Trip {
    id: string
    shortId: string
    start: string
    end: string
    duration: string
    fee: string
    distance: string
    status: string
    inProgress: boolean
}

interface MainItem {
    id: string
    navId: string
    navScrollId: string
    data: Trip
}

interface NavItem {
    id: string
    mainId: string
    label: string
}

interface MainItemQueryResult {
    id: string
    top: number
    dataset: {
        navId: string
        navScrollId: string
    }
}

const tripStatusMap = new Map([
    [TripStatus.IN_PROGRESS, "进行中"],
    [TripStatus.FINISHED, "已完成"]
])

const licStatusMap = new Map([
    [IdentityStatus.NOT_SUBMITTED, "未认证"],
    [IdentityStatus.PENDING, "认证中"],
    [IdentityStatus.VERIFIED, "已认证"]
])

Page({
    scrollStates: {
        mainItems: [] as MainItemQueryResult[]
    },
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
        licStatus: licStatusMap.get(IdentityStatus.NOT_SUBMITTED || 0),
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
        Promise.all([TripService.getTrips()]).then(([trips]) => {
            this.populateTrips(trips.trips)
        })
        const avatarURL = wx.getStorageSync('avatar')
        this.setData({
            avatarURL,
        })
    },

    onShow() {
        ProfileService.getProfile().then(p => {
            this.setData({
                licStatus: licStatusMap.get(p.identityStatus || 0)
            })
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

    populateTrips(trips: TripEntity[]) {
        const mainItems: MainItem[] = []
        const navItems: NavItem[] = []
        let navSel = ''
        let prevNav = ''
        for (let i = 0; i < trips?.length || 0; i++) {
            const trip = trips[i]
            const mainId = 'main-' + i
            const navId = 'nav-' + i
            const shortId = trip.id?.substring(trip.id.length - 6)
            if (!prevNav) {
                prevNav = navId
            }
            const tripData: Trip = {
                id: trip.id!,
                shortId: '****' + shortId,
                start: trip.trip?.start?.poiName || '未知',
                end: '',
                distance: '',
                duration: '',
                fee: '',
                status: tripStatusMap.get(trip.trip?.status!) || '未知',
                inProgress: trip.trip?.status === TripStatus.IN_PROGRESS,
            }
            const end = trip.trip?.end
            if (end) {
                tripData.end = end.poiName || '未知',
                    tripData.distance = end.kmDriven?.toFixed(1) + '公里',
                    tripData.fee = formatFee(end.feeCent || 0)
                const dur = formatDuration(Number((end.timestampSec) || 0) - (Number(trip.trip?.start?.timestampSec) || 0))
                tripData.duration = `${dur.hh}时${dur.mm}分`
            }
            mainItems.push({
                id: mainId,
                navId: navId,
                navScrollId: prevNav,
                data: tripData,
            })
            navItems.push({
                id: navId,
                mainId: mainId,
                label: shortId || '',
            })
            if (i === 0) {
                navSel = navId
            }
            prevNav = navId
        }
        for (let i = 0; i < this.data.navCount - 1; i++) {
            navItems.push({
                id: '',
                mainId: '',
                label: '',
            })
        }
        this.setData({
            mainItems,
            navItems,
            navSel,
        }, () => {
            this.prepareScrollStates()
        })
    },

    prepareScrollStates() {
        wx.createSelectorQuery().selectAll('.main-item')
            .fields({
                id: true,
                dataset: true,
                rect: true,
            }).exec(res => {
            this.scrollStates.mainItems = res[0]
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
        const {avatarUrl} = e.detail
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

    onMainScroll(e: any) {
        const top: number = e.currentTarget?.offsetTop + e.detail?.scrollTop
        if (top === undefined) {
            return
        }

        const selItem = this.scrollStates.mainItems.find(
            v => v.top >= top)
        if (!selItem) {
            return
        }

        this.setData({
            navSel: selItem.dataset.navId,
            navScr: selItem.dataset.navScrollId,
        })
    },
})
