<template>
    <!-- Project tree view -->
    <div class="project-tree" v-if="projectTree">
        <div class="title">
            <svg class="icon icon-arrow" :class="{ 'fold': !showDirectories }" aria-hidden="true"
                @click.stop="toggleDirectories()">
                <use xlink:href="#icon-arrow"></use>
            </svg>
            <span class="default-cursor text-overflow" @click.stop="toggleDirectories()">{{ projectTree.name }}</span>
        </div>
        <div class="tree-wrapper" v-show="showDirectories" :data-root="true" :data-rootpath="projectTree.pathname">
            <tree-folder v-for="(folder, index) of projectTree.folders" :key="index + 'folder'" :folder="folder"
                :depth="depth" />
            <input type="text" class="new-input" v-show="createCache.dirname === projectTree.pathname"
                :style="{ 'margin-left': `${depth * 5 + 15}px` }" ref="input" v-model="createName"
                @keydown.enter="handleInputEnter">
            <tree-file v-for="(file, index) of projectTree.files" :key="index + 'file'" :file="file" :depth="depth" />
            <div class="empty-project" v-if="projectTree.files.length === 0 && projectTree.folders.length === 0">
                <span>Empty project</span>
                <a href="javascript:;" @click.stop="createFile">Create File</a>
            </div>
        </div>
    </div>
    <div v-else class="open-project">
        <div class="centered-group">
            <folder-icon />
            <button class="button-primary" @click="">
                Open Folder
            </button>
        </div>

    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, reactive } from "vue";
import { IProjectTree } from "../../../types/interface";
import FolderIcon from "../../assets/icons/undraw_folder.svg?component";
import TreeFolder from "./TreeFolder.vue";
import TreeFile from "./TreeFile.vue";

const props = defineProps({
    projectTree: {
        type: Object as () => IProjectTree,
        default: null
    }
})

const showDirectories = ref(true)
const depth = ref(0)
const createName = ref("")
const createCache = reactive({
    dirname: ""
})

onMounted(() => {
    console.log(props.projectTree);
})

const createFile = () => {

}

const toggleDirectories = () => {
    showDirectories.value = !showDirectories.value
}

const handleInputEnter = () => {
    console.log(createName.value);
}

</script>


<style lang="less" scoped>
.list-item {
    display: inline-block;
    margin-right: 10px;
}

.list-enter-active,
.list-leave-active {
    transition: all .2s;
}

.list-enter,
.list-leave-to

/* .list-leave-active for below version 2.1.8 */
    {
    opacity: 0;
    transform: translateX(-50px);
}

.tree-view {
    font-size: 14px;
    color: var(--icon-primary);
    display: flex;
    flex-direction: column;
    height: 100%;
}

.tree-view>.title {
    height: 35px;
    line-height: 35px;
    padding: 0 15px;
    display: flex;
    flex-shrink: 0;
    flex-direction: row-reverse;
}

.icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
}

.icon-arrow {
    margin-right: 5px;
    transition: all .25s ease-out;
    transform: rotate(90deg);
    fill: var(--icon-primary);
}

.icon-arrow.fold {
    transform: rotate(0);
}

.opened-files,
.project-tree {
    &>.title {
        height: 30px;
        line-height: 30px;
        font-size: 14px;
    }
}

.project-tree {
    display: flex;
    flex-direction: column;
    overflow: auto;

    &>.title {
        padding-right: 15px;
        display: flex;
        align-items: center;

        &>span {
            flex: 1;
            user-select: none;
        }

        &>a {
            pointer-events: auto;
            cursor: pointer;
            margin-left: 8px;
            color: var(--icon-primary);
            opacity: 0;
        }

        &>a:hover {
            color: var(--scrollbar-color);
        }

        &>a.active {
            color: var(--scrollbar-color);
        }
    }

    &>.tree-wrapper {
        overflow: auto;
        flex: 1;

        &::-webkit-scrollbar:vertical {
            width: 8px;
        }
    }

    flex: 1;
}

.project-tree div.title:hover>a {
    opacity: 1;
}

.open-project {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    padding-bottom: 100px;

    & .centered-group {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    & svg {
        width: 120px;
        fill: var(--theme-color);
    }

    & button.button-primary {
        display: block;
        margin-top: 20px;
    }
}
</style>