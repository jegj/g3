import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { encryptAESGCM } from "../crypto";
import { appendG3Entry, getFileContent, getFileSizeMb } from "../fsdata";
import { GistDataEntry } from "../fsdata/types";
import { createG3FileFactory } from "../g3file";
import { createGistFactory } from "../gist";
import { GistFilesRequest, GistResponse } from "../gist/types";
import { G3Dependecies } from "../types";

const GIST_FILE_SIZE_MB = 10;

export default async function cp(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const createGist = createGistFactory(dependencies);
  const description = argv.description as string;
  const file = argv.file as string;
  const g3File = await createG3File(file);
  const fileSizeMb = await getFileSizeMb(g3File);

  //https://medium.com/@serhiisamoilenko/speeding-up-file-parsing-with-multi-threading-in-nodejs-and-typescript-9e91728cf607

  if (g3File.exists) {
    if (fileSizeMb < GIST_FILE_SIZE_MB) {
    } else {
      //TODO: FOR BIGGER FILES THAT ALREADY EXISTS
    }
  } else {
    if (fileSizeMb < GIST_FILE_SIZE_MB) {
      const content = await getFileContent(g3File);
      const encryptedContent = encryptAESGCM(
        content,
        Buffer.from(config.AES_KEY),
      );
      const gistFiles: GistFilesRequest = {
        [g3File.filename]: {
          content: String(encryptedContent),
        },
      };
      let resp: GistResponse = await createGist(description, gistFiles, false);
      resp = deleteContentFromGistReponse(resp);
      const entries: GistDataEntry[] = [
        {
          id: resp.id,
          gistUrl: resp.url,
          files: resp.files,
        },
      ];
      await appendG3Entry(entries, g3File);
    } else {
      // TODO: FOR BIGGER FILES THAT DOEST NOT EXISTS
    }
  }
}

function deleteContentFromGistReponse(resp: GistResponse) {
  Object.values(resp.files).forEach((file) => {
    if (file && file.content) {
      file.content = "";
    }
  });
  return resp;
}
