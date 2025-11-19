import * as fs from "fs";
import { resolve } from "path";
import { parentPort } from "worker_threads";
import { decryptAESGCM } from "../crypto/";

export const filename = resolve(__filename);

interface DecryptFileInput {
  file: string;
  password: string;
}

export function decryptFile({ file, password }: DecryptFileInput) {
  parentPort?.postMessage(`Decrypting file ${file}`);
  const content = fs.readFileSync(file);
  const decryptedContent = decryptAESGCM(content, Buffer.from(password));
  fs.writeFileSync(file, decryptedContent);
}
