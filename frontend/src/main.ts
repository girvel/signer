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
      output.value = response.data;
    });
};

document.getElementById("verify")!.onclick = ev => {
  console.log(ev);
  output.classList.remove("mod_valid", "mod_invalid");
  api.post('/verify', output.value)
    .then(_ => {
        output.classList.add("mod_valid");
    })
    .catch(err => {
      if (err.status < 500) {
        output.classList.add("mod_invalid");
      } else {
        // TODO report internal error
        // TODO handle all networking errors in one place?
      }
    });
};
