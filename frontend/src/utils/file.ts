import path from "path"
import fs from "fs"
/**
 * *获取上传的文件数据
 * @param { File } file 文件对象
 */
export const readFile = (file: File) => {
    return new Promise((resolve: Function) => {
        try {
            const reader = new FileReader()
            reader.onload = (evt: ProgressEvent<FileReader>) => {
                if (evt.target) {
                    resolve(evt.target.result)
                }
            }
            reader.readAsText(file)
        } catch (error) {
            // @ts-ignore
            window['$message'].error('文件读取失败！')
        }
    })
}

/**
 * * 通过 a 标签下载数据
 * @param url
 * @param filename
 * @param fileSuffix
 */
export const downloadByA = (url: string, filename = new Date().getTime(), fileSuffix?: string) => {
    const ele = document.createElement('a') // 创建下载链接
    ele.download = `${filename}.${fileSuffix}` //设置下载的名称
    ele.style.display = 'none' // 隐藏的可下载链接
    // 字符内容转变成blob地址
    ele.href = url
    // 绑定点击时间
    document.body.appendChild(ele)
    ele.click()
    // 然后移除
    document.body.removeChild(ele)
}

/**
 * * 下载数据
 * @param { string } content 数据内容
 * @param { ?string } filename 文件名称（默认随机字符）
 * @param { ?string } fileSuffix 文件名称（默认随机字符）
 */
export const downloadTextFile = (
    content: string,
    filename = new Date().getTime(),
    fileSuffix?: string
) => {
    // 字符内容转变成blob地址
    const blob = new Blob([content])
    downloadByA(URL.createObjectURL(blob), filename, fileSuffix)
}


/**
 * Check if the both paths point to the same file.
 *
 * @param {string} pathA The first path.
 * @param {string} pathB The second path.
 * @param {boolean} [isNormalized] Are both paths already normalized.
 */
export const isSamePathSync = (pathA: string, pathB: string, isNormalized: boolean = false) => {
    if (!pathA || !pathB) return false
    const a = isNormalized ? pathA : path.normalize(pathA)
    const b = isNormalized ? pathB : path.normalize(pathB)
    if (a.length !== b.length) {
        return false
    } else if (a === b) {
        return true
    } else if (a.toLowerCase() === b.toLowerCase()) {
        try {
            const fiA = fs.statSync(a)
            const fiB = fs.statSync(b)
            return fiA.ino === fiB.ino
        } catch (_) {
            // Ignore error
        }
    }
    return false
}