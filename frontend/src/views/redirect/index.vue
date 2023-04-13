<template>
  <div style="display: flex;justify-content: center;align-items: center;height: 100%;">
    <n-empty description="你什么也找不到">
      <template #extra>
        <n-button size="small" @click="goHome">看看别的</n-button>
      </template>
    </n-empty>
  </div>
</template>
<script lang="ts" setup>
import { useRouter } from 'vue-router'

const router = useRouter()
const goHome = () => {
  router.replace({ path: '/' })
}
</script>
