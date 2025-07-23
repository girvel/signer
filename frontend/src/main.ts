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

const sign = async () => {
  const response = await api.post('/sign', elements.input.value);
  elements.output.value = response.data;
  setValidity("input", "valid");
  setValidity("output", "valid");
};

const verify = async () => {
  try {
    await api.post('/verify', elements.output.value);
    setValidity("output", "valid");
  } catch (err) {
    setValidity("output", "invalid");
    // TODO report internal error
    // TODO handle all networking errors in one place?
  }
};

elements.sign.addEventListener('click', async (ev: MouseEvent) => {
  console.log(ev);
  await sign();
});

elements.input.addEventListener('keydown', async (ev: KeyboardEvent) => {
  if (!(ev.ctrlKey && ev.key == 'Enter')) return;
  ev.preventDefault();
  console.log(ev);
  await sign();
});

elements.verify.addEventListener('click', async (ev: MouseEvent) => {
  console.log(ev);
  await verify();
});

elements.output.addEventListener('keydown', async (ev: KeyboardEvent) => {
  if (!(ev.ctrlKey && ev.key == 'Enter')) return;
  ev.preventDefault();
  console.log(ev);
  await verify();
});

elements.output.addEventListener('input', _ => {
  setValidity("output", "undefined");
  setValidity("input", "undefined");
});

elements.input.addEventListener('input', _ => {
  setValidity("input", "undefined");
});
