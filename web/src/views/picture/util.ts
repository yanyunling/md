import { openFiles, getFileName } from "@/utils/file";
import Compressor from "compressorjs";
import PictureApi from "@/api/picture";
import { ElMessage } from "element-plus";

/**
 * 上传图片
 */
export const uploadPicture = () => {
  return new Promise((resolve, reject) => {
    openFiles(false, ".jpg,.jpeg,.png,.gif").then((fileList) => {
      const fileName = getFileName(fileList[0]);
      if (fileName.ext !== "jpg" && fileName.ext !== "jpeg" && fileName.ext !== "png" && fileName.ext !== "gif") {
        ElMessage.warning("请上传以下格式的图片：jpg、jpeg、png、gif");
        reject();
        return;
      }
      if (fileList[0].size > 1000 * 1000 * 20) {
        ElMessage.warning("图片大小不可超过20MB");
        reject();
        return;
      }
      // 压缩图片
      new Compressor(fileList[0], {
        quality: 0.8,
        maxWidth: 100,
        maxHeight: 100,
        success(result) {
          // 上传图片
          const formData = new FormData();
          formData.append("picture", fileList[0]);
          formData.append("thumbnail", result);
          PictureApi.upload(formData)
            .then((res) => {
              ElMessage.success(res.message);
              resolve(res.data);
            })
            .catch(() => {
              reject();
            });
        },
        error(err) {
          ElMessage.warning("图片压缩失败：" + err.message);
          reject();
        },
      });
    });
  });
};
