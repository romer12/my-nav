<template>
  <div class="container">
    <div class="header">
      <div class="header-title">
        <strong class="site-title">我的导航</strong>
        <!-- <div class="current-item" v-show="headerTabShow" @click="toTop">
          <strong v-if="groupData.length">
            {{ groupData[currentIndex].name }}
            <span class="group-count">（{{ groupData[currentIndex].links ? groupData[currentIndex].links.length : 0 }}）</span>
          </strong>
        </div> -->
        <div class="header-middle">
          <n-input class="search-input" placeholder="输入关键词" v-model="searchText" @input="handleSearchChange" clearable></n-input>
        </div>
        <div class="header-right">
          <n-switch
            v-model:value="dark"
            size="large"
            class="mr-20"
            @update:value="handleThemeChange"
          >            
            <template #unchecked-icon>
              <n-icon :component="Sunny" />
            </template>
            <template #checked-icon>
              <n-icon :component="Moon" />
            </template>
          </n-switch>
        </div>
      </div>
      <div class="header-operation">
        <n-button class="mr-20" type="primary" @click="loadGroupData">
          <template #icon>
            <Refresh />
          </template>
          刷新
        </n-button>
        <n-button class="mr-20" type="primary" @click="handleImport()">
          <template #icon>
            <CloudUploadOutline />
          </template>
          导入json
        </n-button>
        <n-button class="mr-20" type="primary" @click="handleExport">
          <template #icon>
            <DownloadOutline />
          </template>
          导出json
        </n-button>
        <n-button type="primary" @click="handleCreateGroup">
          <template #icon>
            <Add />
          </template>
          添加分组
        </n-button>
      </div>
      </div>

    <!--分组列表-->
    <div class="group-list media-group-class mb-20" v-show="mainBoxShow">
      <div
        class="group-list-item"        
        v-if="groupData"
        v-for="(item,index) in groupData"
        :class="currentIndex === index ? 'group-list-active' : ''"
        @click="switchGroupItem(index)"
      >
        <div class="item-title">
          <strong>{{ item.name }} <span class="group-count">（{{ item.links ? item.links.length : 0 }}）</span></strong>
        </div>
        <div class="item-bottom">
          <n-popover trigger="hover">
            <template #trigger>
              <n-button size="tiny" secondary  circle @click="handleCreateLink(item)">
                <template #icon>
                  <n-icon class="op-icon"><Add /></n-icon>
                </template>
              </n-button>
            </template>
            <span>添加链接到该分组</span>
          </n-popover>

          <n-popover trigger="hover">
            <template #trigger>
              <n-button size="tiny" secondary  circle @click="handleUpdateGroup(item)">
                <template #icon>
                  <n-icon class="op-icon"><CreateOutline /></n-icon>
                </template>
              </n-button>
            </template>
            <span>编辑分组</span>
          </n-popover>

          <n-popover trigger="hover">
            <template #trigger>
              <n-button size="tiny" secondary  circle @click="handleGroupImport(item)">
                <template #icon>
                  <n-icon class="op-icon"><CloudUploadOutline /></n-icon>
                </template>
              </n-button>
            </template>
            <span>导入json</span>
          </n-popover>

          <n-popover trigger="hover">
            <template #trigger>
              <n-button size="tiny" secondary  circle type="error" @click="handleDeleteGroup(item)">
                <template #icon>
                  <n-icon><TrashOutline /></n-icon>
                </template>
              </n-button>
            </template>
            <span>删除分组</span>
          </n-popover>
        </div>
      </div>
    </div>

    <div class="main-box">
      <!--链接列表-->
      <div class="link-list mb-20" v-if="mainBoxShow">
        <div class="media-link-class" v-if="linksData && linksData.length">
          <div class="link-list-item" v-for="item in linksData">
            <span :title="item.url" class="item-link" @click="openUrl(item.url)">{{ item.title }}</span>
            <div class="link-operation">
              <n-popover trigger="hover">
                <template #trigger>
                  <n-button size="tiny" secondary  circle @click="handleUpdateLink(item)">
                    <template #icon>
                      <n-icon class="op-icon"><CreateOutline /></n-icon>
                    </template>
                  </n-button>
                </template>
                <span>编辑链接</span>
              </n-popover>
              <n-popconfirm
                negative-text="取消"
                positive-text="确定"
                @positive-click="deleteLink"
              >
                <template #trigger>
                  <n-button style="margin-left: 10px;" size="tiny" secondary  circle type="error" @click="handleDeleteLink(item)">
                    <template #icon>
                      <n-icon><TrashOutline /></n-icon>
                    </template>
                  </n-button>
                </template>
                确定删除此链接？
              </n-popconfirm>
            </div>
          </div>
        </div>
        
        <n-empty v-else description="好像还没有东西哎"></n-empty>
      </div>

      <!--搜索结果-->
      <div class="search-result mb-20" v-else style="display: flex;flex-direction: column;">
        <div class="search-result-header">
          搜索<n-highlight
          :text="searchText"
          :patterns="patterns"
          :highlight-style="{
            padding: '0 6px',
            borderRadius: themeVars.borderRadius,
            display: 'inline-block',
            color: themeVars.baseColor,
            background: themeVars.primaryColor,
            transition: `all .3s ${themeVars.cubicBezierEaseInOut}`,
          }"
          />的结果如下
        </div>
        <div class="link-list" v-if="searchListData && searchListData.length">
          <div class="media-link-class">
            <div class="link-list-item" v-for="item in searchListData">
              <span :title="item.url" class="item-link" @click="openUrl(item.url)">{{ item.title }}</span>
            </div>
          </div>
        </div>
        <n-empty v-else description="好像还没有东西哎"></n-empty>
      </div>

    </div>
        

    <!--添加、修改分组弹窗-->
    <modal-alert
      v-model:show="showGroupModal"
      :title="groupDialogStatus == 'create' ? groupTextMap[0] : groupTextMap[1]"
      confirm-text="确定"
      cancel-text="关闭"
      @cancel="closeAddGroupModal"
      @confirm="groupSubmit"
    >
      <n-form
        :model="groupForm"
        ref="groupFormRef"
        :rules="groupRules"
        label-width="80"
      >
        <n-form-item label="分组名" path="name">
          <n-input v-model:value="groupForm.name" placeholder="请输入1-10个字的分组名" />
        </n-form-item>
      </n-form>
    </modal-alert>

    <!--删除分组弹窗-->
    <modal-alert
      v-model:show="showDelGroupModal"
      title="警告"
      confirm-text="确定"
      cancel-text="再想想"
      @cancel="showDelGroupModal = false"
      @confirm="delGroupSubmit"
    >
      <p style="color: red;">注意！删除分组同时会清空该分组下的所有链接，确定删除？</p>
    </modal-alert>

    <!--添加、修改链接弹窗-->
    <modal-alert
      v-model:show="linkDialogVisible"
      :title="groupDialogStatus == 'create' ? linkTextMap[0] : linkTextMap[1]"
      confirm-text="确定"
      cancel-text="关闭"
      @cancel="linkDialogVisible = false"
      @confirm="linkSubmit"
    >
      <n-form
        :model="linkForm"
        ref="linkFormRef"
        :rules="linkRules"
        label-width="80"
      >
        <n-form-item label="链接名" path="title">
          <n-input v-model:value="linkForm.title" placeholder="请输入链接名" />
        </n-form-item>
        <n-form-item label="链接地址" path="url">
          <n-input v-model:value="linkForm.url" placeholder="请输入链接地址" />
        </n-form-item>
        <n-form-item label="所属分组" path="group_id">
          <n-select
            v-model:value="linkForm.group_id"
            label-field="name"
            value-field="id"
            clearable
            :options="groupData"/>
        </n-form-item>
      </n-form>
    </modal-alert>

    <!--导入弹窗-->
    <modal-alert
      v-model:show="jsonDialogVisible"
      title="上传json文件并导入"
      confirm-text="确定"
      cancel-text="关闭"
      @cancel="jsonDialogVisible = false"
      @confirm="jsonImportSubmit"
      style="width: 600px;"
    >
      <n-tabs type="line" default-value="google" @update:value="tabSelect" v-model:value="jsonHandleType" animated>
        <!--google标签json导入-->
        <n-tab-pane name="google" tab="google收藏夹json导入">
          <div class="json-example">
            <p>json格式示例：</p>
            <div style="overflow-x: auto;"><pre v-if="jsonHandleType == 'google'">{{ jsonFileExample.google }}</pre></div>
          </div>
        </n-tab-pane>
        <!--普通json导入-->
        <n-tab-pane name="group" tab="导入链接到分组">
          <div class="json-example">
            <p>json格式示例：</p>
            <div style="overflow-x: auto;"><pre v-if="jsonHandleType == 'group'">{{ jsonFileExample.group }}</pre></div>
          </div>
        </n-tab-pane>
        <n-tab-pane name="normal" tab="导入分组和链接数据">
          <div class="json-example">
            <p>json格式示例：</p>
            <div style="overflow-x: auto;"><pre v-if="jsonHandleType == 'normal'">{{ jsonFileExample.normal }}</pre></div>
          </div>
        </n-tab-pane>
      </n-tabs>
      <div class="request-error" v-show="uploadError"><p style="color: red;">{{ uploadError }}</p></div>
      <n-upload
        ref="uploadRef"
        :default-upload="false"
        :multiple="false"
        :file-list="fileList"
        :max="1"
        accept=".json"
        @change="handleUploadChange"
      >
        <n-button>选择文件</n-button>
      </n-upload>
    </modal-alert>
    
  </div>

  <BackToTop />
