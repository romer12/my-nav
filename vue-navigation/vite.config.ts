import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig((env)=>{
  const viteEnv = loadEnv(env.mode, process.cwd()) 

  return {
    plugins: [
      vue(),
      // vueDevTools(),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
    },
    // 开发服务器设置
    server: {
      port: viteEnv.VITE_APP_PORT ? parseInt(viteEnv.VITE_APP_PORT) : 9005, // 默认端口为 9005
    },
    // 打包设置
    build: {
      rollupOptions: {
        output: {
          entryFileNames: 'index.js',
          chunkFileNames: '[name].js',
          assetFileNames: (assetInfo) => {
            if (assetInfo.names.some(name => name.endsWith('.css'))) {
              return 'style.css'
            }
            return '[name].[hash][extname]'; // 其他资源的默认命名
          },
        }
      }
    }
  }
})
