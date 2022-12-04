// pages/register/register.ts
Page({
  /**
   * 页面的初始数据
   */
  data: {
    licNo: '',
    name: '',
    genderIndex: 0,
    genders: ['未知', '男', '女', '其他'],
    birthDate: '1993-06-01',
    licImgURL: '',
    state: 'UNSUBMIITED' as 'UNSUBMIITED' | 'PEDING' | 'VERIFIED',
  },

  onUploadLic() {
    wx.chooseMedia({
      mediaType: ['image'],
      success: res => {
        console.log(res)
        if (res.tempFiles.length > 0) {
          this.setData({
            licImgURL: res.tempFiles[0].tempFilePath
          })
          // TODO: check the lisence and set the info
          setTimeout(() => {
            this.setData({
              licNo: '29852539042895',
              name: '李大锤',
              genderIndex: 1,
              birthDate: '2008-08-08'
            })
          }, 1000);
        }
      }
    })
  },
  onGenderChange(e: any) {
    this.setData({
      genderIndex: e.detail.value
    })
  },

  onBirthDateChange(e: any) {
    this.setData({
      birthDate: e.detail.value
    })
  },

  onSubmit() {
    this.setData({
      state: 'PEDING',
    })

    setTimeout(() => {
      this.onLicVerify()
    }, 3000);
  },

  onResubmit() {
    this.setData({
      state: 'UNSUBMIITED',
      licImgURL: '',
    })
  },

  onLicVerify() {
    this.setData({
      state: 'VERIFIED',
    })
    wx.redirectTo({
      url: '/pages/lock/lock',
    })
  },
})