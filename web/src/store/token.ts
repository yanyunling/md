interface Token {
  tokenKey: string;
}

class Token {
  constructor() {
    this.tokenKey = "Token";
  }

  /**
   * 保存token
   * @param token
   */
  setToken(token: string) {
    localStorage.setItem(this.tokenKey, token);
  }

  /**
   * 清空token
   */
  removeToken() {
    localStorage.removeItem(this.tokenKey);
    location.reload();
  }

  /**
   * 查询token
   */
  getToken() {
    return localStorage.getItem(this.tokenKey);
  }
}

export default new Token();
