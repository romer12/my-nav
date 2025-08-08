<template>
    <n-loading-bar-provider>
        <n-dialog-provider>
            <n-notification-provider>
                <n-message-provider>
                    <slot />
                    <NaiveProviderContent />
                </n-message-provider>
            </n-notification-provider>
        </n-dialog-provider>
    </n-loading-bar-provider>
</template>

<script lang="ts" setup>
import { defineComponent, h } from 'vue'
import {
    NLoadingBarProvider,
    NDialogProvider,
    NNotificationProvider,
    NMessageProvider,
    useLoadingBar,
    useDialog,
    useNotification,
    useMessage
} from 'naive-ui'

// 挂载函数到window对象上
const registerNaiveTools = () => {
    window.$loadingBar = useLoadingBar()
    window.$dialog = useDialog()
    window.$notification = useNotification()
    window.$message = useMessage()
}

// 初始化各个组件
const NaiveProviderContent = defineComponent({
    name: 'NaiveProviderContent',
    setup() {
        registerNaiveTools()
    },
    render() {
        return h('div')
    }
})

</script>

<style scoped>

</style>