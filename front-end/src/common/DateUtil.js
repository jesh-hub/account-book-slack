/**
 * @param {Date} date
 * @param {string} [separator]
 * @returns {string}
 */
export function getDateMonthStr(date, separator = '-') {
  return [
    date.getFullYear(),
    `${date.getMonth() + 1}`.padStart(2, '0')
  ].join(separator);
}