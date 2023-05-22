<template>
    <span :data-filepath="pathname" :class="className" class="file-icon"></span>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import fileIcons from "../../assets/icons/fileIcons/index";
const props = defineProps({
    name: {
        type: String,
        required: true,
        default: 'mock.md'
    },
    pathname: {
        type: String,
        required: true,
        default: undefined
    }
})
const className = computed(() => {
    let classNames = fileIcons.getClassByName(props.name ? props.name : 'mock.md')

    if (!classNames) {
        // Use fallback icon when the icon is unknown.
        classNames = fileIcons.getClassByName('mock.md')
    }
    return classNames?.split(/\s/)
})
</script>

<style lang="less" scoped>
.file-icon {
    flex-shrink: 0;
    margin-right: 5px;
}
</style>
