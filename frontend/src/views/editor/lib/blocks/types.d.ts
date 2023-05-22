interface IBaseBlock {
  /** getUniqueId(): string */
  key: string | null;
  // HTML 标签
  type: keyof HTMLElementTagNameMap;
  editable: boolean;
  parent: null | string;
  preSibling: null | string;
  nextSibling: null | string;
  children: IBlock[];
  text?: string;
  functionType?: TypeBlockKey | TypeBlockValue;
}

interface IParagraphBlock extends IBaseBlock {
  type: "p";
  text?: string;
  children: IParagraphContentBlock[];
}

interface IParagraphContentBlock extends IBaseBlock {
  type: "span";
  text?: string;
}

interface IHeadingBlock extends IBaseBlock {
  type: "h1" | "h2" | "h3" | "h4" | "h5" | "h6";
  text: string;
}

interface ITableBlock extends IBaseBlock {
  type: "table";
  rows: string[][];
}

interface IBlockQuoteBlock extends IBaseBlock {
  type: "pre";
  text: string;
}

interface IBlockCodeBlock extends IBaseBlock {
  type: "pre";
  text: string;
}

interface IInlineCodeBlock extends IBaseBlock {
  type: "span";
  text: string;
}

interface IVideoBlock extends IBaseBlock {
  type: "video";
  src: string;
  name: string;
  poster: string;
  text: string;
  width: number;
  height: number;
}

interface IImageBlock extends IBaseBlock {
  type: "image";
  src: string;
  row: "89ds02j0sk1";
  text: string;
  width: number;
  height: number;
  alt: string;
}

interface IAttachBlock extends IBaseBlock {
  type: "div";
  name: string;
  url: string;
  size: number;
  ext: string;
}

interface IOrdeListBlock extends IBaseBlock {
  type: "ol";
  text: string;
}

interface IUnOrderListBlock extends IBaseBlock {
  type: "ul";
  text: string;
}

interface ITodoListBlock extends IBaseBlock {
  type: "ul";
  text: string;
  checked: boolean;
}

interface IDividerBlock extends IBaseBlock {
  type: "div";
  text: string;
}

export type IBlock = IBaseBlock;

export enum EnumBlock {
  "PARAGRAPH" = "paragraph", // 段落
  "PARAGRAPH_CONTENT" = "paragraphContent", // 段落
  "HEADING" = "heading", // 标题
  "TABLE" = "table", // 表格
  "BLOCKQUOTE" = "blockquote", // 引用块
  "BLOCKCODE" = "blockcode", // 代码块
  "BLOCKCODE_CONTENT" = "blockcodeContent",
  "INLINECODE" = "inlinecode", // 行内代码块
  "VIDEO" = "video", // 视频
  "IMAGE" = "image", // 图片
  "ATTACH" = "attach", // 附件
  "ORDERLIST" = "orderlist", // 有序列表
  "UNORDERLIST" = "unorderlist", // 无序列表
  "TODOLIST" = "todolist", // 任务清单
  "DIVIDER" = "divider", // 分隔符
}

export type TypeBlockKey = keyof typeof EnumBlock;
export type TypeBlockValue = `${EnumBlock}`;

export type FunctionType =
  /** 段落内容 */
  | "paragraphContent"
  /** 代码内容 */
  | "codeContent"
  /** 代码语言输入 */
  | "languageInput"
  /** 单元格内容 */
  | "cellContent"
  /** 脚注 */
  | "footnote"
  | "atxLine"
  | "table"
  | "footnoteInput"
  | "";

export type PreFunnctionType =
  | "fencecode"
  | "html"
  | "multiplemath"
  | "flowchart"
  | "mermaid"
  | "sequence"
  | "plantuml"
  | "vega-lite";

export interface IBlockExtras {
  functionType?: TypeBlockKey | TypeBlockValue;
  text?: string;
}
