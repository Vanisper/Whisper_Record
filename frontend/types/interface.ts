export interface IProjectTree {
    files: IProjectFile[],
    id: string,
    folders: IProjectTree[],
    isDirectory: boolean,
    isFile: boolean,
    isMarkdown: boolean,
    isCollapsed: boolean,
    name: string,
    pathname: string,
}

export interface IProjectFile {
    birthTime: string
    id: string
    isDirectory: string
    isFile: string
    isMarkdown: string
    name: string
    pathname: string
}