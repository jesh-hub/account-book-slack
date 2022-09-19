/**
 * @param {Date} date
 * @param {string} [separator]
 * @returns {string}
 */
export function getDateMonthStr(date, separator = '-') {
  return [
    date.getFullYear(),
    `${date.getMonth() + 1}`.padStart(2, '0'),
  ].join(separator);
}

export function getDateDateStr(date, separator = '-') {
  return [
    date.getFullYear(),
    `${date.getMonth() + 1}`.padStart(2, '0'),
    `${date.getDate()}`.padStart(2, '0'),
  ].join(separator);
}

export function getDateTimeMinStr(date, separator = ':') {
  return [
    `${date.getHours()}`.padStart(2, '0'),
    `${date.getMinutes()}`.padStart(2, '0'),
  ].join(separator);
}

/**
 * @param {Date} date
 * @param {string} cs - chuck separator
 * @param {string} ds - date separator
 * @param {string} ts - time separator
 * @returns {string}
 */
export function getDateStr(date, cs = ' ', ds = '-', ts = ':') {
  return [
    getDateDateStr(date, ds),
    getDateTimeMinStr(date, ts)
  ].join(cs);
}
