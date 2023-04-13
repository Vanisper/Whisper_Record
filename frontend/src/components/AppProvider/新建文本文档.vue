<template>
      <n-notification-provider>
        <n-message-provider placement='bottom-right'>
          <message-content></message-content>
          <slot></slot>
        </n-message-provider>
      </n-notification-provider>
</template>

<script lang="ts" setup>
import { NMessageProvider, NNotificationProvider } from 'naive-ui';
import { MessageContent } from '../Plugins/MessageContent/index';

</script>