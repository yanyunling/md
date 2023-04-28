import dayjs, { ManipulateType } from "dayjs";

/**
 * 格式化时间
 * @param date
 * @param pattern
 * @returns
 */
export const formatTime = (date: dayjs.ConfigType, pattern: string): string => {
  return dayjs(date).format(pattern);
};

/**
 * 转为时间格式
 * @param date
 * @returns
 */
export const parseTime = (date: string | number) => {
  return dayjs(date).toDate();
};

/**
 * 增加指定时间
 * @param date
 * @param value
 * @param unit
 * @returns
 */
export const addTime = (date: Date, value: number, unit: ManipulateType) => {
  return dayjs(date).add(value, unit).toDate();
};
