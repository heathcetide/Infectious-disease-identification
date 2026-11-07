<template>
  <router-view v-slot="{ Component }">
    <keep-alive :max="8" :exclude="['Screen']">
      <component :is="Component" />
    </keep-alive>
  </router-view>
</template>

<script>
export default {
  name: 'ReportLayout',
}
</script>