</template>

<script setup lang="ts">
import {
  NButton,
  NEmpty,
  NSelect,
  NForm,
  NFormItem,
  NInput,
  NPopover,
  NPopconfirm,
  NIcon,
  NUpload,
  NSwitch,
  NHighlight,
  NTabs,
  NTabPane,
  useThemeVars,
  type FormInst,
  type FormItemRule,
  type FormRules,
  type UploadFileInfo,
  type UploadInst
} from 'naive-ui'

import { 
  Add,
  Refresh,
  CloudUploadOutline,
  CreateOutline,
  TrashOutline,
  Sunny,
  Moon,
  DownloadOutline
} from '@vicons/ionicons5'
import ModalAlert from '@/components/common/ModalAlert.vue'
import { ref, onMounted, onUnmounted, computed } from 'vue'
import type { Group, Links } from '@/types/data'
import fetchRequest from '@/utils/api'
import { useThemeStore } from '@/stores/theme'
import BackToTop from '@/components/common/BackToTop.vue'

const apiUrl = import.meta.env.VITE_BASE_API

interface LinkItem {
  title: string
  url: string
  id?: number
  group_id?: number
}

interface ExportItem {
  name: string
  links?: LinkItem[]
}

const headerTabShow = ref(false)
const searchText = ref('') // 搜索关键词
const mainBoxShow = computed(() => {
  return searchText.value ? false : true
})
const cacheListData = ref<Links[]>([])
const searchListData = ref<Links[]>([])
const themeVars = useThemeVars()
const patterns = ref<string[]>([]) // 高亮的文本

