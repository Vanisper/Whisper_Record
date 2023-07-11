<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import { IClickMenuItem } from "web-contextmenu/type/ContextMenuType";
import { useDesignStore } from "../../store/modules/designStore/designStore";
import Editor from "./editor.vue";
import { WhisperMarkDown } from "./lib";
import "./lib/styles/index.less";

const msg = ref("");
const contextmenus = ref<IClickMenuItem[]>([
  {
    text: "删除",
    subText: "BACKSPACE",
    action: (el, axis, menus, isDark) => {
      msg.value = "你点击了删除";
      console.log(el, axis, menus, isDark);
    },
  },
  {
    text: "禁用菜单项",
    disable: true,
  },
  { divider: true },
  {
    text: "多级菜单",
    children: [
      { text: "子菜单1" },
      { text: "子菜单2" },
      {
        text: "三级菜单",
        children: [{ text: "子菜单1" }, { text: "子菜单2" }],
      },
    ],
  },
  { divider: true },
  {
    text: "剪切",
    subText: "CTRL + X",
    action: () => (msg.value = "你点击了剪切"),
  },
  {
    text: "复制",
    subText: "CTRL + C",
    action: () => (msg.value = "你点击了复制"),
  },
  {
    text: "粘贴",
    subText: "CTRL + V",
    action: () => (msg.value = "你点击了粘贴"),
  },
]);

const WMD = ref<null | WhisperMarkDown>(null);
const container = ref<Element>();
const options = ref({ readonly: true });
onMounted(() => {
  WMD.value = new WhisperMarkDown(container.value, options.value);
  console.log();
  setTimeout(() => {
    options.value.readonly = false;
  }, 2000);
});

watch(
  options,
  (value) => {
    console.log(value.readonly);
  },
  {
    deep: true,
  }
);
</script>

<template>
  <!-- <div
    v-if="useDesignStore().darkTheme"
    v-contextmenu.dark="contextmenus"
    style="height: 100%"
  >
    <editor />
  </div>
  <div v-else v-contextmenu="contextmenus" style="height: 100%">
    <editor />
  </div> -->
  <div class="content" v-contextmenu.dark="contextmenus">
    <span ref="container" class="ttt"></span>
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
}</style>
