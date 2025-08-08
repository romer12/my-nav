<template>
  <n-config-provider>
    <app-provider>
      <RouterView />
    </app-provider>
  </n-config-provider>
  
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import { NConfigProvider } from 'naive-ui'
import AppProvider from '@/components/common/AppProvider.vue'
import { useThemeStore } from '@/stores/theme'

const themeStore = useThemeStore()
themeStore.initTheme()
</script>

<style scoped>

</style>