/** 分组操作 */
const currentIndex = ref<number>(0)
const groupData = ref<Group[]>([])
const showGroupModal = ref(false)
const showDelGroupModal = ref(false)
const groupFormRef = ref<FormInst | null>(null)
const groupDialogStatus = ref('create')
const groupTextMap =  ['添加分组','修改分组']
const groupForm = ref({
  id: null,
  name: '',
})
const groupRules: FormRules = {
  name: [
    {
      validator(rule: FormItemRule, value: string) {
        if (!value) {
          return new Error('请输入分组名')
        }
        else if (value.length < 1) {
          return new Error('分组名在1-10个字之间')
        }
        else if (value.length > 10) {
          return new Error('分组名在1-10个字之间')
        }
        return true
      },
    }
  ]
}

const resetForm = () => {
  groupForm.value = {
    id: null,
    name: '',
  }

  linkForm.value = {
    id: null,
    title: '',
    url: '',
    group_id: null
  }
}

const handleCreateGroup = () => {
  resetForm()
  groupDialogStatus.value = 'create'
  showGroupModal.value = !showGroupModal.value
}

const handleUpdateGroup = (item: any) => {
  resetForm()
  groupForm.value.id = item.id
  groupForm.value.name = item.name
  showGroupModal.value = !showGroupModal.value
  groupDialogStatus.value = 'update'
}

const handleDeleteGroup = (item: any) => {
  resetForm()
  groupForm.value.id = item.id
  groupForm.value.name = item.name
  showDelGroupModal.value = !showDelGroupModal.value
}

