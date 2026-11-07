<template>
  <router-view v-slot="{ Component }">
    <keep-alive :max="8">
      <component :is="Component" />
    </keep-alive>
  </router-view>
</template>

<script>
export default {
  name: 'UserLayout',
}
</script>
