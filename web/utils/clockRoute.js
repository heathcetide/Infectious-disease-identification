/** Shared redirect after daily check-in status is known. */

function todayStr() {
  const now = new Date()
  const y = now.getFullYear()
  const m = String(now.getMonth() + 1).padStart(2, '0')
  const d = String(now.getDate()).padStart(2, '0')
  return `${y}-${m}-${d}`
}

function signedResultUrl(data) {
  const status = data && (data.dailyhealthstatus || data.dailyHealthStatus)
  const healthy = status == null ? true : status.isHealth !== false && status.health !== false
  return healthy ? '/pages/clockSuccess/clockSuccess' : '/pages/clockFail/clockFail'
}

/**
 * If already signed today, redirect to success/fail page.
 * Returns true when redirected (caller should stop).
 */
export function redirectIfSignedToday(data, { replace = true } = {}) {
  if (!data || !data.hasSignedToday) return false
  const url = signedResultUrl(data)
  if (replace) {
    uni.redirectTo({ url })
  } else {
    uni.navigateTo({ url })
  }
  return true
}

export { todayStr, signedResultUrl }
