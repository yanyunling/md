import Token from "./token";
import localforage from "localforage";

localforage.config({
  name: "doc",
});

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
    localforage.setItem(Token.getName() + this.docCacheKey, JSON.stringify(currentDoc));
  }

  /**
   * 获取文档内容
   * @returns
   */
  async getDoc(): Promise<CurrentDoc | null> {
    return new Promise((resolve, reject) => {
      localforage
        .getItem<string>(Token.getName() + this.docCacheKey)
        .then((res) => {
          if (res) {
            resolve(JSON.parse(res));
          } else {
            resolve({
              id: "",
              content: "",
              originContent: "",
              type: "",
              updateTime: "",
            });
          }
        })
        .catch((err) => {
          console.error(err);
          resolve({
            id: "",
            content: "",
            originContent: "",
            type: "",
            updateTime: "",
          });
        });
    });
  }

  /**
   * 清空文档缓存
   */
  removeDoc() {
    localforage.removeItem(Token.getName() + this.docCacheKey);
  }
}

export default new DocCache();
