import {
  App,
  ComponentPublicInstance,
  createApp,
  Directive,
  DirectiveBinding,
} from "vue";
import ContextmenuComponent from "./Contextmenu.vue";
import { IClickMenuItem } from "./index.d";

const __ctxmenu__ = "__ctxmenu__";

interface Element extends HTMLElement {
  __ctxmenu__?: (event: MouseEvent) => void;
}

interface Instance extends ComponentPublicInstance {
  axis: { x: number; y: number };
  el: HTMLElement;
  menus: IClickMenuItem[];
  isDark: boolean;
  removeContextMenu: () => void;
}

const contextmenuListener = ({
  el,
  event,
  binding,
}: {
  el: HTMLElement;
  event: MouseEvent;
  binding: DirectiveBinding<string>;
}) => {
  event.stopPropagation();
  event.preventDefault();
  let instance: ComponentPublicInstance | null, mask: HTMLElement | null;

  const menus = binding.value;
  if (!menus) return;

  const isDark = binding.modifiers!.dark;

  const removeContextMenu = () => {
    if (instance) {
      document.body.removeChild(instance.$el);
      instance = null;
    }
    if (mask) {
      mask.removeEventListener("contextmenu", handleMaskContextmenu);
      mask.removeEventListener("click", removeContextMenu);
      document.body.removeChild(mask);
      mask = null;
    }
    el.classList.remove("contextmenu-active");
    document.body.removeEventListener("scroll", removeContextMenu);
    window.removeEventListener("resize", removeContextMenu);
  };

  const handleMaskContextmenu = (event: MouseEvent) => {
    event.preventDefault();
    removeContextMenu();
  };

  removeContextMenu();

  mask = document.createElement("div");
  mask.style.cssText = `
    position: fixed;
    left: 0;
    top: 0;
    width: 100vw;
    height: 100vh;
    z-index: 9998;
  `;
  document.body.appendChild(mask);
  // 传递初始化参数
  const app = createApp(ContextmenuComponent, {
    axis: { x: event.x, y: event.y },
    el,
    menus,
    isDark,
    removeContextMenu,
  });
  instance = app.mount(document.createElement("div"));

  document.body.appendChild(instance.$el);
  el.classList.add("contextmenu-active");

  mask.addEventListener("contextmenu", handleMaskContextmenu);
  mask.addEventListener("click", removeContextMenu);
  document.body.addEventListener("scroll", removeContextMenu);
  window.addEventListener("resize", removeContextMenu);
};

const ContextmenuDirective: Directive = {
  mounted(el: Element, binding: DirectiveBinding) {
    el[__ctxmenu__] = (event: MouseEvent) =>
      contextmenuListener({ el, event, binding });
    el.addEventListener("contextmenu", el[__ctxmenu__]);
  },
  //   updated(el: Element, binding: DirectiveBinding) {
  //     el[__ctxmenu__] = (event: MouseEvent) =>
  //       contextmenuListener({ el, event, binding });
  //     el.addEventListener("contextmenu", el[__ctxmenu__]);
  //   },
  unmounted(el: Element) {
    if (el && el[__ctxmenu__]) {
      el.removeEventListener("contextmenu", el[__ctxmenu__]);
      delete el[__ctxmenu__];
    }
  },
};

export const directives = {
  contextmenu: ContextmenuDirective,
};

export default {
  install(app: App) {
    // 注册具名组件
    // app.component("contextmenu", ContextmenuComponent);
    // 注册全局指令
    app.directive("contextmenu", ContextmenuDirective);
  },
};
