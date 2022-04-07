interface HTMLInputEvent extends Event {
  target: HTMLInputElement & EventTarget;
}

/**
 * 上传文件类型
 */
export const enum InputAccept {
  xlsx = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
  image = "image/*",
  jpg = "image/jpeg",
  png = "image/png",
  gif = "image/gif",
}

/**
 * 上传文件
 * @param multiple
 * @param accept
 * @returns
 */
export const openFiles = (multiple?: boolean, accept?: string) => {
  return new Promise<FileList>((resolve) => {
    const input = document.createElement("input");
    input.type = "file";
    if (accept) {
      input.accept = accept;
    }
    if (multiple) {
      input.multiple = true;
    }
    input.onchange = (e: Event) => {
      let event = e as HTMLInputEvent;
      const elem = event.target;
      resolve(elem.files ?? new FileList());
    };
    input.click();
  });
};

/**
 * 获取文件名和后缀
 * @param file
 * @returns
 */
export const getFileName = (file: File) => {
  let splits = file.name.split(".");
  let name = splits.slice(0, splits.length - 1).join(".");
  let ext = "";
  if (splits.length > 1) {
    ext = splits[splits.length - 1];
  }
  return { name, ext };
};

/**
 * 将文件读取为String
 * @param file
 * @returns
 */
export const readFileToString = (file: File) => {
  return new Promise<string>((resolve) => {
    const reader = new FileReader();
    reader.onload = (e: ProgressEvent<FileReader>) => {
      const text = e.target?.result + "";
      resolve(text);
    };
    reader.readAsText(file);
  });
};

/**
 * 将文件读取为ArrayBuffer
 * @param file
 * @returns
 */
export const readFileToArrayBuffer = (file: File) => {
  return new Promise<string | ArrayBuffer>((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsArrayBuffer(file);
    reader.onload = (e: ProgressEvent<FileReader>) => {
      const data = e.target?.result;
      if (data) {
        resolve(data);
      } else {
        reject();
      }
    };
    reader.onerror = () => {
      reject();
    };
  });
};

/**
 * blob转Base64
 * @param blob
 * @param callback
 */
export const blobToBase64 = (blob: Blob, callback: Function) => {
  let reader = new FileReader();
  reader.onload = (e) => {
    callback(e.target?.result);
  };
  reader.readAsDataURL(blob);
};
