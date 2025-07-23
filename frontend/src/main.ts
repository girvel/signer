import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080',  // TODO hardcoded, use env
  headers: {
    'Content-Type': 'text/plain',
  }
});

const elements = {
  input: document.getElementById("input") as HTMLTextAreaElement,
  output: document.getElementById("output") as HTMLTextAreaElement,
  sign: document.getElementById("sign")! as HTMLInputElement,
  verify: document.getElementById("verify")! as HTMLInputElement,
};

type Validity = "valid" | "invalid" | "undefined";

var validity: Record<string, Validity> = {
  input: "undefined",
  output: "undefined",
};

const setValidity = (target: "input" | "output", value: Validity) => {
  // TODO assert validity[target]
  if (value == validity[target]) return
  const targetElement = elements[target];
  targetElement.classList.remove("mod_valid", "mod_invalid");
  if (value !== "undefined") {
    targetElement.classList.add("mod_" + value);
  }
  validity[target] = value;
};

elements.sign.onclick = ev => {
  console.log(ev);
  api.post('/sign', elements.input.value)
    .then(response => {
      elements.output.value = response.data;
    });
};

elements.verify.onclick = ev => {
  console.log(ev);
  api.post('/verify', elements.output.value)
    .then(_ => {
      setValidity("output", "valid");
    })
    .catch(err => {
      if (err.status < 500) {
        setValidity("output", "invalid");
      } else {
        // TODO report internal error
        // TODO handle all networking errors in one place?
      }
    });
};

elements.output.oninput = _ => {
  setValidity("output", "undefined");
};
