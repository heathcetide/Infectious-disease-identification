/** Shared theme tokens for the CDC screen (cards + ECharts). */

export function isDarkTheme() {
  if (typeof document === 'undefined') return false
  const root = document.documentElement
  return root.classList.contains('dark') || root.getAttribute('data-theme') === 'dark'
}

export function screenSurface() {
  return isDarkTheme() ? '#1a1a1a' : '#fafafa'
}

export function screenChartColors() {
  const dark = isDarkTheme()
  return {
    text: dark ? '#cfd3dc' : '#666666',
    muted: dark ? '#a3a6ad' : '#888888',
    axis: dark ? '#8d9095' : '#999999',
    tooltipBg: dark ? '#1d1e1f' : '#ffffff',
    tooltipText: dark ? '#e5eaf3' : '#333333',
    track: dark ? '#2a2a2a' : '#f4f4f4',
    border: dark ? '#414243' : '#e4e7ed',
  }
}