const closeAddGroupModal = () => {
  showGroupModal.value = false
}
const groupSubmit = async(e: Event) => {
  e.preventDefault()
  groupFormRef.value?.validate(async (errors) => {
    if (!errors) {
      // 发起请求
      switch (groupDialogStatus.value) {
        case 'create':
          await createOrUpdateGroup(groupForm.value, 'create')
          break;  
        case 'update':
          await createOrUpdateGroup(groupForm.value, 'update') 
          break;
      }
    }
    else {
      showGroupModal.value = true
      console.log(errors[0][0].message)
      $message.error('' + errors[0][0].message)
      return false
    }
    return false
  })
}

const createOrUpdateGroup = async(formData: any, type: string) => {
  $loadingBar.start()

  let url = ''
  switch (type) {
    case 'create':
      url = `${apiUrl}/groups`
      break;  
    case 'update':
      url = `${apiUrl}/group/${formData.id}`
      break;
  }
  try {
    const result = await fetchRequest<any>({
      url: url, 
      method: 'POST', 
      body: { name: formData.name }
    })
    showGroupModal.value = false
    $message.success(result.msg)

    await resetForm()
    await loadGroupData()    
  } catch (err: any) {    
    $message.error('请求出错: ' + err.message)
  } finally {
    $loadingBar.finish()
  }
}

const delGroupSubmit = async() => {
  if(!groupForm.value.id || groupForm.value.id <= 0){
    $message.error('id参数错误')
    return false
  }
  $loadingBar.start()

  try {
    const result = await fetchRequest<any>({
      url: `${apiUrl}/groups/${groupForm.value.id}`,
      method: 'DELETE'
    });
    if (groupData) {
      if(currentIndex.value == groupData.value.length){
        if(currentIndex.value - 1 > 0){
          currentIndex.value--
        }        
      }
    }

    showDelGroupModal.value = false
    await loadGroupData()
    
    $message.success(result.msg)
  } catch (err: any) {    
    $message.error('删除分组出错: ' + err.message)
    return false
  } finally{
    $loadingBar.finish()
  }  
}

const switchGroupItem = (index: number) => {
  currentIndex.value = index
  linksData.value = []
  linksData.value = groupData.value[currentIndex.value].links
}

/** 链接 */
const linksData = ref<Links[]>([])
const linkDialogVisible = ref(false)
const linkFormRef = ref<FormInst | null>(null)
const linkForm = ref({
  id: null,
  title: '',
  url: '',
  group_id: null
})
const linkTextMap =  ["添加链接","编辑链接"]
const linkRules: FormRules = {
  title: [{required: true, message: '请输入链接名', trigger: ['input','blur']}],
  url: [{required: true, message: '请输入链接地址', trigger: ['input','blur']}],
  group_id: [
    {
      validator(rule: FormItemRule, value: number) {
        if (!value) {
          return new Error('请选择所属分组')
        }
        else if (value <= 0) {
          return new Error('分组id参数错误')
        }
        return true
      },
    }
  ],
}

const openUrl = (url: string) => {
  // 检测字符串中是否含有 http 或 https
  if (url.startsWith('http://') || url.startsWith('https://')) {
    // 如果包含，打开新标签页并跳转到该地址
    window.open(url, '_blank');
  } else {
    // 如果不包含，添加 http:// 并打开新标签页
    window.open('http://' + url, '_blank');
  }
}

const handleCreateLink = (item: any) => {
  resetForm()
  linkDialogVisible.value = !linkDialogVisible.value
  groupDialogStatus.value = 'create'
  linkForm.value.group_id = item.id
}

const handleUpdateLink = (item: any) => {
  resetForm()
  Object.assign(linkForm.value, item)
  linkDialogVisible.value = !linkDialogVisible.value
  groupDialogStatus.value = 'update'
}

const linkSubmit = async(e: Event) => {
  // 表单验证
  e.preventDefault()
  linkFormRef.value?.validate(async (errors) => {
    if (!errors) {
      // 发起请求
      switch (groupDialogStatus.value) {
        case 'create':
          await createOrUpdateLink(linkForm.value, 'create')
          break;  
        case 'update':          
          await createOrUpdateLink(linkForm.value, 'update') 
          break;
      }
    }
    else {
      linkDialogVisible.value = true
      console.log(errors[0][0].message)
      $message.error('' + errors[0][0].message)
      return false
    }
    return false
  })
}

