<template>
  <div 
    v-show="show" 
    class="fixed-top"
    @click="scrollToTop"
  >
    <div class="fixed-top-content">
      <n-icon class="icon"><ArrowUp /></n-icon>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ArrowUp } from '@vicons/ionicons5'
import { NIcon } from 'naive-ui'
import { ref, onMounted, onUnmounted } from 'vue'

const show = ref(false)
const scrollThreshold = ref(300) // 显示按钮的滚动阈值

// 监听滚动事件
const handleScroll = () => {
  show.value = window.scrollY > scrollThreshold.value
}

// 滚动到顶部
const scrollToTop = () => {
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

// 组件挂载时添加滚动监听
onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

// 组件卸载时移除滚动监听
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style scoped>
.fixed-top{
  position: fixed;
  right: 8%;
  bottom: 20%;
  z-index: 50;
  cursor: pointer;
  transition: all .3s;
}
.fixed-top:hover{
  transform: translate(0,-2px);
}
.fixed-top .fixed-top-content{
  width: 40px;
  height: 40px;
  background-color: #ffffff;
  border-radius: 50%;
  box-shadow: -2px 5px 10px 0 rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
}
.fixed-top .fixed-top-content:hover{
  background-color: #e5353e;
}
.fixed-top .fixed-top-content:hover .icon{
  color: #ffffff;
}
.fixed-top .fixed-top-content .icon{
  color: #4a5565;
  font-size: 20px;
}
</style>