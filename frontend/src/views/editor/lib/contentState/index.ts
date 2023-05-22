import { WhisperMarkDown } from "..";
import { IWMD_DEFAULT_OPTION } from "../types";
import { IBlock, IBlockExtras, TypeBlockKey } from "../blocks/types";
import { BLOCK_SCOPED_SYMBOL } from "@babel/types";
import { getUniqueId } from "../utils/random";
import escapeCharactersMap, {
  escapeCharacters,
} from "../parser/escapeCharacter";

class ContentState {
  blocks: IBlock[];
  constructor(context: WhisperMarkDown, options: IWMD_DEFAULT_OPTION) {
    this.blocks = [this.createBlockP()];
    console.log(this.blocks);
  }

  /**
   * 这是一个 MarkText 中的块，它可以表示一个段落（在 GFM 中称为块语法）或者一个段落内的某一行。
   * 一个 `span` 块必须在一个 `p` 块或者 `pre` 块中，并且 `p` 块中只能有 `span` 块作为其子节点。
   */
  createBlock(type: keyof HTMLElementTagNameMap, extras: IBlockExtras = {}) {
    const key = getUniqueId();
    const blockData: IBlock = {
      key,
      text: "",
      type: type,
      editable: true,
      parent: null,
      preSibling: null,
      nextSibling: null,
      children: [],
    };

    // give span block a default functionType `paragraphContent`
    if (type === "span" && !extras.functionType) {
      blockData.functionType = "paragraphContent";
    }

    if (extras.functionType === "blockcodeContent" && extras.text) {
      const CHAR_REG = new RegExp(`(${escapeCharacters.join("|")})`, "gi");
      extras.text = extras.text.replace(CHAR_REG, (_, p) => {
        return escapeCharactersMap[p];
      });
    }

    Object.assign(blockData, extras);
    return blockData;
  }
  createBlockP(text = "") {
    const pBlock = this.createBlock("p");
    const contentBlock = this.createBlock("span", { text });
    this.appendChild(pBlock, contentBlock);
    return pBlock;
  }
  appendChild(parent: IBlock, block: IBlock) {
    const len = parent.children.length;
    const lastChild = parent.children[len - 1];
    parent.children.push(block);
    block.parent = parent.key;
    if (lastChild) {
      lastChild.nextSibling = block.key;
      block.preSibling = lastChild.key;
    } else {
      block.preSibling = null;
    }
    block.nextSibling = null;
  }

  public importMarkdown(md: string | undefined) {
    console.log("打印传入的markdown数据:", md);
  }

  /**
   * importBlock
   */
  public importBlock(blocks: IBlock[]) {
    console.log(blocks);
  }
}

export { ContentState };
