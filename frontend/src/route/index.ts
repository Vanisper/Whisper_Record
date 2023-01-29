import type {App} from "vue";

import {createRouter, createWebHashHistory, RouteRecordRaw} from "vue-router";
import {Layout} from "./constant";
import {createRouterGuards} from "./router-guards";
import {PageEnum} from "../enums/pageEnum";
import {HttpErrorPage, LoginRoute, RedirectRoute, ReloadRoute} from "./base";


const RootRoute: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Root",
        redirect: PageEnum.BASE_HOME,
        component: Layout,
        meta: {
            title: "Root"
        },
        children: [...HttpErrorPage,]
    },
]

export const constantRouter: any[] = [LoginRoute, ...RootRoute, RedirectRoute, ReloadRoute]

const router = createRouter({
    history: createWebHashHistory(),
    routes: constantRouter,
    strict: true
})

export function setupRouter(app: App) {
    app.use(router);
    createRouterGuards(router)
}

export default router