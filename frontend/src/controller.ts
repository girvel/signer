import { Api } from "./api";

export type Validity = "valid" | "invalid" | "undefined";
export type Elements = {
  input: HTMLTextAreaElement,
  output: HTMLTextAreaElement,
  sign: HTMLInputElement,
  verify: HTMLInputElement,
};

export const CreateController = (elements: Elements) => {
  return {
    _validity: {
      input: "undefined",
      output: "undefined",
    },

    SetValidity(target: "input" | "output", value: Validity) {
      if (value == this._validity[target]) return
      const targetElement = elements[target];
      targetElement.classList.remove("mod_valid", "mod_invalid");
      if (value !== "undefined") {
        targetElement.classList.add("mod_" + value);
      }
      this._validity[target] = value;
    },

    async Sign() {
      const response = await Api.post('/sign', elements.input.value);
      elements.output.value = response.data;
      this.SetValidity("input", "valid");
      this.SetValidity("output", "valid");
    },

    async Verify() {
      try {
        await Api.post('/verify', elements.output.value);
        this.SetValidity("output", "valid");
      } catch (err) {
        this.SetValidity("output", "invalid");
        // TODO report internal error
        // TODO handle all networking errors in one place?
      }
    },
  };
};
