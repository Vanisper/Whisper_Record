import { IWMD_DEFAULT_OPTION } from "./types";
import "./styles/index.less";
import { getUniqueId } from "./utils/random";

import { ContentState } from "./contentState";

import {
  h,
  VNode,
  init,
  toVNode,
  attributesModule,
  propsModule,
  datasetModule,
  eventListenersModule,
  styleModule,
  classModule,
} from "snabbdom";

const patch = init([
  attributesModule,
  propsModule,
  datasetModule,
  eventListenersModule,
  styleModule,
  classModule,
]);

/**
 * [ 确保容器是一个div || 将dom元素修正为div ]
 */
function getContainer(
  originContainer: HTMLElement | Element | undefined,
  options: Readonly<IWMD_DEFAULT_OPTION>
) {
  if (!originContainer) return undefined;

  const attributes = originContainer.attributes;
  const rootDom = createRootDom(options);
  // copy attributes from origin container to new div element
  const attrs: Record<string, string | number | boolean> = {};
  Array.from(attributes).forEach((attr) => {
    attrs[attr.name] = attr.value;
  });
  const container = h(
    "div",
    {
      attrs: attrs,
      class: {
        editor: true,
      },
    },
    [rootDom]
  );
  patch(originContainer, container);
  return {
    elm: container.elm as Element,
    vnode: container,
  };
}
// 创建根dom元素
const createRootDom = (options: Readonly<IWMD_DEFAULT_OPTION>) => {
  const { spellcheckEnabled, readonly } = options;
  return h("article", {
    attrs: {
      contenteditable: !readonly, // 将容器设置为可编辑
      translate: "no", // 阻止浏览器翻译内容
      autocorrect: false,
      autocomplete: "off",
      // NOTE: 浏览器无法自动纠正拼写错误的单词，除非像 MarkText 这样实现了自定义拼写纠正功能。
      spellcheck: `${!!spellcheckEnabled}`,
    },
    class: {
      "editor-article": true,
      "editor-article-readonly": readonly!,
    },
    on: {
      keydown: () => {},
    },
  });
};

class WhisperMarkDown {
  container: Element | undefined;
  containerVNode: VNode | undefined;
  options: IWMD_DEFAULT_OPTION;
  contentState: ContentState;

  constructor(container: Element | undefined, options: IWMD_DEFAULT_OPTION) {
    this.container = getContainer(container, options)?.elm;
    this.containerVNode = getContainer(container, options)?.vnode;
    this.options = options;

    this.contentState = new ContentState(this, this.options);
    if (!this.container) return;
    this.init();
  }
  // 初始化
  private init() {
    this.mutationObserver();
  }
  // 变化观察者
  private mutationObserver() {
    const { container } = this;
    if (!container) return;
    /**观察者的选项（要观察的变化）
     * attributeFilter：观察哪些属性的变化，如果设置为 null 或空数组，则观察所有属性的变化。
     * attributeOldValue：如果为 true，则记录属性变化之前的值。
     * attributes：如果为 true，则观察属性的变化。
     * characterData：如果为 true，则观察 target 的文本节点的变化。
     * characterDataOldValue：如果为 true，则记录文本变化前的其老值。
     * childList：如果为 true，则观察目标节点的子元素列表发生变化，如添加或删除子元素等。
     * subtree：如果为 true，则观察目标节点及其子孙节点的变化。
     */
    const config = { childList: true, subtree: true };
    const observer = new MutationObserver((mutationsList) => {
      /**
       * mutationsList.length:
       * 1：新块的输入
       * 2：
       */

      for (const mutation of mutationsList) {
        if (mutation.type === "childList") {
          const { removedNodes, previousSibling, target, addedNodes } =
            mutation;
          if (
            previousSibling !== null &&
            target instanceof HTMLElement &&
            target.tagName == "DIV"
          ) {
            // console.log(previousSibling);
          } else if (target instanceof HTMLElement) {
            // 此时的target是每个段落块
            // console.log(addedNodes);
          }
        }
        // console.log(mutation);
      }
    });
    observer.observe(container, config);
  }

  // 传入markdown
  private setMarkdown(markdown?: string | undefined) {
    let newMarkdown = markdown;
    this.contentState.importMarkdown(newMarkdown);
  }
}

export { WhisperMarkDown };
