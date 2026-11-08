<template>
  <div>
    <div class="positionInfo">
      <span>经度: {{ longitude }}°</span>
      <span>纬度: {{ latitude }}°</span>
      <span>地点: {{ address }}</span>
    </div>
    <div ref="mapContainer" class="map-container"></div>
  </div>
</template>

<script>
import locationIcon from '@/assets/location.svg';
import { loadAMap } from '@/utils/amap';

export default {
  name: 'MapComponent',
  data() {
    return {
      longitude: 94.3668, // 默认经度
      latitude: 29.6524, // 默认纬度
      address: '', // 地址信息
      map: null, // 保存地图实例
      marker: null, // 保存标记实例
    };
  },
  methods: {

    getPosition(longitude, latitude, fallbackAddress = '') {
        this.destroyMap();
        this.longitude = longitude;
        this.latitude = latitude;
        this.address = fallbackAddress || '';
        this.initializeMap(this.longitude, this.latitude, fallbackAddress);
    },
    async initializeMap(longitude, latitude, fallbackAddress = '') {
      const AMap = await loadAMap();
      const LatLng = [longitude, latitude];

      this.map = new AMap.Map(this.$refs.mapContainer, {
        center: LatLng,
        zoom: 13,
        resizeEnable: true,
      });

      const customIcon = new AMap.Icon({
        size: new AMap.Size(48, 48),
        image: locationIcon,
        imageSize: new AMap.Size(48, 48),
      });

      this.marker = new AMap.Marker({
        position: LatLng,
        icon: customIcon,
        map: this.map,
      });

      this.getAddress(AMap, LatLng, fallbackAddress);
    },

    getAddress(AMap, latlng, fallbackAddress = '') {
      if (fallbackAddress) {
        this.address = fallbackAddress;
      }
      const applyFallback = () => {
        if (!this.address) {
          this.address = fallbackAddress || `${latlng[1]}, ${latlng[0]}`;
        }
      };
      try {
        // AMap JSAPI 2.0 requires plugin load; bare `new AMap.Geocoder()` often
        // surfaces "please check the key and referer".
        AMap.plugin('AMap.Geocoder', () => {
          try {
            const geocoder = new AMap.Geocoder();
            geocoder.getAddress(latlng, (status, result) => {
              if (status === 'complete' && result?.regeocode?.formattedAddress) {
                this.address = result.regeocode.formattedAddress;
              } else {
                applyFallback();
              }
            });
          } catch (e) {
            applyFallback();
          }
        });
      } catch (e) {
        applyFallback();
      }
    },

    // 清理地图资源，防止内存泄露
    destroyMap() {
      if (this.map) {
        this.map.destroy(); // 销毁地图实例
        this.map = null;
        this.marker = null;
      }
    }
  },
  mounted() {
    // this.initializeMap(this.longitude,this.latitude); // 页面加载时立即初始化地图
  },
  beforeDestroy() {
    this.destroyMap(); // 在组件销毁前清理地图
  }
};
</script>

<style scoped>
.map-container {
  width: 100%;
  height: 500px;
  border-radius: 5px;
}
.positionInfo {
  font-size: 14px;
  margin-bottom: 10px;
}
span {
  margin-right: 30px;
}
</style>
