<template>
  <div style="--wails-draggable: drag" class="side-bar">
    <!-- <n-button text style="
        font-size: 26px;
        --n-text-color-hover: #ffdb00;
        --n-text-color-pressed: #ffdb00;
        --n-text-color-focus: #ffdb00;
      " data-route-path="/home" @click="goTo" :focusable="false">
      <n-icon>
        <article-icon />
      </n-icon>
    </n-button> -->
    <n-button text style="
        font-size: 24px;
        --n-text-color-hover: #ffdb00;
        --n-text-color-pressed: #ffdb00;
        --n-text-color-focus: #ffdb00;
      " data-route-path="/editor" @click="goTo" :focusable="false">
      <n-icon>
        <documents-icon />
      </n-icon>
    </n-button>
    <n-button text style="
        font-size: 24px;
        --n-text-color-hover: #ffdb00;
        --n-text-color-pressed: #ffdb00;
        --n-text-color-focus: #ffdb00;
        margin-top: auto;
        margin-bottom: 10px;
      " @click="OpenPreviewURL" :focusable="false">
      <n-icon>
        <link-to-icon />
      </n-icon>
    </n-button>
    <n-button text style="
        font-size: 24px;
        --n-text-color-hover: #ffdb00;
        --n-text-color-pressed: #ffdb00;
        --n-text-color-focus: #ffdb00;
        margin-bottom: 10px;
      " data-route-path="/setting" @click="goTo" :focusable="false">
      <n-icon>
        <setting-icon />
      </n-icon>
    </n-button>
  </div>
  <n-layout has-sider>
    <n-space :wrap="false" :wrap-item="false" style="width: 100%">
      <n-layout-sider ref="layoutSider" bordered :collapsed="isCollapsed" :on-update:collapsed="collapseHandler"
        collapse-mode="width" :collapsed-width="0" :width="SideBarWidth" content-style="overflow: hidden;"
        show-trigger="bar">
        <div style="height: 100%;user-select: none;display: flex;">
          <!-- {{ route.fullPath }} -->
          <tree :project-tree="projectTree" v-contextmenu.dark="contextmenus" />
        </div>
        <div class="drag-bar" ref="dragBar"></div>
      </n-layout-sider>

      <n-layout-content content-style="margin-right: 0px;position: relative;overflow: hidden;">
        <app-provider />
        <router-view />
      </n-layout-content>
    </n-space>
  </n-layout>
</template>

<script lang="ts" setup>
import { AddFile, AddFolder, RemoveFile, OpenPreviewURL } from "../../wailsjs/go/backend/App";
import { AppProvider } from "../components/AppProvider/index";
import Tree from "./sideBar/Tree.vue"
import { DocumentsOutline as DocumentsIcon, SettingsOutline as SettingIcon, } from "@vicons/ionicons5";
import { DocumentOnePage20Regular as ArticleIcon, CubeLink20Regular as LinkToIcon } from "@vicons/fluent";
import { useRouter, useRoute } from "vue-router";
import { onMounted, ref, watch, reactive, computed } from "vue";
import { IClickMenuItem } from "web-contextmenu/type/ContextMenuType";
import { usePostsStore } from "../store/modules/postsStore/postsStore";

const projectTree = computed(() => usePostsStore().projectTree!)

