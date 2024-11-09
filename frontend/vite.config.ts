import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    port: 3000,
    // fileの変更の監視（dockerを使うため）
    watch: {
      usePolling: true,
    }
  },
  resolve: {
    alias: [{ find: '@', replacement: '/src' }]
  },
})
