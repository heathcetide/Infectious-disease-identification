<script>
import { tip } from '../../utils/tip'
import { nearbyHospitalsService } from '../../services/hospital'
import { getStatusBarHeight } from '../../utils/system'

export default {
  data() {
    return {
      loading: true,
      list: [],
      diseaseHint: '',
      lat: 0,
      lng: 0,
      statusBarHeight: 20,
      // In-page map panel (no extra route — avoids empty stub pages in mp build)
      mapOpen: false,
      mapName: '',
      mapAddress: '',
      mapPhone: '',
      mapLat: 36.4,
      mapLng: 94.9,
      mapMarkers: [],
    }
  },
  onLoad(options) {
    this.statusBarHeight = getStatusBarHeight(20)
    this.diseaseHint = decodeURIComponent(options?.disease || '')
    this.locateAndLoad()
  },
  onBackPress() {
    if (this.mapOpen) {
      this.closeMap()
      return true
    }
    return false
  },
  methods: {
    goBack() {
      if (this.mapOpen) {
        this.closeMap()
        return
      }
      const pages = getCurrentPages()
      if (pages.length > 1) {
        uni.navigateBack({ delta: 1 })
      } else {
        uni.redirectTo({ url: '/pages/clock/clock' })
      }
    },
    locateAndLoad() {
      this.loading = true
      uni.getLocation({
        type: 'gcj02',
        success: (res) => {
          this.lat = res.latitude
          this.lng = res.longitude
          this.fetchList()
        },
        fail: () => {
          this.lat = 36.4
          this.lng = 94.9
          this.fetchList()
        },
      })
    },
    async fetchList() {
      try {
        const res = await nearbyHospitalsService({ lat: this.lat, lng: this.lng, limit: 20 })
        const ok = res && (res.code === 200 || res.code === 1)
        this.list = ok ? res.data?.list || [] : []
      } catch (e) {
        console.error(e)
        tip.fail('加载医院失败')
        this.list = []
      } finally {
        this.loading = false
      }
    },
    hasDistance(item) {
      const d = Number(item?.distanceKm)
      return Number.isFinite(d) && d >= 0 && d < 99999999
    },
    formatDistance(item) {
      const d = Number(item?.distanceKm) || 0
      if (d < 1) return `${Math.round(d * 1000)}m`
      return `${d.toFixed(1)}km`
    },
    itemAt(e) {
      const idx = Number(e?.currentTarget?.dataset?.index)
      if (!Number.isFinite(idx) || idx < 0 || idx >= this.list.length) return null
      return this.list[idx]
    },
    callPhone(e) {
      const item = this.itemAt(e)
      const phone = item?.hospitalPhoneNumber
      if (!phone) {
        tip.fail('暂无电话')
        return
      }
      uni.makePhoneCall({ phoneNumber: String(phone).replace(/-/g, '') })
    },
    openMap(e) {
      const item = this.itemAt(e)
      if (!item) {
        tip.fail('医院信息无效')
        return
      }
      const lat = Number(item.latitude)
      const lng = Number(item.longitude)
      if (!Number.isFinite(lat) || !Number.isFinite(lng) || (lat === 0 && lng === 0)) {
        tip.fail('暂无坐标')
        return
      }
      this.mapName = item.hospitalName || '医院'
      this.mapAddress = item.address || ''
      this.mapPhone = item.hospitalPhoneNumber || ''
      this.mapLat = lat
      this.mapLng = lng
      this.mapMarkers = [
        {
          id: 1,
          latitude: lat,
          longitude: lng,
          title: this.mapName,
          width: 28,
          height: 28,
          callout: {
            content: this.mapName,
            display: 'ALWAYS',
            padding: 8,
            borderRadius: 6,
            fontSize: 13,
          },
        },
      ]
      this.mapOpen = true
    },
    closeMap() {
      this.mapOpen = false
    },
    callMapPhone() {
      if (!this.mapPhone) {
        tip.fail('暂无电话')
        return
      }
      uni.makePhoneCall({ phoneNumber: String(this.mapPhone).replace(/-/g, '') })
    },
    openSystemMap() {
      uni.openLocation({
        latitude: this.mapLat,
        longitude: this.mapLng,
        name: this.mapName,
        address: this.mapAddress || this.mapName,
        scale: 15,
        fail: (err) => {
          console.warn('openLocation fail', err)
          tip.fail('当前环境无法打开系统地图，请直接查看上方地图位置')
        },
      })
    },
  },
}
</script>

