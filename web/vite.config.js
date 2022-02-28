import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        host: '0.0.0.0',
        proxy: {
            '/api': {
                target: 'https://eea.a1ansong.com/',
                changeOrigin: true,
            },
        },
    },
    plugins: [vue()]
})