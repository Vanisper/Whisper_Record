<script lang="ts" setup>
import WindowHeader from "./components/WindowHeader.vue";
import WindowSider from "./components/WindowSider.vue";
import DarkMode from "./components/buttons/DarkMode.vue";

import { useDarkThemeHook, useThemeOverridesHook } from "./hooks";

import { Icon } from "@vicons/utils";
import LogoImage from "./assets/logo.svg";

// 暗黑主题
const darkTheme = useDarkThemeHook()

// 主题配置
const overridesTheme = useThemeOverridesHook()
</script>

<template>
  <n-config-provider :theme="darkTheme" :theme-overrides="overridesTheme"
    style="display: flex;flex-direction: column;height: 100%;">
    <!-- 当setup返回的是一个异步组件时,父级组件调用时,需要使用Suspense包裹异步组件 -->
    <Suspense>
      <window-header>
        <template #extend-left>
          <Icon size="24" style="margin: 0 10px 0 16px;">
            <LogoImage />
          </Icon>
          <span style="font-size: 15px;">微记</span>
        </template>
        <template #extend-right>
          <dark-mode style="zoom: 0.13333" />
        </template>
      </window-header>
    </Suspense>
    <main style="background-color: var(--background-secondary);display: flex;flex-grow: 1;overflow: hidden;">
      <window-sider></window-sider>
    </main>
  </n-config-provider>
</template>

<style></style>
