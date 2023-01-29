import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";

function pathResolve(dir: string) {
  return resolve(process.cwd(), ".", dir);
}
// https://vitejs.dev/config/
export default defineConfig({
  base: "/",
  // 路径重定向
  resolve: {
    alias: [
      // {
      //   find: /\/#\//,
      //   replacement: pathResolve("types"),
      // },
      {
        find: "@",
        replacement: pathResolve("src"),
      },
      //   {
      //     find: "vue-i18n",
      //     replacement: "vue-i18n/dist/vue-i18n.cjs.js", //解决i8n警告
      //   },
    ],
    dedupe: ["vue"],
  },
  plugins: [vue()],
  css: {
    preprocessorOptions: {
      less: {
        math: "always",
      },
    },
  },
});
