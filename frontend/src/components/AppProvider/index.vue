<template>
  <n-loading-bar-provider :to="loadingBarTargetRef as HTMLElement" container-style="position: absolute;">

    <n-dialog-provider>
      <dialog-content></dialog-content>
      <loading-content></loading-content>
      <n-notification-provider>
        <n-message-provider placement='bottom-right'>
          <message-content></message-content>
          <!-- 将loading进度条放置在指定位置(局部使用) -->
          <div ref="loadingBarTargetRef" style="
            position: relative;
            inset: 0;
            border-radius: var(--n-border-radius);
            /* overflow: hidden; */
            pointer-events: none;
          "></div>
          <slot></slot>
        </n-message-provider>
      </n-notification-provider>
    </n-dialog-provider>
  </n-loading-bar-provider>
</template>

<script lang="ts" setup>
import { NDialogProvider, NLoadingBarProvider, NMessageProvider, NNotificationProvider } from 'naive-ui';

import { MessageContent } from '../Plugins/MessageContent/index';
import { DialogContent } from '../Plugins/DialogContent/index';
import { LoadingContent } from '../Plugins/LoadingContent/index';
import { ref } from 'vue';
const loadingBarTargetRef = ref<string | false | HTMLElement>(false)
</script>
