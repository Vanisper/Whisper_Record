<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import Editor from "./editor.vue";
import "./lib/styles/index.less";

const msg = ref("");

const $editor = ref<Element>();

onMounted(() => { });
</script>

<template>
  <div class="content">
    <editor ref="$editor" />
  </div>
</template>

<style lang="less" scoped>
.content {
  overflow-x: hidden;
  overflow-y: auto;
  position: absolute;
  height: auto;
  width: inherit;

  left: 0;
  right: 0;
  height: 100%;
  background-color: inherit;

  &::-webkit-scrollbar {
    width: 12px;
    height: 8px;
    background-color: rgba(0, 0, 0, 0);
    position: absolute;
  }

  &::-webkit-scrollbar-corner {
    background: 0 0;
  }

  &::-webkit-scrollbar-thumb {
    background: var(--scrollbar-color);
    // background-clip: padding-box;
  }

  #editor {
    cursor: text;

    position: relative;
    left: 0;
    right: 0;
    min-height: 100%;
    padding-left: 30px;
    padding-right: 30px;
    background-color: inherit;
    // overflow-x: visible;
    overflow: hidden;

    padding-bottom: 70px;
    // transition: 0.4s padding-top ease-out;
  }
}
</style>
