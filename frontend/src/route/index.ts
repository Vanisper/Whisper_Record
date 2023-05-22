import type { App } from "vue";

import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import { Layout } from "./constant";
import { createRouterGuards } from "./router-guards";
import { PageEnum } from "../enums/pageEnum";
import {
  EditorRoute,
  HttpErrorPage,
  SettingRoute,
  LoginRoute,
  RedirectRoute,
  ReloadRoute,
} from "./base";
import modules from "./modules";

const RootRoute: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Root",
    redirect: PageEnum.BASE_EDITOR,
    component: Layout,
    meta: {
      title: "Root",
      keepAlive: true,
    },
    children: [...HttpErrorPage, modules.homeRouters, EditorRoute],
  },
];

export const constantRouter: any[] = [...RootRoute, SettingRoute, LoginRoute, RedirectRoute, ReloadRoute];

const router = createRouter({
  history: createWebHashHistory(),
  routes: constantRouter,
  strict: true,
});

export function setupRouter(app: App) {
  app.use(router);
  createRouterGuards(router);
}

export default router;
