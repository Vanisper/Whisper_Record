import { EditorView } from "prosemirror-view";
import { EditorState } from "prosemirror-state";
import { Schema } from "prosemirror-model";
import { addListNodes } from "prosemirror-schema-list";
import {
  schema as schema_md,
  defaultMarkdownParser,
  defaultMarkdownSerializer,
} from "prosemirror-markdown";
import { exampleSetup } from "prosemirror-example-setup";
import { menu } from "./plugins/MenuPlugin";
import { baseKeymap } from "prosemirror-commands";
import { keymap } from "prosemirror-keymap";
// 代码视图
import { basicSetup, EditorView as CodeView } from "codemirror";
import { EditorState as CodeState } from "@codemirror/state";
import { markdown, markdownLanguage } from "@codemirror/lang-markdown";
import { languages } from "@codemirror/language-data";
import { oneDark } from "@codemirror/theme-one-dark";
import {
  bracketMatching,
  defaultHighlightStyle,
  HighlightStyle,
  indentOnInput,
  syntaxHighlighting,
} from "@codemirror/language";
import { tags } from "@lezer/highlight";
import {
  keymap as codeKeymap,
  highlightActiveLine,
  highlightActiveLineGutter,
  lineNumbers,
} from "@codemirror/view";
import {
  defaultKeymap as codeDefaultKeymap,
  history as codeHistory,
} from "@codemirror/commands";
import { usePostsStore } from "../../store/modules/postsStore/postsStore";
import { WritePostsContent } from "../../../wailsjs/go/backend/App";

// 将 prosemirror-schema-list 和基本 schema 放在一起形成一个支持 list 的 schema
export const mySchema = new Schema({
  nodes: addListNodes(schema_md.spec.nodes, "paragraph block*", "block"),
  marks: schema_md.spec.marks,
});
// ...exampleSetup({ schema: schema_md })
const plugins = [...exampleSetup({ schema: schema_md, menuBar: false }), menu];

const myHighlightStyle = HighlightStyle.define([
  {
    tag: tags.heading1,
    fontSize: "1.6em",
    fontWeight: "bold",
  },
  {
    tag: tags.heading2,
    fontSize: "1.4em",
    fontWeight: "bold",
  },
  {
    tag: tags.heading3,
    fontSize: "1.2em",
    fontWeight: "bold",
  },
]);

export class ProseMirrorView {
  view: EditorView;
  editorDom: HTMLElement;
  codeView: CodeView;
  codeDom: HTMLElement;
  target: Node;
  content: string;
  currView: "code" | "editor";
  constructor(target: Node, content: string) {
    this.currView = "editor";
    this.target = target;
    this.content = content;
    this.view = new EditorView(this.target, {
      state: EditorState.create({
        doc: defaultMarkdownParser.parse(content)!,
        plugins: plugins,
      }),
      dispatchTransaction: async (transaction) => {
        let newState = this.view.state.apply(transaction);
        this.view.updateState(newState);
        // 更新 markdown 内容、更新代码视图
        this.content = defaultMarkdownSerializer.serialize(this.view.state.doc);
        usePostsStore().isChange = (this.content == await usePostsStore().getCurrentFileContent());
        this.codeView.setState(this.codeState);
      },
    });
    this.codeView = new CodeView({
      parent: this.target as Element,
      state: this.codeState,
      dispatch: async (transaction) => {
        this.codeView.update([transaction]);
        // 更新 markdown 内容、更新渲染视图
        this.content = transaction.newDoc.toJSON().join("\n");
        usePostsStore().isChange = !(this.content == await usePostsStore().getCurrentFileContent());
        this.view.updateState(this.state);
      },
    });
    this.editorDom = (this.target as HTMLElement).querySelector(
      ".ProseMirror"
    )!;
    this.codeDom = (this.target as HTMLElement).querySelector(".cm-editor")!;
    this.init();
  }

  init() {
    const SourceCode = document.querySelector(".SourceCode") as HTMLElement;
    SourceCode?.addEventListener("click", (_ev) => {
      this.switchView();
    });

    document.addEventListener("keydown", async (event) => {
      if (event.ctrlKey && event.key === "s") {
        event.preventDefault();
        if (usePostsStore().currentFile) {
          if (await WritePostsContent(usePostsStore().currentFile!.pathname, this.content)) {
            window.$message.success("保存成功")
            console.log(this.content == await usePostsStore().getCurrentFileContent());

            usePostsStore().isChange = !(this.content == await usePostsStore().getCurrentFileContent());
          }
        }
      }
    });
  }

  get state() {
    return EditorState.create({
      doc: defaultMarkdownParser.parse(this.content)!,
      plugins: [...plugins,],
    });
  }

  get codeState() {
    return CodeState.create({
      doc: this.content,
      extensions: [
        codeKeymap.of([...codeDefaultKeymap]),
        lineNumbers(),
        highlightActiveLineGutter(),
        codeHistory(),
        indentOnInput(),
        bracketMatching(),
        syntaxHighlighting(myHighlightStyle),
        syntaxHighlighting(defaultHighlightStyle),
        highlightActiveLine(),
        markdown({
          base: markdownLanguage,
          codeLanguages: languages,
          addKeymap: true,
        }),
        oneDark,
        CodeView.theme({
          "&": {
            backgroundColor: "transparent !important",
            height: "100%",
            "flex-grow": 1,
            display: this.currView == "code" ? "" : "none !important",
          },
          ".cm-scroller": {
            "flex-grow": 1,
          },
        }),
        CodeView.lineWrapping,
        CodeView.updateListener.of((update) => {
          if (update.changes) {
            // onChange && onChange(update.state)
            // console.log(update);
          }
        }),
      ],
    });
  }

  switchView() {
    this.currView = this.currView == "editor" ? "code" : "editor";
    switch (this.currView) {
      case "editor":
        this.codeDom.style.setProperty("display", "none", "important");
        this.editorDom.style.display = "";
        break;

      default:
        this.codeDom.style.setProperty("display", "flex", "important");
        this.editorDom.style.display = "none";
        break;
    }
  }

  updateContent(content: string) {
    this.content = content;
    this.view.updateState(this.state);
    this.codeView.setState(this.codeState);
  }

  focus() {
    this.view.focus();
  }

  destroy() {
    this.view.destroy();
  }
}