const createOrUpdateLink = async(formData: any, type: string) => {
  let url = ''
  switch (type) {
    case 'create':
      url = `${apiUrl}/links`
      formData.id = null
      break;  
    case 'update':
      url = `${apiUrl}/link`
      break;
  }
  try {
    const result = await fetchRequest<any>({
      url: url, 
      method: 'POST', 
      body: { 
        id: formData.id,
        title: formData.title,
        url: formData.url,
        group_id: formData.group_id
      }
    })
    linkDialogVisible.value = false
    $message.success(result.msg)
    await resetForm()
    await loadGroupData()
  } catch (err: any) {    
    $message.error('请求出错: ' + err.message)
    return false
  }
}

const handleDeleteLink = (item: any) => {
  resetForm()
  Object.assign(linkForm.value,item)
}
const deleteLink = async () => {
  $loadingBar.start()
  if(!linkForm.value.id) {
    $message.error('id参数错误')
    return false
  }

  try {
    const result = await fetchRequest<any>({
      url: `${apiUrl}/links/${linkForm.value.id}`, 
      method: 'DELETE'
    })        
    $message.success(result.msg)
    resetForm()
    await loadGroupData()
  } catch (err: any) {
    $message.error('请求出错: ' + err.message)
    return false
  } finally {
    $loadingBar.finish()
  }
}

const debounce = (func:Function, delay: number) => {
  let timeout:number;
  return (...args:any) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => {
      func(...args);
    }, delay);
  }
}

const processInput = (text:string) => {
  if (!text) {
    searchText.value = ''
    patterns.value = []
    return false
  }
  searchText.value = text.toLowerCase()
  patterns.value.push(searchText.value)
  searchListData.value = []
  // 在缓存的数据中查找
  if (cacheListData.value) {
    cacheListData.value.forEach((item,index) => {
      if (item.title.toLowerCase().includes(text)) {
        searchListData.value.push(item)
      }
    })
  }
}

const handleSearchChange = debounce((text:string) => {
  processInput(text);
}, 300)

/** 导入json */
const uploadData = ref({})
const uploadError = ref('')
const jsonDialogVisible = ref(false)
const jsonHandleType = ref('google')
const uploadRef = ref<UploadInst | null>(null)
const fileList = ref<UploadFileInfo[]>([])
const jsonFileExample = {
  'google': `  
    该导入仅支持chrome浏览器导出的json文件！    
    windows 系统：
      C:/Users/Administrator/AppData/Local/Google/Chrome/User Data/Default/Bookmarks
    mac 系统：
      /Users/一般是你的电脑账号名/Library/Application Support/Google/Chrome/Default/Bookmarks
  `,
  'group': `
    [
      {
        "title": "淘宝",
        "url": "https://www.taobao.com"
      },
      {
        "title": "百度",
        "url": "https://www.baidu.com"
      }
    ]
  `,
  'normal': `
    [
      {
        name: '分组1'
        links: [
          {
            title: '链接1',
            url: 'http://wwww.example.com'
          },
          {
            title: '链接2',
            url: 'http://wwww.example.com'
          }
        ]
      }
    ]
  `
}

const tabSelect = (value:string) => {
  jsonHandleType.value = value  
}

// 导入chrome书签json文件数据
const handleImport = async() => {
  jsonDialogVisible.value = !jsonDialogVisible.value
  resetFileForm()
  uploadData.value = {
    op_type: jsonHandleType.value,
    group_id: null
  }
}


// 导出数据为json
const handleExport = async() => {
  // 处理数据成上面的格式
  if (!groupData.value || groupData.value.length === 0) {
    $message.error('数据为空')
    return false
  }

  let exportData: ExportItem[] = []
  groupData.value.forEach((item: any) => {
    let linksArr: LinkItem[] = []
    if (item.links && item.links.length > 0) {      
      item.links.forEach((link:any) => {
        linksArr.push({
          title: link.title,
          url: link.url
        })
      })
      
      exportData.push({
        name: item.name,
        links: linksArr
      })
    } else {
      exportData.push({
        name: item.name
      })
    }
  })

  // 将数据转换为 JSON 字符串
  const jsonString = JSON.stringify(exportData, null, 2); // 格式化为可读的 JSON

  // 创建一个 Blob 对象
  const blob = new Blob([jsonString], { type: 'application/json' });

  // 创建一个 URL 对象
  const url = URL.createObjectURL(blob);

  // 创建一个链接元素
  const a = document.createElement('a');
  a.href = url;
  a.download = 'data.json'; // 设置下载文件名

  // 触发下载
  document.body.appendChild(a);
  a.click();

  // 清理
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
}
// 从json文件导入链接数据到某个组
const handleGroupImport = async(item: any) => {
  jsonDialogVisible.value = !jsonDialogVisible.value
  jsonHandleType.value = 'group'
  resetFileForm()
  uploadData.value = {
    op_type: jsonHandleType.value,
    group_id: item.id
  }  
}

