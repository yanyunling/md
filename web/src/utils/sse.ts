import { host, context } from "@/config";
import SSEApi from "@/api/sse";

/**
 * SSE连接
 */
export class SSE {
  private static eventSource: EventSource = null;
  private static reconnectTimeout = null;
  private static onMessageList: ((message: SSEMessage<any>) => void)[] = [];

  /**
   * 开启SSE连接
   */
  static open(isReconnect?: boolean) {
    // 关闭已有链接
    if (!isReconnect) {
      this.close();
    }

    // 获取SSE临时token
    SSEApi.getSSEToken()
      .then((res) => {
        this.eventSource = new EventSource(host + context + "/sse/" + res.data);

        // 监听消息事件
        this.eventSource.onmessage = (e) => {
          try {
            let message: SSEMessage<any> = JSON.parse(e.data);
            if (!message.type) {
              console.error("SSE消息格式不符合标准");
              return;
            }
            // 心跳不进行回调
            if (message.type === "heart") {
              return;
            }
            // 消息事件回调
            for (let func of this.onMessageList) {
              func(message);
            }
          } catch (e) {
            console.error("SSE消息反序列化失败", e);
          }
        };

        this.eventSource.onopen = () => {
          console.info("SSE连接开启");
        };

        this.eventSource.onerror = (e) => {
          console.error("SSE错误", e);
          this.reconnect();
        };
      })
      .catch(() => {
        this.reconnect();
      });
  }

  /**
   * 关闭SSE连接
   */
  static close() {
    this.onMessageList = [];
    if (this.reconnectTimeout) {
      clearTimeout(this.reconnectTimeout);
      this.reconnectTimeout = null;
    }
    if (this.eventSource) {
      this.eventSource.close();
      this.eventSource = null;
      console.info("SSE连接关闭");
    }
  }

  /**
   * 监听消息
   * @param callback
   */
  static onMessage(callback: (message: SSEMessage<any>) => void) {
    this.onMessageList.push(callback);
  }

  /**
   * 移除消息监听
   * @param callback 未传入则删除最后一个监听
   */
  static removeOnMessage(callback?: (message: SSEMessage<any>) => void) {
    if (callback) {
      // 传入回调方法则移除指定监听
      this.onMessageList = this.onMessageList.filter((item) => item !== callback);
    } else if (this.onMessageList.length > 0) {
      // 未传入回调方法则删除最后一个监听
      this.onMessageList.pop();
    }
  }

  /**
   * 重连
   */
  private static reconnect() {
    if (this.reconnectTimeout) {
      return;
    }
    this.reconnectTimeout = setTimeout(() => {
      this.reconnectTimeout = null;
      if (this.eventSource) {
        this.eventSource.close();
        this.eventSource = null;
      }
      this.open(true);
    }, 5000);
  }
}