const contextmenus = ref<IClickMenuItem[]>([
  {
    text: "新建文件",
    // subText: "CREATE_FILE",
    action: async (el, event, axis, menus, item, isDark) => {
      const target = event.target as HTMLElement;
      const regex = /\.md$/;

      if (target.dataset.filepath && target.dataset.filepath !== "") {
        const path = target.dataset.filepath.replace(/\\/g, '/')
        if (regex.test(path)) {
          const folderPath = path.split('/').slice(0, -1).join('/');
          const name = prompt("输入文件名");
          if (name) {
            if (await AddFile(folderPath + "/" + name + ".md")) {
              window.$message.success("文件创建成功")
            } else {
              window.$message.error("文件创建失败")
            }
            await usePostsStore().updateProjectTree()
          } else {
            window.$message.info("取消文件创建")
          }
        }
      } else if (target.dataset.folderpath && target.dataset.folderpath !== "") {
        const path = target.dataset.folderpath.replace(/\\/g, '/')
        const folderPath = path
        const name = prompt("输入文件名");
        if (name) {
          if (await AddFile(folderPath + "/" + name + ".md")) {
            window.$message.success("文件创建成功")
          } else {
            window.$message.error("文件创建失败")
          }
          await usePostsStore().updateProjectTree()
        } else {
          window.$message.info("取消文件创建")
        }
      } else if (target.dataset.root == "true" && target.dataset.rootpath && target.dataset.rootpath !== "") {
        const path = target.dataset.rootpath.replace(/\\/g, '/')
        const folderPath = path
        const name = prompt("输入文件名");
        if (name) {
          if (await AddFile(folderPath + "/" + name + ".md")) {
            window.$message.success("文件创建成功")
          } else {
            window.$message.error("文件创建失败")
          }
          await usePostsStore().updateProjectTree()
        } else {
          window.$message.info("取消文件创建")
        }
      }
    },
  },
  {
    text: "新建文件夹",
    // subText: "CREATE_FOLDER",
    action: async (el, event, axis, menus, item, isDark) => {
      const target = event.target as HTMLElement;

      if (target.dataset.filepath && target.dataset.filepath !== "") {
        const path = target.dataset.filepath.replace(/\\/g, '/')
        const regex = /\.md$/;
        if (regex.test(path)) {
          const folderPath = path.split('/').slice(0, -1).join('/');
          const name = prompt("输入文件夹名");
          if (name) {
            if (await AddFolder(folderPath + "/" + name)) {
              window.$message.success("文件夹创建成功")
            } else {
              window.$message.error("文件夹创建失败")
            }
            await usePostsStore().updateProjectTree()
          } else {
            window.$message.info("取消文件夹创建")
          }
        }
      } else if (target.dataset.folderpath && target.dataset.folderpath !== "") {
        const path = target.dataset.folderpath.replace(/\\/g, '/')
        const folderPath = path
        const name = prompt("输入文件夹名");
        if (name) {
          if (await AddFolder(folderPath + "/" + name)) {
            window.$message.success("文件夹创建成功")
          } else {
            window.$message.error("文件夹创建失败")
          }
          await usePostsStore().updateProjectTree()
        } else {
          window.$message.info("取消文件夹创建")
        }
      } else if (target.dataset.root == "true" && target.dataset.rootpath && target.dataset.rootpath !== "") {
        const path = target.dataset.rootpath.replace(/\\/g, '/')
        const folderPath = path
        const name = prompt("输入文件夹名");
        if (name) {
          if (await AddFolder(folderPath + "/" + name)) {
            window.$message.success("文件夹创建成功")
          } else {
            window.$message.error("文件夹创建失败")
          }
          await usePostsStore().updateProjectTree()
        } else {
          window.$message.info("取消文件夹创建")
        }
      }
    },
  },
  { divider: true },
  // {
  //   text: "多级菜单",
  //   children: [
  //     { text: "子菜单1" },
  //     { text: "子菜单2" },
  //     {
  //       text: "三级菜单",
  //       children: [{ text: "子菜单1" }, { text: "子菜单2" }],
  //     },
  //   ],
  // },
  // { divider: true },
  {
    text: "删除",
    // subText: "CTRL + X",
    action: async (el, event, axis, menus, item, isDark) => {
      if (!confirm("你确定要删除这个文件（夹）吗？")) {
        window.$message.info("取消删除")
        return false
      }
      const target = event.target as HTMLElement;
      if (target.dataset.filepath && target.dataset.filepath !== "") {
        if (await RemoveFile(target.dataset.filepath)) {
          window.$message.success("删除成功")
        } else {
          window.$message.error("删除失败")
        }
        await usePostsStore().updateProjectTree()
      } else if (target.dataset.folderpath && target.dataset.folderpath !== "") {
        if (await RemoveFile(target.dataset.folderpath)) {
          window.$message.success("删除成功")
        } else {
          window.$message.error("删除失败")
        }
        await usePostsStore().updateProjectTree()
      }
    },
  },
]);

const router = useRouter();
const route = useRoute();

const goTo = ($event: PointerEvent) => {
  const target = $event.currentTarget as HTMLButtonElement;
  router.push(target.dataset.routePath || "/");
};

const dragBar = ref<HTMLElement>();
const layoutSider = ref<any>();
const SideBarWidth = ref(240);
// const isCollapsed = computed(() => SideBarWidth.value <= 50 ? true : false);
const isCollapsed = ref(false);
const collapseHandler = (collapsed: boolean) => {
  isCollapsed.value = collapsed;
  !collapsed && (SideBarWidth.value = 240);
}
watch(SideBarWidth, (value) => {
  if (value <= 50) {
    isCollapsed.value = true;
  }
})
onMounted(async () => {
  await usePostsStore().updateProjectTree()
  console.log(usePostsStore().projectTree);

  let startX = 0
  let sideBarWidth = SideBarWidth.value
  let startWidth = SideBarWidth.value
  const oldTransition = getComputedStyle((layoutSider.value.$el as HTMLElement), null).transition;
  const newTransition = oldTransition.match(/[^,]+?\s+cubic-bezier\(.+?\)/g)?.filter(item => !/\bwidth\b/.test(item)).join(",");
  const mouseUpHandler = (_event: MouseEvent) => {
    (layoutSider.value.$el as HTMLElement).style.transition = "";
    document.removeEventListener('mousemove', mouseMoveHandler, false)
    document.removeEventListener('mouseup', mouseUpHandler, false)
  }

  const mouseMoveHandler = (event: MouseEvent) => {
    (layoutSider.value.$el as HTMLElement).style.transition = newTransition!;
    const offset = event.clientX - startX;
    sideBarWidth = startWidth + offset;
    SideBarWidth.value = sideBarWidth;
  }
  const mouseDownHandler = (event: MouseEvent) => {
    startX = event.clientX
    startWidth = SideBarWidth.value
    document.addEventListener('mousemove', mouseMoveHandler, false)
    document.addEventListener('mouseup', mouseUpHandler, false)
    console.log(startX, startWidth);
  }

  dragBar.value?.addEventListener('mousedown', mouseDownHandler, false)
})


</script>

<style lang="less" scoped>
.side-bar {
  width: 45px;
  min-width: 45px;
  position: relative;
  display: flex;
  align-items: center;
  flex-direction: column;
  padding-top: 10px;

  &::after {
    content: "";
    position: absolute;
    right: 0;
    top: 0;
    height: 100%;
    width: 1px;
    background-color: var(--border-primary);
  }

  button {
    width: 100%;
    height: 45px;
  }
}

.drag-bar {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  height: 100%;
  width: 5px;
  cursor: e-resize;
  transition: background-color .3s cubic-bezier(.4, 0, .2, 1);

  &:hover {
    // #ffdb00  var(--scrollbar-color);
    background-color: #ffdd007e;
  }
}
</style>
