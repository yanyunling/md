import Token from "./token";
import localforage from "localforage";

const store = localforage.createInstance({
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
   * @returns
   */
  setDoc(currentDoc: CurrentDoc) {
    return store.setItem(Token.getName() + this.docCacheKey, JSON.parse(JSON.stringify(currentDoc)));
  }

  /**
   * 获取文档内容
   * @returns
   */
  async getDoc(): Promise<CurrentDoc> {
    return new Promise((resolve, reject) => {
      store
        .getItem<CurrentDoc>(Token.getName() + this.docCacheKey)
        .then((res) => {
          resolve(res);
        })
        .catch((err) => {
          console.error(err);
          resolve({
            id: "",
            content: "",
            originMD5: "",
            type: "",
            updateTime: "",
          });
        });
    });
  }

  /**
   * 清空文档缓存
   * @returns
   */
  removeDoc() {
    return store.removeItem(Token.getName() + this.docCacheKey);
  }
}

export default new DocCache();
