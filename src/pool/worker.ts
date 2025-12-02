import * as fs from "fs";
import { resolve } from "path";
import { parentPort } from "worker_threads";
import { G3Config } from "../config";
import { decryptAESGCM, encryptAESGCM } from "../crypto/";
import { appendG3FSEntry } from "../fsdata";
import { GistDataEntry } from "../fsdata/types";
import { G3File } from "../g3file";
import { createGistFactory } from "../gist";
import { GistFilesRequest, GistResponse } from "../gist/types";
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
  sortableFileName: string;
  start: number;
  end: number;
  chunkIndex: number;
  g3File: G3File;
  config: G3Config;
}

export async function processGistChunk(params: ProcessFileChunkParam) {
  const { filePath, sortableFileName, start, end, chunkIndex, g3File, config } =
    params;
  const chunkSize = end - start;
  parentPort?.postMessage(
    `Processing gist chunk ${chunkIndex} of size ${chunkSize}`,
  );
  try {
    const dependencies: G3Dependecies = { config };
    const createGist = createGistFactory(dependencies);
    const fd = fs.openSync(filePath, "r");
    const buffer = Buffer.alloc(chunkSize);
    const bytesRead = fs.readSync(fd, buffer, 0, chunkSize, start);
    console.log("bytesRead:", bytesRead);
    const encryptedContent = encryptAESGCM(
      buffer.toString("utf-8"),
      config.AES_KEY,
    );
    console.log(
      "--->g3File.sortableFileName:",
      sortableFileName,
      g3File.description,
      g3File.filepath,
    );
    const gistFiles: GistFilesRequest = {
      [sortableFileName]: {
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
