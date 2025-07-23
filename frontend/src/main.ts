import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080',  // TODO hardcoded, use env
  headers: {
    'Content-Type': 'text/plain',
  }
});

const input = document.getElementById("input") as HTMLTextAreaElement;
const output = document.getElementById("output") as HTMLTextAreaElement;

document.getElementById("sign")!.onclick = ev => {
  console.log(ev);
  api.post('/sign', input.value)
    .then(response => {
      console.log(response);
      output.value = response.data;
    });
};

document.getElementById("verify")!.onclick = ev => {
  console.log(ev);
};
