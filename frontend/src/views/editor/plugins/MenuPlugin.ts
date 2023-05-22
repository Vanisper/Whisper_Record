import { Command, Plugin } from "prosemirror-state";
import { EditorView } from "prosemirror-view";
import { toggleMark, setBlockType, wrapIn } from "prosemirror-commands";
import { schema } from "prosemirror-markdown";

import "./style.less";

import SourceCodeIcon from "./icons/menu/SourceCode.svg?raw";
import BoldIcon from "./icons/menu/Bold.svg?raw";
import ItalicIcon from "./icons/menu/Italic.svg?raw";
import StrikeThroughIcon from "./icons/menu/StrikeThrough.svg?raw";
import ImageIcon from "./icons/menu/Image.svg?raw";
import CodeIcon from "./icons/menu/Code.svg?raw";
import FormulaIcon from "./icons/menu/Formula.svg?raw";
import TableIcon from "./icons/menu/Table.svg?raw";
import EmojiIcon from "./icons/menu/Emoji.svg?raw";

function loadCSSString(cssString: string) {
  const style = document.createElement("style");
  style.textContent = cssString;
  document.head.appendChild(style);
}

interface PluginItems {
  command: Command;
  dom: HTMLElement;
}

class MenuView {
  items: PluginItems[];
  editorView: EditorView;
  dom: HTMLDivElement;
  constructor(items: PluginItems[], editorView: EditorView) {
    this.items = items;
    this.editorView = editorView;

    this.dom = document.createElement("div");
    this.dom.className = "menubar-wrapper";
    this.dom.appendChild(
      (() => {
        const menubar = document.createElement("div");
        menubar.classList.add("menubar");
        const menuitemWrapper = document.createElement("div");
        menuitemWrapper.classList.add("menuitem-wrapper");
        menubar.appendChild(menuitemWrapper);
        return menubar;
      })()
    );
    items.forEach(({ dom }) => {
      dom.classList.add("menuitem");
      this.dom.childNodes[0].childNodes[0].appendChild(dom);
    });
    this.update();

    this.dom.addEventListener("mousedown", (e) => {
      e.preventDefault();
      editorView.focus();
      items.forEach(({ command, dom }) => {
        if (dom.contains(e.target as Node))
          command(editorView.state, editorView.dispatch, editorView);
      });
    });
  }

  update() {
    this.items.forEach(({ command, dom }) => {
      let active = command(this.editorView.state, undefined, this.editorView);
      dom.style.display = active ? "" : "none";
    });
  }

  destroy() {
    this.dom.remove();
  }
}

function menuPlugin(items: PluginItems[]) {
  return new Plugin({
    view(editorView) {
      let menuView = new MenuView(items, editorView);
      editorView.dom?.parentNode?.insertBefore(menuView.dom, editorView.dom);
      return menuView;
    },
  });
}

// Helper function to create menu icons
function icon(innerHTML: string, name: string) {
  let span = document.createElement("span");
  span.className = "menuicon " + name;
  span.title = name;
  span.innerHTML = innerHTML;
  return span;
}

// Create an icon for a heading at the given level
function heading(level: number) {
  return {
    command: setBlockType(schema.nodes.heading, { level }),
    dom: icon("H" + level, "heading"),
  };
}
export const menu = menuPlugin([
  {
    command: (state, dispatch?, view?) => {
      return true;
    },
    dom: icon(SourceCodeIcon, "SourceCode"),
  },
  { command: toggleMark(schema.marks.strong), dom: icon(BoldIcon, "strong") },
  { command: toggleMark(schema.marks.em), dom: icon(ItalicIcon, "em") },
  {
    command: setBlockType(schema.nodes.image),
    dom: icon(ImageIcon, "image"),
  },
  {
    command: setBlockType(schema.nodes.code_block),
    dom: icon(CodeIcon, "code-block"),
  },

  { command: wrapIn(schema.nodes.blockquote), dom: icon(">", "blockquote") },
]);
