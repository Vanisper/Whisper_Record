<template>
  <div class="v-contextmenu" v-show="status" :style="{
    left: style.left,
    top: style.top,
  }" @contextmenu.prevent>
    <ContextmenuContent :menus="menus" :isDark="isDark" :subMenuPosition="style.subMenuPosition"
      :clickMenuItem="clickMenuItem" />
  </div>
</template>

<script lang="ts" setup>
import {
  ComponentInternalInstance,
  computed,
  getCurrentInstance,
  nextTick,
  onBeforeUnmount,
  ref,
  onMounted,
} from "vue";

import ContextmenuContent from "./ContextmenuContent.vue";
import { IClickMenuItem } from "./index.d";

const MENU_WIDTH = 170;
const MENU_HEIGHT = 30;
const MENU_PADDING = 5;
const DIVIDER_HEIGHT = 11;
const SUB_MENU_WIDTH = 120;
const Instance = getCurrentInstance() as ComponentInternalInstance | null;
const props = defineProps({
  axis: {
    type: Object as () => { x: number; y: number },
    default() {
      return { x: 0, y: 0 };
    },
  },
  el: {
    type: HTMLElement,
    default() {
      return null;
    },
  },
  event: {
    type: MouseEvent,
    default() {
      return null;
    },
  },
  menus: {
    type: Array<IClickMenuItem>,
    default() {
      return [{ text: "" }];
    },
  },
  isDark: {
    type: Boolean,
    default: false,
  },
  removeContextMenu: {
    type: Function,
    default() {
      return () => { };
    },
  },
});

const status = ref(false);

const style = computed<{
  left: string;
  top: string;
  subMenuPosition: "right" | "left";
}>(() => {
  const { x, y } = props.axis;

  const normalMenuCount = props.menus.filter(
    (menu) => !menu.divider && !menu.hide
  ).length;
  const dividerMenuCount = props.menus.filter((menu) => menu.divider).length;

  const menuWidth = MENU_WIDTH;
  const menuHeight =
    normalMenuCount * MENU_HEIGHT +
    dividerMenuCount * DIVIDER_HEIGHT +
    MENU_PADDING * 2;

  const maxMenuWidth = MENU_WIDTH + SUB_MENU_WIDTH - 10;

  const screenWidth = document.body.clientWidth;
  const screenHeight = document.body.clientHeight;

  const left = screenWidth <= x + menuWidth ? x - menuWidth : x;
  const top = screenHeight <= y + menuHeight ? y - menuHeight : y;

  const subMenuPosition = screenWidth <= left + maxMenuWidth ? "right" : "left";

  return {
    left: left + "px",
    top: top + "px",
    subMenuPosition,
  };
});

onMounted(() => {
  nextTick(() => (status.value = true));
});

onBeforeUnmount(() => {
  document.body.removeChild(Instance?.proxy?.$el);
});

function clickMenuItem(item: IClickMenuItem) {
  if (item.disable || item.children) return;

  status.value = false;
  item.action && item.action(props.el, props.event, props.axis, props.menus, props.isDark);

  props.removeContextMenu();
}
</script>

<style lang="less">
.v-contextmenu {
  position: fixed;
  z-index: 9999;
  user-select: none;
}
</style>
