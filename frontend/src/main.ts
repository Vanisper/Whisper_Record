import {createApp} from 'vue'
import App from './App.vue'
import {AppProvider} from "./components/AppProvider/index";
import {setupDirectives, setupNaive} from "./plugins";
import router, {setupRouter} from "./route";
import {setHtmlTheme} from "./utils";
import {setupStore} from "./store";

async function appInit() {
    const appProvider = createApp(AppProvider)
    const app = createApp(App)
    // 注册全局常用 naive-ui 组件
    setupNaive(app)
    // 注册全局自定义指令
    setupDirectives(app)

    // 挂载状态管理
    setupStore(app)

    // 解决路由守卫，Axios中可使用，Dialog，Message 等全局组件
    appProvider.mount('#appProvider', true)
    // 挂载路由
    setupRouter(app)
    // 路由准备就绪后挂载APP实例
    await router.isReady()

    // Store 准备就绪后处理主题色
    setHtmlTheme()

    // 挂载到页面
    app.mount('#app', true)

    // 挂载到 window
    window['$vue'] = app
}

void appInit()
