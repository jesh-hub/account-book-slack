import axios from 'axios';
import { useEffect, useState } from 'react';

const API_END_POINT = process.env.REACT_APP_ABS;

export function useGetRequest(url, params) {
  const [processing, setProcessing] = useState(false);
  const [response, setResponse] = useState([]);

  useEffect(() => {
    const fetch = async () => {
      try {
        setProcessing(true);
        const { data } = await axios.get(`${API_END_POINT}${url}`, { params });
        setResponse(data);
      } finally {
        setProcessing(false);
      }
    };

    fetch().then();
  // eslint-disable-next-line
  }, []);

  return [response, processing];
}
