import { ModalResult } from "./types"

// components/modal/modal.ts
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    showModal: Boolean,
    showCancel: Boolean,
    title: String,
    content: String,
  },

  options: {
    addGlobalClass: true,
  },

  /**
   * 组件的初始数据
   */
  data: {
    resolve: undefined as ((r: ModalResult) => void) | undefined,
  },

  /**
   * 组件的方法列表
   */
  methods: {
    onCancel(e: any) {
      console.log(e)
      this.hideModal('cancel')
    },

    onOK(e: any) {
      console.log(e)
      this.hideModal('ok')
    },

    hideModal(res: ModalResult) {
      console.log(res)
      this.setData({
        showModal: false,
      })
      this.triggerEvent(res)
      if (this.data.resolve) {
        this.data.resolve(res)
      }
    },

    showModal(): Promise<ModalResult> {
      this.setData({
        showModal: true,
      })
      return new Promise((resolve) => {
        this.data.resolve = resolve
      })
    }
  }
})
