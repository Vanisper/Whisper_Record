<template>
  <header id="header" style="--wails-draggable:drag">
    <div class="header-wrapper">
      <span class="sticky iconfont icon-zhiding" @click="OnTopWindow"></span>
      <span class="minimize iconfont icon-shouqi"></span>
      <span class="maximize iconfont icon-quanping"></span>
      <span class="close iconfont icon-guanbi" @click="closeWindow"></span>
    </div>
  </header>
</template>

<script setup lang="ts">
import {CloseWindow, OnTopWindow} from "../../wailsjs/go/backend/App";
import {onMounted} from "vue";

function closeWindow() {
  CloseWindow(true)
}

onMounted(() => {
  const header = document.querySelector("header") as HTMLElement
  header.addEventListener("mousedown", (event) => {
    console.log(1)
    header.style.cursor = "grabbing"
    event.preventDefault()
  })
  header.addEventListener("mouseup", (event) => {
    console.log(2)
    header.style.cursor = "grab"
    event.preventDefault()
  })
})
</script>

<style scoped lang="less">
@import url("src/main");

header {
  width: 100vw;
  height: 60px;
  background-color: @background-secondary;
  cursor: grab;

  .header-wrapper {
    padding: 5px;

    span {
      cursor: pointer;
      display: inline-block;
      height: 16px;
      width: 16px;
      overflow: hidden;
      padding: 2px;
      border-radius: 50%;
      position: relative;

      &::before {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
      }
    }

    .sticky {
      background-color: pink;
    }

    .minimize {
      background-color: forestgreen;
    }

    .maximize {
      background-color: deepskyblue;
    }

    .close {
      background-color: orangered;
    }
  }


}
</style>