import request from "@/utils/request";

class PictureApi {
  /**
   * 分页查询
   * @param tableCondition
   * @returns
   */
  page(tableCondition: PageCondition<null>) {
    return request<PageResult<PicturePageResult>>({
      method: "post",
      url: "/pic/page",
      data: tableCondition,
    });
  }

  /**
   * 根据id删除图片
   * @param id
   * @returns
   */
  delete(id: string) {
    return request({
      method: "post",
      url: "/pic/delete",
      data: { id: id },
    });
  }
}

export default new PictureApi();
