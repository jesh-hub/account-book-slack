import axios from 'axios';
import { useEffect, useState } from 'react';

const API_END_POINT = process.env.REACT_APP_ABS;

export function useGetRequest(url, params = null) {
  const [processing, setProcessing] = useState(false);
  const [response, setResponse] = useState([]);

  useEffect(() => {
    const fetch = async () => {
      try {
        setProcessing(true);
        setResponse(_ => []);
        const { data } = await axios.get(`${API_END_POINT}${url}`, { params });
        setResponse(response => data || response);
      } finally {
        setProcessing(false);
      }
    };

    fetch().then();
  }, [url, params]);

  return [response, processing];
}

export function doPostRequest(url, params) {
  return axios.post(`${API_END_POINT}${url}`, params);
}

export function doPutRequest(url, params) {
  return axios.put(`${API_END_POINT}${url}`, params);
}
