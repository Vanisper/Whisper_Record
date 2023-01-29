<script lang="ts" setup>

import WindowHeader from "./components/WindowHeader.vue";
import DarkMode from "./components/buttons/DarkMode.vue";
import {AppProvider} from "./components/AppProvider/index";

import {NConfigProvider} from "naive-ui";
import {useDarkThemeHook, useThemeOverridesHook} from "./hooks";

// 暗黑主题
const darkTheme = useDarkThemeHook()

// 主题配置
const overridesTheme = useThemeOverridesHook()
</script>

<template>
  <n-config-provider
      :theme="darkTheme"
      :theme-overrides="overridesTheme"
  >
    <!-- 当setup返回的是一个异步组件时,父级组件调用时,需要使用Suspense包裹异步组件 -->
    <Suspense>
      <window-header>
        <template #extend-left>
          logo
        </template>
        <template #extend-center>
          搜索
        </template>
        <template #extend-right>
          <dark-mode style="zoom: 0.2"/>
        </template>
      </window-header>
    </Suspense>
    <app-provider>
      <router-view></router-view>
    </app-provider>
  </n-config-provider>
</template>

<style>

</style>
