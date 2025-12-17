import { GistResponse } from "./types";

export function deleteContentFromGistReponse(resp: GistResponse) {
  Object.values(resp.files).forEach((file) => {
    if (file && file.content) {
      file.content = "";
    }
  });
  return resp;
}
