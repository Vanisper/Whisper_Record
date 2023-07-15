<template>
  <onTop />
  <div title="最小化" class="frameless-titlebar-button frameless-titlebar-minimize" @click="WindowMinimise">
    <svg width="10" height="10">
      <path v-for="(d, index) in minimizePath" :d="d" />
    </svg>
  </div>
  <div title="最大化" class="frameless-titlebar-button frameless-titlebar-toggle" @click="maximiseWindow">
    <svg width="10" height="10">
      <path v-for="(d, index) in (isMaximised ? restorePath_win11 : maximizePath_win11)" :d="d" />
    </svg>
  </div>
  <div title="关闭" class="frameless-titlebar-button frameless-titlebar-close" @click="closeWindow">
    <svg width="10" height="10">
      <path v-for="(d, index) in closePath" :d="d" />
    </svg>
  </div>
</template>

<script setup lang="ts">
import { minimizePath, maximizePath, closePath, restorePath, restorePath_win11, maximizePath_win11 } from "../../assets/window-controls";
import { CloseWindow, MaximiseWindow, OnTopWindow, WindowIsOnToped } from "../../../wailsjs/go/backend/App";
import { ref } from "vue";
import { WindowIsMaximised, WindowMinimise } from "../../../wailsjs/runtime";
import onTop from "./items/onTop.vue";

async function closeWindow() {
  // wails 运行时关闭程序 传参true表示隐藏窗口而不关闭
  await CloseWindow(false)
}

async function maximiseWindow() {
  await MaximiseWindow()
  isMaximised.value = !isMaximised.value
}

const isMaximised = ref(await WindowIsMaximised())
const isOntoped = ref(await WindowIsOnToped())

window.onresize = async (event) => {
  isMaximised.value = await WindowIsMaximised()
}
</script>

<style scoped lang="less">
@import url("../../main.less");

.frameless-titlebar-button {
  position: relative;
  display: block;
  width: 46px;
  height: @header-height;
  fill: var(--text-primary);
  cursor: pointer;

  svg {
    position: absolute;
    display: inline-flex;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
  }

  &.frameless-titlebar-close:hover {
    background-color: rgb(228, 79, 79);

    svg {
      fill: #fff;
    }
  }

  &.frameless-titlebar-minimize:hover,
  &.frameless-titlebar-toggle:hover {
    background-color: var(--button-hover);
  }
}
</style>