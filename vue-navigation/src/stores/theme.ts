import { ref, watch, onMounted, computed } from 'vue'
import { defineStore } from 'pinia'

export const useThemeStore = defineStore('theme', () => {
  const isDark = computed( () => {
    if(currentTheme.value === 'dark'){
      return true
    }else{
      return false
    }
  })
  const currentTheme = ref('')

  const setTheme = (theme: string) => {
    currentTheme.value = theme
    localStorage.setItem('theme',theme)
  }

  const initTheme = () => {
    const themeName = localStorage.getItem('theme')
    if(!themeName){
      localStorage.setItem('theme','light')
      currentTheme.value = 'light'
    }else{
      currentTheme.value = themeName
    }
  }

  watch(()=>isDark.value,(dark)=>{
    console.log('is dark')
    console.log(dark)
    if (dark)
      document.documentElement.classList.add('dark')
    else
      document.documentElement.classList.remove('dark')
  })

  onMounted(()=>{
    if (isDark.value)
      document.documentElement.classList.add('dark')
    else
      document.documentElement.classList.remove('dark')
  })

  return { currentTheme, initTheme, setTheme, isDark}
})
