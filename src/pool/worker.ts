import * as fs from "fs";
import { resolve } from "path";
import { parentPort } from "worker_threads";
import { decryptAESGCM, encryptAESGCM } from "../crypto/";
import { G3File } from "../g3file";
import { GistFilesRequest, GistResponse } from "../gist/types";
import { GistDataEntry } from "../fsdata/types";
import { appendG3FSEntry } from "../fsdata";
import { G3Config } from "../config";
import { createGistFactory } from "../gist";
import { G3Dependecies } from "../types";

export const filename = resolve(__filename);

interface DecryptFileInput {
  file: string;
  password: string;
}

export function decryptFile({ file, password }: DecryptFileInput) {
  parentPort?.postMessage(`Decrypting file ${file}`);
  const content = fs.readFileSync(file, { encoding: "utf-8" });
  const decryptedContent = decryptAESGCM(content, password);
  fs.writeFileSync(file, decryptedContent);
}

export interface ProcessFileChunkParam {
  filePath: string;
  start: number;
  end: number;
  chunkIndex: number;
  g3File: G3File;
  config: G3Config;
}

export async function processGistChunk(params: ProcessFileChunkParam) {
  const { filePath, start, end, chunkIndex, g3File, config } = params;
  const chunkSize = end - start;
  try {
    const dependencies: G3Dependecies = { config };
    const createGist = createGistFactory(dependencies);
    const fd = fs.openSync(filePath, "r");
    const buffer = Buffer.alloc(chunkSize);
    const bytesRead = fs.readSync(fd, buffer, 0, chunkSize, start);
    const encryptedContent = encryptAESGCM(
      bytesRead.toString(),
      config.AES_KEY,
    );
    const gistFiles: GistFilesRequest = {
      [g3File.sortableFileName]: {
        content: String(encryptedContent),
      },
    };
    let resp: GistResponse = await createGist(
      g3File.description,
      gistFiles,
      false,
    );
    resp = deleteContentFromGistReponse(resp);
    const entries: GistDataEntry[] = [
      {
        id: resp.id,
        gistUrl: resp.url,
        gistPullUrl: resp.git_pull_url,
        files: resp.files,
      },
    ];
    await appendG3FSEntry(entries, g3File);
    fs.closeSync(fd);
  } catch (error) {
    throw new Error(`Failed to read chunk ${chunkIndex}: ${error}`);
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
