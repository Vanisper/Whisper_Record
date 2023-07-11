<template>
  <header id="header" style="--wails-draggable: drag">
    <div class="header-wrapper">
      <div class="header-left" @dblclick="maximiseWindow">
        <slot name="extend-left" />
      </div>
      <div class="header-center">
        <slot name="extend-center" />
      </div>
      <div class="header-right" style="--wails-draggable: none">
        <slot name="extend-right" />
        <window-controls />
        <!-- <old-window-controls /> -->
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { MaximiseWindow } from "../../wailsjs/go/backend/App";
import { onMounted, ref } from "vue";
import { WindowIsMaximised } from "../../wailsjs/runtime";
import WindowControls from "./WindowControls/index.vue";
import OldWindowControls from "./WindowControls/old.vue";

async function maximiseWindow() {
  await MaximiseWindow();
  isMaximised.value = !isMaximised.value;
}

let isMaximised = ref(await WindowIsMaximised());

window.onresize = async (event) => {
  isMaximised.value = await WindowIsMaximised();
};

onMounted(() => {
  // (window as any).wails.flags.deferDragToMouseMove = true;
});
</script>

<style scoped lang="less">
@import url("../main.less");

header {
  width: 100vw;
  overflow: hidden;
  height: @header-height;
  //background-color: transparent;
  //background-color: @background-secondary;
  background-color: var(--background-secondary);
  display: flex;
  align-items: center;
  user-select: none;
  position: relative;

  &::after {
    content: "";
    position: absolute;
    bottom: 0;
    height: 1px;
    width: 100%;
    background-color: var(--border-primary);
  }

  .header-wrapper {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;

    .header-left {
      margin-right: auto;
      display: flex;
      align-items: center;
    }

    .header-right {
      margin-left: auto;
      display: flex;
      flex-direction: row;

      span {
        cursor: pointer;
        display: inline-block;
        height: 100%;
        //width: 47px;
        overflow: hidden;
        padding: 2px;
        position: relative;
        transition: background-color 300ms;

        &::before {
          position: absolute;
          top: 50%;
          left: 50%;
          transform: translate(-50%, -50%);
        }
      }
    }
  }
}
</style>
