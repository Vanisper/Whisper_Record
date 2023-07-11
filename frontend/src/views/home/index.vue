<template>
    <h1>HOME</h1>
    <div v-if="useDesignStore().darkTheme" v-contextmenu.dark="contextmenus" style="height: 100%;">
        {{ msg }}
    </div>
    <div v-else v-contextmenu="contextmenus" style="height: 100%;">
        {{ msg }}
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { IClickMenuItem } from "web-contextmenu/type/ContextMenuType";
import { useDesignStore } from '../../store/modules/designStore/designStore';

const msg = ref("")
const contextmenus = ref<IClickMenuItem[]>([
    {
        text: '删除',
        subText: 'BACKSPACE',
        action: (el, axis, menus, isDark) => {
            msg.value = '你点击了删除';
            console.log(el, axis, menus, isDark);
        }
    },
    {
        text: '禁用菜单项',
        disable: true,
    },
    { divider: true },
    {
        text: '多级菜单',
        children: [
            { text: '子菜单1' },
            { text: '子菜单2' },
            {
                text: '三级菜单',
                children: [
                    { text: '子菜单1' },
                    { text: '子菜单2' },
                ],
            },
        ],
    },
    { divider: true },
    {
        text: '剪切',
        subText: 'CTRL + X',
        action: () => msg.value = '你点击了剪切'
    },
    {
        text: '复制',
        subText: 'CTRL + C',
        action: () => msg.value = '你点击了复制'
    },
    {
        text: '粘贴',
        subText: 'CTRL + V',
        action: () => msg.value = '你点击了粘贴'
    },
])
</script>