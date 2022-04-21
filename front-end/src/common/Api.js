import axios from 'axios';

/**
 * @param {string} url
 * @param params
 * @returns {Promise<any>}
 */
export async function get(url, params) {
  const { data } = await axios.get(url, { params });
  return data;
}

export async function getWithDispatch(dispatch, url, params) {
  dispatch({ processing: true });
  const data = await get(url, params);
  dispatch({ processing: false, data });
}
