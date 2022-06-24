import XLSX from "xlsx";

export class Excel {
  /**
   * 获取excel的sheet
   * @param data
   * @param sheetNum 从0开始
   * @returns
   */
  static getExcelSheet(data: string | ArrayBuffer, sheetNum: number) {
    const workbook = XLSX.read(data, { type: "array" });
    const firstSheetName = workbook.SheetNames[sheetNum];
    return workbook.Sheets[firstSheetName];
  }

  /**
   * 获取表格数据
   * @param data
   * @param sheetNum 从0开始
   * @returns
   */
  static getExcelJson<T>(data: string | ArrayBuffer, sheetNum: number) {
    const sheet = Excel.getExcelSheet(data, sheetNum);
    return XLSX.utils.sheet_to_json<T>(sheet);
  }

  /**
   * 获取表头行
   * @param sheet
   * @returns
   */
  static getExcelHeaders(data: string | ArrayBuffer, sheetNum: number) {
    const sheet = Excel.getExcelSheet(data, sheetNum);
    const headers = [];
    if (sheet["!ref"]) {
      const range = XLSX.utils.decode_range(sheet["!ref"]);
      let C;
      const R = range.s.r;
      for (C = range.s.c; C <= range.e.c; ++C) {
        let cell = sheet[XLSX.utils.encode_cell({ c: C, r: R })];
        let hdr = "UNKNOWN " + C;
        if (cell && cell.t) {
          hdr = XLSX.utils.format_cell(cell);
        }
        headers.push(hdr);
      }
    }
    return headers;
  }

  /**
   * 格式化excel中的时间
   * @param date
   * @returns
   */
  static formatExcelDate(date: number) {
    const utcDays = Math.floor(date - 25569);
    const utcValue = utcDays * 86400;
    const dateInfo = new Date(utcValue * 1000);
    const fractionalDay = date - Math.floor(date) + 0.0000001;
    let totalSeconds = Math.floor(86400 * fractionalDay);
    const seconds = totalSeconds % 60;
    totalSeconds -= seconds;
    const hours = Math.floor(totalSeconds / (60 * 60));
    const minutes = Math.floor(totalSeconds / 60) % 60;
    const result = new Date(dateInfo.getFullYear(), dateInfo.getMonth(), dateInfo.getDate(), hours, minutes, seconds);
    if (result instanceof Date && !isNaN(result.getTime())) {
      return result;
    }
    return null;
  }
}
