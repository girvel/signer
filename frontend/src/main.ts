import { CreateController } from "./controller";
import { GetEnv } from "./env";

const elements = {
  input: document.getElementById("input") as HTMLTextAreaElement,
  output: document.getElementById("output") as HTMLTextAreaElement,
  sign: document.getElementById("sign")! as HTMLInputElement,
  verify: document.getElementById("verify")! as HTMLInputElement,
  ownerName: document.getElementById("owner-name") as HTMLSpanElement,
  apiUrl: document.getElementById("api-url") as HTMLSpanElement,
};

const controller = CreateController(elements);
const env = GetEnv();

elements.ownerName.textContent = env.VITE_OWNER_NAME; // TODO maybe fetch it? or not?
elements.apiUrl.textContent = env.VITE_API_URL;

elements.sign.addEventListener('click', async (ev: MouseEvent) => {
  console.log(ev);
  await controller.Sign();
});

elements.input.addEventListener('keydown', async (ev: KeyboardEvent) => {
  if (!(ev.ctrlKey && ev.key == 'Enter')) return;
  ev.preventDefault();
  console.log(ev);
  await controller.Sign();
});

elements.verify.addEventListener('click', async (ev: MouseEvent) => {
  console.log(ev);
  await controller.Verify();
});

elements.output.addEventListener('keydown', async (ev: KeyboardEvent) => {
  if (!(ev.ctrlKey && ev.key == 'Enter')) return;
  ev.preventDefault();
  console.log(ev);
  await controller.Verify();
});

elements.output.addEventListener('input', _ => {
  controller.SetValidity("output", "undefined");
  controller.SetValidity("input", "undefined");
});

elements.input.addEventListener('input', _ => {
  controller.SetValidity("input", "undefined");
});
