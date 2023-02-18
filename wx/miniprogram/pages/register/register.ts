import {routing} from "../../utils/routing"
import {ProfileService} from "../../service/profile";
import {Identity, IdentityStatus, Profile} from "../../service/proto_gen/rental/rental";
import {parseString} from "../../utils/format";
import {CoolCar} from "../../service/request";


function formatDate(dt: Date) {
    const y = dt.getFullYear()
    const m = dt.getMonth() + 1
    const d = dt.getDate()
    return `${parseString(y)}-${parseString(m)}-${parseString(d)}`
}

// pages/register/register.ts
Page({
    /**
     * 页面的初始数据
     */
    profileRefresher: 0,
    redirectURL: '',
    data: {
        licNo: '',
        name: '',
        genderIndex: 0,
        genders: ['未知', '男', '女'],
        birthDate: '2006-01-02',
        licImgURL: '',
        state: IdentityStatus[IdentityStatus.NOT_SUBMITTED],
    },

    renderProfile(p: Profile) {
        this.renderIdentity(p.identity!)
        this.setData({
            state: IdentityStatus[p.identityStatus || 0],
        })
    },

    renderIdentity(i: Identity) {
        this.setData({
            licNo: i?.licNumber || '',
            name: i?.name || '',
            genderIndex: i?.gender || 0,
            birthDate: i?.birthDate || formatDate(new Date()),
        })
    },

    onLoad(opt: Record<'redirect', string>) {
        const o: routing.RegisterOpts = opt
        if (o.redirect) {
            this.redirectURL = decodeURIComponent(o.redirect)
        }
        ProfileService.getProfile().then(this.renderProfile)
        ProfileService.getProfilePhoto().then(p => {
            this.setData({
                licImgURL: p.url || '',
            })
        })
    },

    onUnload() {
        this.clearProfileRefresher()
    },

    onUploadLic() {
        wx.chooseMedia({
            mediaType: ['image'],
            success: async res => {
                console.log(res)
                if (res.tempFiles.length === 0) {
                    return
                }
                this.setData({
                    licImgURL: res.tempFiles[0].tempFilePath
                })

                const photoRes = await ProfileService.createProfilePhoto()
                if (!photoRes.uploadUrl) {
                    return
                }
                await CoolCar.uploadFile({
                    localPath: res.tempFiles[0].tempFilePath,
                    url: photoRes.uploadUrl,
                })

                const identity = await ProfileService.completeProfilePhoto()
                this.renderIdentity(identity)
            }
        })
    },
    onGenderChange(e: any) {
        this.setData({
            genderIndex: Number(e.detail.value)
        })
    },

    onBirthDateChange(e: any) {
        this.setData({
            birthDate: e.detail.value
        })
    },

    onSubmit() {
        ProfileService.submitProfile({
            licNumber: this.data.licNo,
            name: this.data.name,
            gender: this.data.genderIndex,
            birthDate: this.data.birthDate
        }).then(p => {
            this.renderProfile(p)
            this.scheduleProfileRefresher()
        })
    },

    scheduleProfileRefresher() {
        this.profileRefresher = setInterval(() => {
            ProfileService.getProfile().then(p => {
                this.renderProfile(p)
                if (p.identityStatus !== IdentityStatus.PENDING) {
                    this.clearProfileRefresher()
                }

                if (p.identityStatus === IdentityStatus.VERIFIED) {
                    this.onLicVerify()
                }
            })
        }, 1000)
    },

    clearProfileRefresher() {
        if (this.profileRefresher) {
            clearInterval(this.profileRefresher)
            this.profileRefresher = 0
        }
    },

    onResubmit() {
        ProfileService.clearProfile().then(this.renderProfile)
        ProfileService.clearProfilePhoto().then(() => {
            this.setData({
                licImgURL: '',
            })
        })
    },

    onLicVerify() {
        if (this.redirectURL) {
            wx.redirectTo({
                url: this.redirectURL,
            })
        }
    },
})