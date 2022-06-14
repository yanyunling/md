interface Doc {
  id: string;
  name: string;
  content: string;
  bookId: string;
  createTime?: number;
  updateTime?: number;
}

interface CurrentDoc {
  id: string;
  content: string;
  originContent: string;
  updateTime: string;
}
