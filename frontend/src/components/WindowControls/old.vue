<template>
  <!-- 置顶——状态切换 -->
  <span :title="isOntoped ? '取消置顶' : '置顶'" class="sticky iconfont icon-zhiding" @click="onTopWindow"></span>
  <span title="最小化" class="minimize iconfont icon-zuixiaohua" @click="WindowMinimise"></span>
  <!-- 最大化应用程序——状态切换 -->
  <span :title="isMaximised ? '取消最大化' : '最大化'" class="maximize iconfont"
        :class="isMaximised ? 'icon-shouqi' : 'icon-quanping'" @click="maximiseWindow"></span>
  <span title="关闭" class="close iconfont icon-guanbi" @click="closeWindow"></span>
</template>

<script setup lang="ts">
import {CloseWindow, MaximiseWindow, OnTopWindow, WindowIsOnToped} from "../../../wailsjs/go/backend/App";
import {ref} from "vue";
import {WindowIsMaximised, WindowMinimise} from "../../../wailsjs/runtime";

async function closeWindow() {
  // wails 运行时关闭程序 传参true表示隐藏窗口而不关闭
  await CloseWindow(false)
}

async function maximiseWindow() {
  await MaximiseWindow()
  isMaximised.value = !isMaximised.value
}

async function onTopWindow() {
  // OnTopWindow
  await OnTopWindow()
  isOntoped.value = await WindowIsOnToped()
  if (isOntoped.value) {
    style.value.stickyRotate = "-45deg"
    style.value.stickyColor = "limegreen"
  } else {
    style.value.stickyRotate = "0deg"
    style.value.stickyColor = "var(--text-primary)"
  }
}

let isMaximised = ref(await WindowIsMaximised())
let isOntoped = ref(await WindowIsOnToped())
const style = ref({
  stickyRotate: isOntoped.value ? "-45deg" : "0deg",
  stickyColor: isOntoped.value ? "limegreen" : "var(--text-primary)",
})

window.onresize = async (event) => {
  isMaximised.value = await WindowIsMaximised()
}
</script>

<style scoped lang="less">
@import url("../../main.less");
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
.sticky {
  //color: @sticky-color;
  color: v-bind("style.stickyColor");
  width: 43px;

  &::before {
    transition: transform 300ms;
    transform: translate(-50%, -50%) rotate(v-bind("style.stickyRotate"));
  }

  &:hover {
    background-color: var(--button-hover);
    //color: @background-primary;
  }
}

.minimize {
  color: var(--text-primary);
  width: 43px;

  &:hover {
    background-color: var(--button-hover);
    //color: @background-primary;
  }
}

.maximize {
  color: var(--text-primary);
  width: 43px;

  &:hover {
    background-color: var(--button-hover);
    //color: @background-primary;
  }
}

.close {
  color: var(--text-primary);
  width: 41px;

  &:hover {
    background-color: @close-hover-color;
    color: @background-primary;
  }
}
</style>