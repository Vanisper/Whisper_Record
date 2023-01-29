<template>
  <n-loading-bar-provider>
    <n-dialog-provider>
      <dialog-content></dialog-content>
      <loading-content></loading-content>
      <n-notification-provider>
        <n-message-provider>
          <message-content></message-content>
          <slot></slot>
        </n-message-provider>
      </n-notification-provider>
    </n-dialog-provider>
  </n-loading-bar-provider>
</template>

<script lang="ts" setup>
import {NDialogProvider, NLoadingBarProvider, NMessageProvider, NNotificationProvider} from 'naive-ui'

import {MessageContent} from '../Plugins/MessageContent/index'
import {DialogContent} from '../Plugins/DialogContent/index'
import {LoadingContent} from '../Plugins/LoadingContent/index'</script>
