/** Unified WeChat mini-program feedback helpers. Prefer these over ad-hoc toast/modal. */

type ToastIcon = 'success' | 'error' | 'loading' | 'none'

function toast(title: string, icon: ToastIcon = 'none', duration = 2000) {
  const text = (title || '').trim() || '操作失败'
  uni.showToast({
    title: text.length > 20 ? text.slice(0, 20) : text,
    icon,
    duration,
    mask: icon === 'loading',
  })
}

export const tip = {
  success(title = '操作成功', duration = 2000) {
    toast(title, 'success', duration)
  },
  fail(title = '操作失败') {
    toast(title, 'none')
  },
  loading(title = '加载中') {
    toast(title, 'loading', 60000)
  },
  hide() {
    uni.hideToast()
  },
  /**
   * End a loading toast, show success, optionally navigate after delay.
   * Avoid calling hide() in the same tick as success (it cancels the toast).
   */
  successThen(title: string, navigate?: () => void, delayMs = 1200) {
    try {
      uni.hideToast()
    } catch {
      /* ignore */
    }
    setTimeout(() => {
      toast(title || '操作成功', 'success', Math.max(delayMs, 1500))
      if (navigate) {
        setTimeout(navigate, delayMs)
      }
    }, 40)
  },
  /** Confirm dialog; resolves true if user confirms. */
  confirm(content: string, title = '提示'): Promise<boolean> {
    return new Promise((resolve) => {
      uni.showModal({
        title,
        content,
        confirmText: '确定',
        cancelText: '取消',
        success: (res) => resolve(!!res.confirm),
        fail: () => resolve(false),
      })
    })
  },
  /** Alert-only dialog (no cancel). */
  alert(content: string, title = '提示'): Promise<void> {
    return new Promise((resolve) => {
      uni.showModal({
        title,
        content,
        showCancel: false,
        confirmText: '确定',
        success: () => resolve(),
        fail: () => resolve(),
      })
    })
  },
}

export default tip
