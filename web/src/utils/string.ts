/**
 * 生成uuid
 * @returns
 */
export const uuid = () => {
  const s4 = () => (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
  return s4() + s4() + s4() + s4();
};

/**
 * 生成随机字符串
 * @param {*} len 长度
 * @param {*} onlyNumber 数字型
 */
export const randomStr = (len: number, onlyNumber?: boolean) => {
  let result = "";
  if (len <= 0) {
    return result;
  }
  let optional;
  if (onlyNumber) {
    optional = "0123456789";
  } else {
    optional = "abcdefghijklmnopqrstuvwxyz0123456789";
  }
  const optionalLen = optional.length;
  for (let i = 0; i < len; i++) {
    result += optional.charAt(Math.floor(Math.random() * optionalLen));
  }
  return result;
};

/**
 * 获取字符串实际px
 * @param text 字符串内容
 * @param fontPx 字号
 * @param fontFamily 字体，默认Arial
 * @returns
 */
export const stringWidth = (text: string, fontPx: number, fontFamily: string = "Arial") => {
  const canvas = document.createElement("canvas");
  const context = canvas.getContext("2d");
  if (!context) {
    return 0;
  }
  context.font = fontPx + "px " + fontFamily;
  return context.measureText(text).width;
};
