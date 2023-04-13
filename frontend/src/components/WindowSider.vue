<template>
    <div class="side-bar">
        <n-button text
            style="font-size: 26px;--n-text-color-hover: #ffdb00;--n-text-color-pressed:#ffdb00;--n-text-color-focus: #ffdb00;"
            data-route-path="/home" @click="goTo" :focusable="false">
            <n-icon>
                <article-icon />
            </n-icon>
        </n-button>
        <n-button text
            style="font-size: 24px;--n-text-color-hover: #ffdb00;--n-text-color-pressed:#ffdb00;--n-text-color-focus: #ffdb00;"
            data-route-path="/editor" @click="goTo" :focusable="false">
            <n-icon>
                <documents-icon />
            </n-icon>
        </n-button>
    </div>
    <n-layout has-sider>
        <n-space :wrap="false" :wrap-item="false" style="width: 100%;">
            <n-layout-sider bordered collapse-mode="width" :collapsed-width="0" content-style="overflow: hidden;padding: "
                show-trigger="bar">
                {{ $route.fullPath }}
            </n-layout-sider>
            <n-layout-content content-style="">
                <app-provider>
                    <router-view></router-view>
                </app-provider>
            </n-layout-content>
        </n-space>
    </n-layout>
</template>


<script lang="ts" setup>
import { AppProvider } from "../components/AppProvider/index";
import { DocumentsOutline as DocumentsIcon, } from "@vicons/ionicons5";
import { DocumentOnePage20Regular as ArticleIcon } from "@vicons/fluent";
import { useRouter } from "vue-router";
const router = useRouter()

const goTo = ($event: PointerEvent) => {
    const target = $event.currentTarget as HTMLButtonElement;
    router.push(target.dataset.routePath || "/")

}
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
</style>