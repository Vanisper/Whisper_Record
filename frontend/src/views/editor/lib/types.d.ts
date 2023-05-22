import Color from "color";

export interface IWMD_DEFAULT_OPTION {
  readonly?: boolean;
  spellcheckEnabled?: boolean; // 拼写检查
  theme?: ITheme; // 主题配置
  markdown?: string;
}

export interface ITheme {
  dark: {
    primary: string;
    secondary: string;
  };
  light: {
    primary: string;
    secondary: string;
  };
  custom: {
    primary: string;
    secondary: string;
  };
}
