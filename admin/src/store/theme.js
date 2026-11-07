import { defineStore } from 'pinia';

const STORAGE_KEY = 'admin-theme-mode';
const PRIMARY = '#285AC8';

function hexToRgb(hex) {
  const h = hex.replace('#', '');
  const full = h.length === 3 ? h.split('').map((c) => c + c).join('') : h;
  const n = parseInt(full, 16);
  return { r: (n >> 16) & 255, g: (n >> 8) & 255, b: n & 255 };
}

function mix(hex, ratio, toward = '#ffffff') {
  const a = hexToRgb(hex);
  const b = hexToRgb(toward);
  const r = Math.round(a.r + (b.r - a.r) * ratio);
  const g = Math.round(a.g + (b.g - a.g) * ratio);
  const bl = Math.round(a.b + (b.b - a.b) * ratio);
  return `#${[r, g, bl].map((x) => x.toString(16).padStart(2, '0')).join('')}`;
}

export function applyThemeMode(mode) {
  if (typeof document === 'undefined') return;
  const root = document.documentElement;
  const isDark = mode === 'dark';
  root.setAttribute('data-theme', isDark ? 'dark' : 'light');
  root.classList.toggle('dark', isDark);

  // keep brand primary stable; only surface/text flip
  root.style.setProperty('--el-color-primary', PRIMARY);
  root.style.setProperty('--el-color-primary-light-3', mix(PRIMARY, 0.3));
  root.style.setProperty('--el-color-primary-light-5', mix(PRIMARY, 0.5));
  root.style.setProperty('--el-color-primary-light-7', mix(PRIMARY, 0.7));
  root.style.setProperty('--el-color-primary-light-8', mix(PRIMARY, 0.8));
  root.style.setProperty('--el-color-primary-light-9', mix(PRIMARY, 0.9));
  root.style.setProperty('--el-color-primary-dark-2', mix(PRIMARY, 0.2, '#000000'));
  root.style.setProperty('--admin-brand', PRIMARY);

  if (isDark) {
    root.style.setProperty('--el-bg-color', '#141414');
    root.style.setProperty('--el-bg-color-page', '#0a0a0a');
    root.style.setProperty('--el-bg-color-overlay', '#1d1e1f');
    root.style.setProperty('--el-text-color-primary', '#e5eaf3');
    root.style.setProperty('--el-text-color-regular', '#cfd3dc');
    root.style.setProperty('--el-text-color-secondary', '#a3a6ad');
    root.style.setProperty('--el-text-color-placeholder', '#8d9095');
    root.style.setProperty('--el-border-color', '#4c4d4f');
    root.style.setProperty('--el-border-color-light', '#414243');
    root.style.setProperty('--el-fill-color-blank', '#1d1e1f');
    root.style.setProperty('--el-fill-color-light', '#262727');
    root.style.setProperty('--admin-page-bg', '#0a0a0a');
    root.style.setProperty('--admin-card-bg', '#1a1a1a');
    root.style.setProperty('--admin-header-bg', '#141414');
    root.style.setProperty('--admin-menu-hover', '#2a2a2a');
    root.style.setProperty('--admin-menu-active-bg', 'rgba(64, 128, 255, 0.18)');
  } else {
    root.style.setProperty('--el-bg-color', '#ffffff');
    root.style.setProperty('--el-bg-color-page', '#f2f3f5');
    root.style.setProperty('--el-bg-color-overlay', '#ffffff');
    root.style.setProperty('--el-text-color-primary', '#303133');
    root.style.setProperty('--el-text-color-regular', '#606266');
    root.style.setProperty('--el-text-color-secondary', '#909399');
    root.style.setProperty('--el-text-color-placeholder', '#a8abb2');
    root.style.setProperty('--el-border-color', '#dcdfe6');
    root.style.setProperty('--el-border-color-light', '#e4e7ed');
    root.style.setProperty('--el-fill-color-blank', '#ffffff');
    root.style.setProperty('--el-fill-color-light', '#f5f7fa');
    root.style.setProperty('--admin-page-bg', '#ffffff');
    root.style.setProperty('--admin-card-bg', '#fafafa');
    root.style.setProperty('--admin-header-bg', '#ffffff');
    root.style.setProperty('--admin-menu-hover', '#f0f2f5');
    root.style.setProperty('--admin-menu-active-bg', 'rgba(40, 90, 200, 0.12)');
  }
}

export const useThemeStore = defineStore('theme', {
  state: () => ({
    mode: (() => {
      try {
        const saved = localStorage.getItem(STORAGE_KEY);
        return saved === 'dark' ? 'dark' : 'light';
      } catch {
        return 'light';
      }
    })(),
  }),
  getters: {
    isDark: (s) => s.mode === 'dark',
  },
  actions: {
    init() {
      applyThemeMode(this.mode);
    },
    setMode(mode) {
      this.mode = mode === 'dark' ? 'dark' : 'light';
      applyThemeMode(this.mode);
      try {
        localStorage.setItem(STORAGE_KEY, this.mode);
      } catch {
        /* ignore */
      }
    },
    toggle() {
      this.setMode(this.mode === 'dark' ? 'light' : 'dark');
    },
  },
});
