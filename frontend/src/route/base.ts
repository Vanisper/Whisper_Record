import { RouteRecordRaw } from "vue-router";
import type { AppRouteRecordRaw } from "./types";
import { ErrorPage403, ErrorPage404, ErrorPage500, Layout } from "./constant";

import { PageEnum } from "../enums/pageEnum";
import { Reload } from "../components/Reload/index";

export const EditorRoute: RouteRecordRaw = {
  path: PageEnum.BASE_EDITOR,
  name: PageEnum.BASE_EDITOR_NAME,
  component: () => import("../views/editor/index.vue"),
  meta: {
    title: "编辑器",
  },
};

export const LoginRoute: RouteRecordRaw = {
  path: "/login",
  name: "Login",
  component: () => import("../views/login/index.vue"),
  meta: {
    title: "登录",
  },
};

export const HttpErrorPage: RouteRecordRaw[] = [
  {
    path: "/error/404",
    name: PageEnum.ERROR_PAGE_NAME_404,
    component: ErrorPage404,
    meta: {
      title: PageEnum.ERROR_PAGE_NAME_404,
    },
  },
  {
    path: "/error/403",
    name: PageEnum.ERROR_PAGE_NAME_403,
    component: ErrorPage403,
    meta: {
      title: PageEnum.ERROR_PAGE_NAME_403,
    },
  },
  {
    path: "/error/500",
    name: PageEnum.ERROR_PAGE_NAME_500,
    component: ErrorPage500,
    meta: {
      title: PageEnum.ERROR_PAGE_NAME_500,
    },
  },
];

export const ErrorPageRoute: AppRouteRecordRaw = {
  path: "/:path(.*)*",
  name: "ErrorPage",
  component: ErrorPage404,
  meta: {
    title: PageEnum.ERROR_PAGE_NAME_404,
    hideBreadcrumb: true,
  },
};

export const ReloadRoute: AppRouteRecordRaw = {
  path: PageEnum.RELOAD,
  name: PageEnum.RELOAD_NAME,
  component: Reload,
  meta: {
    title: PageEnum.RELOAD_NAME,
  },
};

export const RedirectRoute: AppRouteRecordRaw = {
  path: PageEnum.REDIRECT,
  name: PageEnum.REDIRECT_NAME,
  component: Layout,
  meta: {
    title: PageEnum.REDIRECT_NAME,
  },
  children: [
    {
      path: "/redirect/:path(.*)",
      name: PageEnum.REDIRECT_NAME,
      component: () => import("@/views/redirect/index.vue"),
      meta: {
        title: PageEnum.REDIRECT_NAME,
        hideBreadcrumb: true,
      },
    },
  ],
};
