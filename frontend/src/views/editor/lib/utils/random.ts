const ID_PREFIX = "n-";
let id = 0;

export const getUniqueId = () => `${ID_PREFIX}${id++}`;

export const getLongUniqueId = () =>
  `${getUniqueId()}-${(+new Date()).toString(32)}`;

export function createGUID() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
    let r = (Math.random() * 16) | 0;
    let v = c === "x" ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}
