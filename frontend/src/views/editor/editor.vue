<script lang="ts" setup>
import { usePostsStore } from "../../store/modules/postsStore/postsStore";
import { ProseMirrorView } from "./EditorView";
import "./editor.css";

import { onMounted, ref } from "vue";

const $editor = ref<HTMLElement>();

onMounted(() => {
  let view = new ProseMirrorView(
    $editor.value!, ""
  );
  window.view = view;
});
</script>

<template>
  <div v-show="usePostsStore().currentFileTitle" ref="$editor" id="editor"
    style="--bg: var(--n-color); display: flex; flex-direction: column"></div>
</template>

<style lang="less"></style>
