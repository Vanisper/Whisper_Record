import { RouteRecordRaw } from "vue-router";
import { PageEnum } from "../../enums/pageEnum";

const importPath = {
  "PageEnum.BASE_HOME_NAME": () => import("../../views/home/index.vue"),
};

const homeRouters: RouteRecordRaw = {
  path: PageEnum.BASE_HOME,
  name: PageEnum.BASE_HOME_NAME,
  component: importPath["PageEnum.BASE_HOME_NAME"],
  //   redirect: PageEnum.BASE_HOME_ITEMS,
  meta: {
    title: "主页",
    isRoot: true,
  },
  children: [],
};

export default homeRouters;