const resetFileForm = () => {
  uploadRef.value?.clear()
  fileList.value = []
  uploadError.value = ''
}

const jsonImportSubmit = () => {

  if(!fileList.value.length) {
    $message.error('请先选择文件！')
    return false
  }
  $loadingBar.start()
  jsonDialogVisible.value = false

  fileList.value.forEach(file => {
    uploadFile(file, uploadData.value)
  })
}

const uploadFile = async(file: any,formOtherData: any) => {
  const formData = new FormData();
  formData.append('file', file.file); // 将文件添加到 FormData 中

  let url = `${apiUrl}/import?type=` + formOtherData.op_type + '&group_id='+ formOtherData.group_id

  try {    
    const response = await fetch(url, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      $message.error('请求发生错误')
      return false
    }   

    const data = await response.json(); // 解析 JSON 响应
    if(data.code !== 0) {
      $loadingBar.finish()
      jsonDialogVisible.value = true
      uploadError.value = data.msg
      $message.error(data.msg)
      return false
    }
    jsonDialogVisible.value = false
    $loadingBar.finish()
    // 重置上传表单
    resetFileForm()
    // 重新加载数据
    await loadGroupData()
    $message.success(data.msg)
  } catch (error) {
    // 处理错误
    throw error;
  }
}

const handleUploadChange = (options: { fileList: UploadFileInfo[] }) => {
  console.log('选择发生了改变')
  resetFileForm()
  //fileList.value = []
  // 更新文件列表
  // console.log(options.fileList)
  fileList.value = options.fileList;
}

/** 主题 */
const dark = ref(false)
const themeStore = useThemeStore()
function handleThemeChange() {
  let theme = dark.value ? 'dark' : 'light'
  themeStore.setTheme(theme)
}

// 加载数据
const loadGroupData = async() => {
  $loadingBar.start()

  try {
    groupData.value = []; // 清空之前的数据
    const result = await fetchRequest<any>({
      url: `${apiUrl}/groups`,

    })
    cacheListData.value = []

    groupData.value = result.data    
    if(groupData.value.length){
      if(!groupData.value[currentIndex.value]){
        if(currentIndex.value - 1 > 0){
          currentIndex.value--
        }
      }
      linksData.value = groupData.value[currentIndex.value].links || []
      if (!cacheListData.value || cacheListData.value.length == 0) {
        result.data.forEach((item:any) => {
          if (item.links.length > 0) {
            item.links.forEach((link:any) => {
              cacheListData.value.push(link)
            })
          }
        })    
      }
    }else{
      linksData.value = []
    }
    
    // console.log('数据')
    // console.log(groupData.value[currentIndex.value])
    
  } catch (err: any) {
    console.log(err)
    $message.error(err.message)
    return false
  } finally {
    $loadingBar.finish()
  }
}

