<template>
  <div class="button-box">
    <div class="sun" @click="switchDark">
      <div class="line"></div>
      <div class="line"></div>
      <div class="line"></div>
      <div class="line"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useDesignStore} from '../../store/modules/designStore/designStore'
import {setHtmlTheme} from '../../utils'

const designStore = useDesignStore()


const switchDark = (event: Event) => {
  const target = event.target as HTMLElement
  target.classList.toggle("night")
  // 当前点击位置
  const rect = target.getBoundingClientRect()
  // 窗口buffer数据
  const bufferData: {
    buffer: Uint8Array;
    width: number;
    height: number;
  } | null = null
  designStore.changeTheme()
  setHtmlTheme()
}
</script>

<style scoped lang="less">
.button-box {
  width: 240px;
  height: 240px;
  background-color: #e91e63;
  border-radius: 6px;

  transform: rotate(0);

  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;

  .sun {
    width: 100px;
    height: 100px;
    background-color: #fff;
    border-radius: 50%;
    cursor: pointer;
    position: relative;
    border: 5px solid #e91e63;

    &::after {
      content: "";
      width: 100px;
      height: 100px;
      background-color: #e91e63;
      border-radius: 50%;
      position: absolute;
      top: -100px;
      right: -100px;
      transition: 0.5s;
    }


    div {
      height: 3px;
      background-color: #fff;
      position: absolute;
      top: 50%;
      left: 50%;

      z-index: -1;

      transition: width 0.5s;

      &.line:nth-child(1) {
        width: 160px;
        transform: translate(-50%, -50%);
      }

      &.line:nth-child(2) {
        width: 160px;
        transform: translate(-50%, -50%) rotate(90deg);
      }

      &.line:nth-child(3) {
        width: 140px;
        transform: translate(-50%, -50%) rotate(45deg);
      }

      &.line:nth-child(4) {
        width: 140px;
        transform: translate(-50%, -50%) rotate(-45deg);
      }
    }

    &.night {
      &::after {
        top: -40px;
        right: -40px;
      }

      div.line {
        width: 0;
      }
    }
  }
}
</style>