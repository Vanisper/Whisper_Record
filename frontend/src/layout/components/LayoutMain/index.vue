<template>
  <router-view>
    <!-- 缓存处理 -->
    <template #default="{ Component, route }">
      <component v-if='!route.meta.keepAlive' :is="Component" :key="route.fullPath"></component>
      <keep-alive v-else>
        <component :is="Component" :key="route.fullPath"></component>
      </keep-alive>
    </template>
  </router-view>
  <!-- <router-view v-if="!$route.meta.keepAlive" />
  <router-view v-if="$route.meta.keepAlive" v-slot="{ Component }">
    <keep-alive>
      <component :is="Component" />
    </keep-alive>
  </router-view> -->
</template>