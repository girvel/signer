export type Env = {
  VITE_API_URL: string,
  VITE_OWNER_NAME: string,
};

export const GetEnv = (): Env => {
  let result: Record<string, string> = {};
  for (let varname of ["VITE_API_URL", "VITE_OWNER_NAME"]) {
    result[varname] = import.meta.env[varname];
    if (result[varname] === undefined) {
      throw new Error(`Missing $${varname} env variable`);
    }
  }
  return result as Env;
};
