import Token from "./token";

interface DocCache {
  docCacheKey: string;
}

class DocCache {
  constructor() {
    this.docCacheKey = "DocCache";
  }

  /**
   * 缓存文档内容
   * @param currentDoc
   */
  setDoc(currentDoc: CurrentDoc) {
    localStorage.setItem(Token.getName() + this.docCacheKey, JSON.stringify(currentDoc));
  }

  /**
   * 获取文档内容
   * @returns
   */
  getDoc(): CurrentDoc {
    try {
      let cacheJson = localStorage.getItem(Token.getName() + this.docCacheKey);
      if (cacheJson) {
        return JSON.parse(cacheJson);
      }
    } catch (e) {}
    return {
      id: "",
      content: "",
      originContent: "",
      updateTime: "",
    };
  }

  /**
   * 清空文档缓存
   */
  removeDoc() {
    localStorage.removeItem(Token.getName() + this.docCacheKey);
  }
}

export default new DocCache();
