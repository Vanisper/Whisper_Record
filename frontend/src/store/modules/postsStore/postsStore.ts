import { defineStore } from "pinia"
import { ref, computed, watch, toRaw, nextTick, onMounted } from "vue"
import { IProjectFile, IProjectTree } from "../../../../types/interface"
import { GetPostsContent, GetPostsList } from "../../../../wailsjs/go/backend/App"

export const usePostsStore = defineStore("usePostsStore", () => {
    const currentFile = ref<IProjectFile>()
    const currentFileOld = ref<IProjectFile | undefined>(currentFile.value)
    const getCurrentFileContent = async () => await GetPostsContent(currentFile.value!.pathname)
    const currentFileTitle = ref<string>()
    const isChange = ref(false)
    const projectTree = ref<IProjectTree>()
    const updateProjectTree = async () => {
        projectTree.value = JSON.parse(await GetPostsList())
    }
    watch(currentFile, (value) => {
        if (value) {
            currentFileTitle.value = value.name
        }
    })

    return {
        isChange,
        currentFile,
        currentFileOld,
        currentFileTitle,
        getCurrentFileContent,
        projectTree,
        updateProjectTree
    }
})