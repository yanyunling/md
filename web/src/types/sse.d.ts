interface SSEMessage<T> {
  type: string;
  data: T;
}
