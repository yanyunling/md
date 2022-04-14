import { authRequest } from "@/utils/request";
import sha256 from "crypto-js/sha256";

class TokenApi {
  /**
   * 注册
   * @param name
   * @param password
   * @returns
   */
  signUp(name: string, password: string) {
    return authRequest({
      method: "post",
      url: "/token/register",
      data: { name: name, password: sha256(password).toString() },
    });
  }

  /**
   * 登录
   * @param name
   * @param password
   * @returns
   */
   signIn(name: string, password: string) {
    return authRequest<TokenResult>({
      method: "post",
      url: "/token/get",
      data: { name: name, password: sha256(password).toString() },
    });
  }
}

export default new TokenApi();
