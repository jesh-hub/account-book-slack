import axios from 'axios';
import { useEffect, useReducer, useState } from 'react';

/** @typedef {any} ResponseDataType */
/**
 * @param {string} url
 * @param {any} params
 * @param {Array.<any>} deps
 * @param {ResponseDataType} [defaultRes]
 * @returns {[boolean,any]} - processing, responseData
 */
export default function useRequest(url, params, deps, defaultRes) {
  const [isProcessing, setIsProcessing] = useState(false);
  const [
    responseData,
    /** @type {function(data: ResponseDataType): void} */ dispatch
  ] = useReducer((_, /** @type {ResponseDataType} */ action) => action, defaultRes);

  async function fetchData() {
    setIsProcessing(true);
    try {
      const { data } = await axios.get(`${process.env.REACT_APP_ABS}${url}`, { params });
      dispatch(data ?? defaultRes);
    } catch (err) {
      // TODO useContext
    } finally {
      setIsProcessing(false);
    }
  }

  useEffect(() => {
    fetchData().then();
    // deps 외에는 의존할 변수가 없어서 disable한다.
    // eslint-disable-next-line
  }, deps);

  return [isProcessing, responseData];
}
