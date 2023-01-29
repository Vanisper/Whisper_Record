import {Router} from "vue-router";
import {PageEnum} from "../enums/pageEnum";


export function createRouterGuards(router: Router) {
    router.beforeEach((to, from, next) => {
        const Loading = window["$loading"];
        Loading && Loading.start()
        // 查找是否存在这个路由
        const isErrorPage = router.getRoutes().findIndex((item) => item.name === to.name)
        if (isErrorPage === -1) {
            next({name: PageEnum.ERROR_PAGE_NAME_404})
        }
        // Todo 判断是否登录

        // next
        next()
    })

    router.afterEach((to, _, failure) => {
        const Loading = window.$loading;
        document.title = (to?.meta?.title as string) || document.title
        Loading && Loading.finish()
    })

    // 错误处理
    router.onError((error) => {
        console.log(error, "路由错误")
    })
}