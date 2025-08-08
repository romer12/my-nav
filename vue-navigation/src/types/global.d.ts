// 全局导入naive ui的四个独立组件，并在NaiveProvider.vue里注册这四个组件，挂载到windows对象上
// 使用方法：window.$message.info('提示信息')
// 也可以简写，因为下面已经定义有常量：$message.info('提示信息')
// 如此一来，这四个组件就可以在项目的页面内随意使用啦，不用useMessage什么的，直接按照上面的用法来用就可以
declare global {
    const $loadingBar: import('naive-ui').LoadingBarApi
    const $dialog: import('naive-ui').DialogApi
    const $message: import('naive-ui').MessageApi
    const $notification: import('naive-ui').NotificationApi

    interface Window {
        $loadingBar: import('naive-ui').LoadingBarApi
        $dialog: import('naive-ui').DialogApi
        $message: import('naive-ui').MessageApi
        $notification: import('naive-ui').NotificationApi
    }
}

export {}