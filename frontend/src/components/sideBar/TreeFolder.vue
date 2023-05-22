<template>
    <div class="side-bar-folder" :data-folderpath="folder.pathname">
        <div class="folder-name" @click="folderNameClick" :data-folderpath="folder.pathname"
            :style="{ 'padding-left': `${(depth * 20) + 20}px` }" :class="[{ 'active': folder.id === activeItem.id }]"
            :title="folder.pathname" ref="folder">
            <svg :data-folderpath="folder.pathname" class="icon" aria-hidden="true">
                <use :data-folderpath="folder.pathname"
                    :xlink:href="`#${folder.isCollapsed ? 'icon-folder-close' : 'icon-folder-open'}`"></use>
            </svg>
            <input type="text" @click.stop="noop" class="rename" v-if="renameCache === folder.pathname" v-model="newName"
                ref="renameInput" @keydown.enter="rename">
            <span :data-folderpath="folder.pathname" v-else class="text-overflow">{{ folder.name }}</span>
        </div>
        <div class="folder-contents" v-if="!folder.isCollapsed">
            <tree-folder v-for="(childFolder, index) of folder.folders" :key="index + 'folder'" :folder="childFolder"
                :depth="depth + 1" />
            <input type="text" v-if="createCache.dirname === folder.pathname" class="new-input"
                :style="{ 'margin-left': `${depth * 5 + 15}px` }" ref="input" @keydown.enter="handleInputEnter"
                v-model="createName">
            <tree-file v-for="(file, index) of folder.files" :key="index + 'file'" :file="file" :depth="depth + 1" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { nextTick, reactive, ref } from 'vue';
import { IProjectTree } from '../../../types/interface';
import TreeFile from "./TreeFile.vue";

const props = defineProps({
    depth: {
        type: Number,
        default: 0
    },
    folder: {
        type: Object as () => IProjectTree,
        default: null
    }
})
const createName = ref('')
const newName = ref('')

// 全局缓存变量
const createCache = reactive({
    dirname: ""
})
const renameCache = ref("")
const activeItem = reactive({
    id: ""
})

const handleInputEnter = () => {
    console.log(createName.value);
}

function folderNameClick() {
    props.folder.isCollapsed = !props.folder.isCollapsed
}
function noop() { }
const renameInput = ref<HTMLInputElement>()
function focusRenameInput() {
    nextTick(() => {
        if (renameInput.value) {
            renameInput.value.focus()
            newName.value = props.folder.name
        }
    })
}
function rename() {
    if (newName.value) {
        console.log(newName.value);
    }
}

</script>

<style lang="less" scoped>
.icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
}

.side-bar-folder {
    &>.folder-name {
        cursor: default;
        user-select: none;
        display: flex;
        align-items: center;
        height: 30px;
        padding-right: 15px;

        &>svg {
            flex-shrink: 0;
            color: var(--icon-primary);
            margin-right: 5px;
        }

        &:hover {
            background: var(--scrollbar-color);
        }
    }
}

.new-input,
input.rename {
    outline: none;
    height: 22px;
    margin: 5px 0;
    padding: 0 6px;
    color: var(--sideBarColor);
    border: 1px solid var(--floatBorderColor);
    background: var(--floatBorderColor);
    width: 70%;
    border-radius: 3px;
}
</style>