<template>
  <view class="page">
    <view class="nav" :style="{ paddingTop: statusBarHeight + 'px' }">
      <view class="nav-inner">
        <view class="back" @tap="goBack">
          <text class="back-arrow">‹</text>
          <text class="back-text">返回</text>
        </view>
        <text class="nav-title">{{ mapOpen ? '医院位置' : '就近医院' }}</text>
        <view class="nav-right"></view>
      </view>
    </view>

    <!-- List -->
    <view class="body" v-show="!mapOpen">
      <view class="banner" v-if="diseaseHint">
        <text class="banner-title">高风险提示</text>
        <text class="banner-desc">「{{ diseaseHint }}」评分可能较高，请优先前往就近医院就诊。</text>
      </view>
      <view class="hint">以下医院来自后台医院基础数据，按与您当前位置距离排序。</view>

      <view v-if="loading" class="empty">加载中…</view>
      <view v-else-if="!list.length" class="empty">暂无医院数据</view>
      <view v-else class="list">
        <view
          class="card"
          v-for="(item, i) in list"
          :key="'h-' + (item.hospitalId != null ? item.hospitalId : i)"
        >
          <view class="row-top">
            <text class="name">{{ item.hospitalName }}</text>
            <text class="dist" v-if="hasDistance(item)">{{ formatDistance(item) }}</text>
          </view>
          <text class="addr">{{ item.address || '地址待完善' }}</text>
          <text class="phone">{{ item.hospitalPhoneNumber || '电话待完善' }}</text>
          <view class="actions">
            <button class="btn ghost" size="mini" :data-index="i" @tap="callPhone">拨打电话</button>
            <button class="btn primary" size="mini" :data-index="i" @tap="openMap">导航前往</button>
          </view>
        </view>
      </view>
    </view>

    <!-- Inline map panel -->
    <view class="map-panel" v-if="mapOpen">
      <map
        class="map"
        :latitude="mapLat"
        :longitude="mapLng"
        :scale="15"
        :markers="mapMarkers"
        :show-location="true"
        enable-scroll
        enable-zoom
      />
      <view class="sheet">
        <text class="sheet-name">{{ mapName }}</text>
        <text class="sheet-addr">{{ mapAddress || '地址待完善' }}</text>
        <text class="sheet-coord">{{ mapLat.toFixed(5) }}, {{ mapLng.toFixed(5) }}</text>
        <view class="actions">
          <button v-if="mapPhone" class="btn ghost" size="mini" @tap="callMapPhone">拨打电话</button>
          <button class="btn primary" size="mini" @tap="openSystemMap">打开系统地图</button>
          <button class="btn ghost" size="mini" @tap="closeMap">关闭</button>
        </view>
      </view>
    </view>
  </view>
</template>

<style>
.page {
  min-height: 100vh;
  background: #eef3fb;
  box-sizing: border-box;
}
.nav {
  background: #fff;
  border-bottom: 1px solid #e2e8f0;
  position: sticky;
  top: 0;
  z-index: 20;
}
.nav-inner {
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 12px;
}
.back {
  min-width: 72px;
  display: flex;
  align-items: center;
  gap: 2px;
}
.back-arrow {
  font-size: 28px;
  line-height: 1;
  color: #0f172a;
  margin-top: -2px;
}
.back-text {
  font-size: 15px;
  color: #0f172a;
}
.nav-title {
  flex: 1;
  text-align: center;
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
}
.nav-right {
  min-width: 72px;
}
.body {
  padding: 16px;
  box-sizing: border-box;
}
.banner {
  background: #fff1f0;
  border: 1px solid #fca5a5;
  border-radius: 12px;
  padding: 12px 14px;
  margin-bottom: 12px;
}
.banner-title {
  display: block;
  font-size: 15px;
  font-weight: 700;
  color: #b91c1c;
}
.banner-desc {
  display: block;
  margin-top: 6px;
  font-size: 13px;
  color: #7f1d1d;
  line-height: 1.5;
}
.hint {
  font-size: 12px;
  color: #64748b;
  margin-bottom: 12px;
}
.card {
  background: #fff;
  border-radius: 12px;
  padding: 14px;
  margin-bottom: 10px;
  box-shadow: 0 4px 14px rgba(40, 90, 200, 0.06);
}
.row-top {
  display: flex;
  justify-content: space-between;
  gap: 8px;
  align-items: flex-start;
}
.name {
  flex: 1;
  font-size: 15px;
  font-weight: 700;
  color: #0f172a;
}
.dist {
  font-size: 12px;
  color: #285ac8;
  font-weight: 600;
}
.addr,
.phone {
  display: block;
  margin-top: 6px;
  font-size: 13px;
  color: #64748b;
}
.actions {
  display: flex;
  gap: 8px;
  margin-top: 12px;
  flex-wrap: wrap;
}
.btn {
  margin: 0;
  border-radius: 8px;
}
.btn::after {
  border: none;
}
.btn.ghost {
  background: #f1f5f9;
  color: #334155;
}
.btn.primary {
  background: #285ac8;
  color: #fff;
}
.empty {
  padding: 48px 0;
  text-align: center;
  color: #64748b;
}
.map-panel {
  display: flex;
  flex-direction: column;
  min-height: calc(100vh - 88px);
}
.map {
  width: 100%;
  height: 70vh;
}
.sheet {
  flex: 1;
  background: #fff;
  border-radius: 16px 16px 0 0;
  margin-top: 0;
  padding: 20px 16px calc(28px + env(safe-area-inset-bottom));
  box-shadow: 0 -2px 12px rgba(15, 23, 42, 0.06);
  position: relative;
  z-index: 2;
}
.sheet-name {
  display: block;
  font-size: 17px;
  font-weight: 700;
  color: #0f172a;
}
.sheet-addr,
.sheet-coord {
  display: block;
  margin-top: 8px;
  font-size: 13px;
  color: #64748b;
  line-height: 1.45;
}
</style>