const toTop = () => {
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

onMounted(() => {
  loadGroupData()
  window.addEventListener('scroll', ()=>{
    if(window.scrollY > 400){
      headerTabShow.value = true
    }else{
      headerTabShow.value = false
    }
  })
  dark.value = themeStore.currentTheme === 'light' ? false : true
})

// 组件卸载时移除滚动监听
onUnmounted(() => {
  window.removeEventListener('scroll', ()=>{})
})
</script>

<style lang="scss" scoped>
.mb-10{
  margin-bottom: 10px;
}
.mb-20{
  margin-bottom: 20px;
}
.mr-20{
  margin-right: 20px;
}
.container{
  position: relative;
  margin: 10px auto;
  padding: 10px;
  border-radius: 8px;
  width: 70vw;
  display: flex;
  flex-direction: column;
  transition: ease-in-out 1s;
  background-color: white;

  /** 头部 */
  .header{
    padding: 10px 0;
    margin: 10px 0 20px 0;
    width: 100%;    
    background-color: white;
    position: sticky;
    top: 0;
    left: 0;
    z-index: 10;

    .header-title {
      display: flex;
      justify-content: space-between;
      align-items: center;
      .current-item{
        display: flex;
        align-items: center;
        background-color: #29292a;
        color: white;
        border-radius: 8px;
        overflow: hidden;
        padding: 5px 8px;
        font-size: 12px;
        animation: fadeInOut 1s;
        cursor: pointer;
      }
      @keyframes fadeInOut {
        0% {
          opacity: 0; /* 开始时透明 */
        }
        100% {
          opacity: 1; /* 结束时再次透明 */
        }
      }
    }
    /** 操作 */
    .header-operation {
      display: flex;
      align-items: center;
      padding: 10px 0;
      margin-top: 10px;
      flex-wrap: wrap;
      gap: 10px;
    }
  }
  /** 分组列表 */  
  .group-list{
    display: grid;
    gap: 20px;
    padding: 0 5px 20px 5px;
    background-color: white;
    border-bottom: 1px solid #c9c9c9;

    .group-list-item{
      padding: 10px;
      border: 1px solid #ddd;
      border-radius: 8px;
      overflow: hidden;
      cursor: pointer;
      font-size: 12px;
      display: flex;
      flex-direction: column;
      &:hover{
        box-shadow: 0 0 10px rgba(0,0,0,0.1);
      }
      /** 标题和操作按钮 */
      .item-title{
        display: flex;
        margin-bottom: 10px;        
      }
      .item-bottom{
        display: flex;
        justify-content: space-around;
      }
      .group-count{
        color: #666;
      }
    }
    .group-list-active{
      background-color: rgba(24, 160, 88,.1);
      color: #18a058;
      .group-count{
        color: #18a058;
      }
    }
  }

  /** 链接列表 */
  .link-list{
    background-color: white;
    .media-link-class{
      display: grid;
      gap: 20px;
      padding: 0 5px;

      .link-list-item{
        display: flex;
        justify-content: space-between;
        border: 1px solid #ddd;
        border-radius: 8px;
        cursor: pointer;
        font-size: 14px;
        padding: 5px 10px;
        &:hover{
          background-color: #ddd;
          a{
            color: #0077ff;
          }
        }

        .item-link{
          display: flex;
          align-items: center;
          flex: 1;
          color: #066bd3;
          text-decoration: none;
          &:hover{
            color: #0077ff;
          }
        }
        .link-operation{
          display: flex;
          align-items: center;
        }
      }
    }
  }

  /** 搜索结果 */
  .search-result {
    .search-result-header {
      padding-bottom: 20px;
    }
  }
}

.json-example{
  padding: 10px;
  background-color: #29292a;
  color: #ccc;
  border-radius: 8px;
  margin-bottom: 10px;

  p{
    padding: 0;
    margin: 0;
  }
}

.op-icon{
  color: #7a7a7a;
  &:hover{
    color: #424242;
  }
}

/*屏幕简单适配*/
@media (max-width: 600px){
  .media-group-class {
    grid-template-columns: repeat(1,1fr);
  }
  .media-link-class {
    grid-template-columns: repeat(1,1fr);
  }
  .site-title, .current-item, .header-right {
    display: none;
  }
  .header-middle {
    width: 100%;
  }
}
@media (min-width: 601px) and (max-width: 911px){
  .media-group-class {
    grid-template-columns: repeat(3,1fr);
  }
  .media-link-class {
    grid-template-columns: repeat(1,1fr);
  }
  .site-title, .current-item, .header-right {
    display: none;
  }
  .header-middle {
    width: 100%;
  }
}
@media (min-width: 912px) and (max-width: 1200px){
  .media-group-class {
    grid-template-columns: repeat(3,1fr);
  }
  .media-link-class {
    grid-template-columns: repeat(2,1fr);
  }
}
@media (min-width: 1201px) and (max-width: 1500px){
  .media-group-class {
    grid-template-columns: repeat(5,1fr);
  }
  .media-link-class {
    grid-template-columns: repeat(2,1fr);
  }
}
@media (min-width: 1501px){
  .media-group-class {
    grid-template-columns: repeat(7,1fr);
  }
  .media-link-class {
    grid-template-columns: repeat(2,1fr);
  }
}
</style>
