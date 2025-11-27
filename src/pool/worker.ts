import * as fs from "fs";
import { resolve } from "path";
import { parentPort } from "worker_threads";
import { decryptAESGCM } from "../crypto/";

export const filename = resolve(__filename);

interface DecryptFileInput {
  file: string;
  password: string;
}

interface ProcessFileChunkParam {
  filePath: string;
  chunkSize: number;
  offset: number;
  chunkNumber: number;
}

export function decryptFile({ file, password }: DecryptFileInput) {
  parentPort?.postMessage(`Decrypting file ${file}`);
  const content = fs.readFileSync(file, { encoding: "utf-8" });
  const decryptedContent = decryptAESGCM(content, password);
  fs.writeFileSync(file, decryptedContent);
}

export function processGistChunk(params: ProcessFileChunkParam) {
  const fd = fs.openSync(params.filePath, "r");
  const buffer = Buffer.alloc(params.chunkSize);
  const bytesRead = fs.readSync(fd, buffer, 0, params.chunkSize, params.offset);
  //TODO: DO SOMETHING
  fs.closeSync(fd);
}
