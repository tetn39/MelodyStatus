import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from "node:path"

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
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
})
