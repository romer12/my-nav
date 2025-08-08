<template>
    <n-modal
        v-model:show="props.show"
        preset="card"
        :title="title"
        style="width: 400px;"
        :on-update:show="handleUpdateShow"
        @close="handleClose('close')"
    >
        <slot />
        <template #footer>
          <div class="footer">
            <n-button class="cancel-btn" @click="handleClose('button')">{{ props.cancelText }}</n-button>
            <n-button type="primary" @click="handleConfirm">{{ props.confirmText }}</n-button>
          </div>
      </template>
    </n-modal>
</template>

<script lang="ts" setup>
import { defineProps } from 'vue'
import {
    NModal,
    NButton
} from 'naive-ui'

interface Props {
  show: boolean
  title?: string
  confirmText?: string
  cancelText?: string
  maskClosable?: boolean // 新增：是否允许点击遮罩层关闭
}

const props = withDefaults(defineProps<Props>(), {
  title: '提示',
  confirmText: '确定',
  cancelText: '取消',
  maskClosable: true // 默认允许点击遮罩层关闭
})

const emit = defineEmits(['update:show', 'confirm', 'cancel'])

// 统一关闭处理方法
const handleClose = (triggerBy?: string) => {
    emit('cancel', triggerBy) // 将触发来源传递给父组件
    emit('update:show', false)  
}

const handleConfirm = (e: Event) => {
  emit('confirm', e)
  emit('update:show', true)
}

// 完成受控模式
// 阻止 Naive UI 内部自动关闭
const handleUpdateShow = (value: boolean) => {
  if (!value) {
    handleClose('mask')
  }
}
</script>

<style scoped>
.footer{
    display: flex;
    justify-content: flex-end;
}
.cancel-btn{
    margin-right: 20px;
}
</style>