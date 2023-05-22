// Because the sidebar also use the file icons, So I put this file out of floatBox directory.
import '@marktext/file-icons/build/index.css'
// @ts-ignore
import fileIcons from '@marktext/file-icons'

type IconTables = typeof fileIcons

interface FileIcons extends IconTables {
  getClassByName: (name: string) => string | null
  getClassByLanguage: (lang: string) => string | null
}

(fileIcons as FileIcons).getClassByName = function (name: string) {
  const icon = fileIcons.matchName(name)

  return icon ? icon.getClass(0, false) : null
};

(fileIcons as FileIcons).getClassByLanguage = function (lang: string) {
  const icon = fileIcons.matchLanguage(lang)

  return icon ? icon.getClass(0, false) : null
};

export default fileIcons as FileIcons

