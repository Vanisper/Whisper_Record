import { ProseMirrorView } from "../src/views/editor/EditorView";
import { MessageApiInjection } from "naive-ui/lib/message/src/MessageProvider";
import { LoadingBarApiInjection } from "naive-ui/lib/loading-bar/src/LoadingBarProvider";
import { DialogApiInjection } from "naive-ui/lib/dialog/src/DialogProvider";
declare global {
    interface Window {
        $loading: LoadingBarApiInjection
        $message: MessageApiInjection
        $dialog: DialogApiInjection

        view: ProseMirrorView
        // 语言
        $t: any
        $vue: any
        // 键盘按键记录
        $KeyboardActive?: { [T: string]: boolean }
        onKeySpacePressHold?: Function

        // 编辑 JSON 的存储对象
        opener: any
    }

}

// declare type Recordable<T = any> = Record<string, T>
