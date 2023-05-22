<template>
    <n-button @click="killPreview">关闭预览进程</n-button>
    <n-button @click="ReStartPreview">重启预览进程</n-button>
    <!-- <pre>
        {{ (usePostsStore().projectTree) }}
    </pre> -->
</template>

<script lang="ts" setup>
import { KillPreview, ReStartPreview } from "../../../wailsjs/go/backend/App";
import { usePostsStore } from "../../store/modules/postsStore/postsStore";

const killPreview = async () => {
    await KillPreview()
}
</script>

<style lang="less" scoped></style>