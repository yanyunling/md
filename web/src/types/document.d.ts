interface Doc {
  id: string;
  name: string;
  content: string;
  type?: string;
  published?: boolean;
  bookId: string;
  createTime?: number;
  updateTime?: number;
}

interface CurrentDoc {
  id: string;
  content: string;
  originContent: string;
  type: string;
  updateTime: string;
}
