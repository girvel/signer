import axios from 'axios';

export const Api = axios.create({
  baseURL: 'http://localhost:8080',  // TODO hardcoded, use env
  headers: {
    'Content-Type': 'text/plain',
  }
});
