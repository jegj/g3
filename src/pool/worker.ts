import * as fs from "fs";
import { resolve } from "path";
import { parentPort } from "worker_threads";
import { decryptAESGCM } from "../crypto/";

export const filename = resolve(__filename);

interface DecryptFileInput {
  file: string;
  password: string;
}

export interface ProcessFileChunkParam {
  filePath: string;
  start: number;
  end: number;
  chunkIndex: number;
}

export function decryptFile({ file, password }: DecryptFileInput) {
  parentPort?.postMessage(`Decrypting file ${file}`);
  const content = fs.readFileSync(file, { encoding: "utf-8" });
  const decryptedContent = decryptAESGCM(content, password);
  fs.writeFileSync(file, decryptedContent);
}

export function processGistChunk(params: ProcessFileChunkParam) {
  const { filePath, start, end, chunkIndex } = params;
  const chunkSize = end - start;
  try {
    const fd = fs.openSync(filePath, "r");
    const buffer = Buffer.alloc(chunkSize);
    const bytesRead = fs.readSync(fd, buffer, 0, chunkSize, start);
    //TODO: DO SOMETHING
    fs.closeSync(fd);
  } catch (error) {
    throw new Error(`Failed to read chunk ${chunkIndex}: ${error}`);
  }
}
