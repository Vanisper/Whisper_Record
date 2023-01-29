<template>
  <router-view>
    <!-- 缓存处理 -->
    <template #default="{ Component, route }">
      <component
          v-if='!route.meta["keep-alive"]'
          :is="Component"
          :key="route.fullPath"
      ></component>
      <keep-alive v-else>
        <component :is="Component" :key="route.fullPath"></component>
      </keep-alive>
    </template>
  </router-view>
</template>