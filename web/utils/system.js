/** Status bar height without deprecated getSystemInfoSync when possible. */
export function getStatusBarHeight(fallback = 20) {
  try {
    if (typeof uni.getWindowInfo === 'function') {
      const win = uni.getWindowInfo()
      if (win && win.statusBarHeight != null) {
        return Number(win.statusBarHeight) || fallback
      }
    }
  } catch (e) {
    /* fall through */
  }
  try {
    // Legacy fallback for older runtimes / non-WeChat hosts.
    const sys = uni.getSystemInfoSync()
    return (sys && sys.statusBarHeight) || fallback
  } catch (e) {
    return fallback
  }
}
