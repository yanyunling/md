interface SSEMessage<T> {
  type: string;
  data: T;
}

type onMessage<T> = (message: SSEMessage<T>) => void;
