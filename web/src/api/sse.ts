import request from "@/utils/request";

class SSEApi {
  /**
   * 获取SSE临时token
   * @returns
   */
  getSSEToken() {
    return request({
      method: "post",
      url: "/sse-token",
    });
  }
}

export default new SSEApi();
