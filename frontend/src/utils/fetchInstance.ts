import axios from 'axios';

export const fetchInstance = () => {
  if (typeof window === 'undefined') return axios.create();
  return axios.create({
    baseURL: `${location.origin}/api`,
    headers: {
      'Content-Type': 'application/json',
    },
  });
};
