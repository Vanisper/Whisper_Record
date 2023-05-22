<template>
    <div v-if="file.isMarkdown" :title="file.pathname" :data-filepath="file.pathname" class="side-bar-file"
        :style="{ 'padding-left': `${(depth * 20) + 20}px`, 'opacity': file.isMarkdown ? 1 : 0.75 }"
        @click="handleFileClick()"
        :class="[{ 'current': usePostsStore().currentFile?.pathname === file.pathname, 'active': file.id === activeItem.id }]">
        <file-icon :pathname="file.pathname" :name="file.name" />
        <input type="text" @click.stop="noop" class="rename" v-if="renameCache === file.pathname" v-model="newName"
            ref="renameInput" @keydown.enter="rename">
        <span :data-filepath="file.pathname" v-else>{{ file.name }}</span>
    </div>
</template>

<script lang="ts" setup>
import { nextTick, ref, reactive } from 'vue';
import { IProjectFile } from '../../../types/interface';
import { isSamePathSync } from '../../utils/file';
import FileIcon from './Icon.vue'
import { usePostsStore } from '../../store/modules/postsStore/postsStore';

const props = defineProps({
    file: {
        type: Object as () => IProjectFile,
        default: null
    },
    depth: {
        type: Number,
        default: 0
    }
})
const newName = ref("")
const renameInput = ref<HTMLInputElement>()
// 缓存数据
const renameCache = ref("")
const activeItem = reactive({
    id: ""
})
const tabs = reactive<IProjectFile[]>([
    {
        birthTime: "",
        id: "",
        isDirectory: "",
        isFile: "",
        isMarkdown: "",
        name: "",
        pathname: "",
    }
])

function noop() { }
function focusRenameInput() {
    nextTick(() => {
        if (renameInput.value) {
            renameInput.value.focus()
            newName.value = props.file.name
        }
    })
}
function rename() {
    console.log(newName.value);
}
async function handleFileClick() {
    if (usePostsStore().isChange) {
        window.$message.warning("请先保存当前文档！")
        return
    }
    const { isMarkdown, pathname } = props.file
    if (!isMarkdown) return
    const openedTab = tabs.find(file => isSamePathSync(file.pathname, pathname))
    if (openedTab) {
        if (usePostsStore().currentFile === openedTab) {
            return
        }
        console.log('UPDATE_CURRENT_FILE', openedTab);
    } else {
        usePostsStore().currentFile = props.file
        window.view.updateContent(await usePostsStore().getCurrentFileContent())
    }
}
</script>

<style lang="less" scoped>
.side-bar-file {
    display: flex;
    position: relative;
    align-items: center;
    cursor: default;
    user-select: none;
    height: 30px;
    box-sizing: border-box;
    padding-right: 15px;

    &:hover {
        background: var(--scrollbar-color);
    }

    &>span {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    &::before {
        content: '';
        position: absolute;
        display: block;
        left: 0;
        background: var(--theme-color);
        width: 2px;
        height: 0;
        top: 50%;
        transform: translateY(-50%);
        transition: all .2s ease;
    }
}

.side-bar-file.current::before {
    height: 100%;
}

.side-bar-file.current>span {
    color: var(--theme-color);
}

.side-bar-file.active>span {
    color: var(--sideBarTitleColor);
}

input.rename {
    height: 22px;
    outline: none;
    margin: 5px 0;
    padding: 0 8px;
    color: var(--sideBarColor);
    border: 1px solid var(--floatBorderColor);
    background: var(--floatBorderColor);
    width: 100%;
    border-radius: 3px;
}
</style>
