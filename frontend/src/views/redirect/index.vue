<template>
  <n-empty description="你什么也找不到">
    <template #extra>
      <n-button size="small" @click="goHome">看看别的</n-button>
    </template>
  </n-empty>
</template>
<script lang="ts" setup>
import {useRouter} from 'vue-router'

const router = useRouter()
const goHome = () => {
  router.replace({path: '/'})
}
</script>
