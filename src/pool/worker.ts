import * as fs from "fs";
import { resolve } from "path";
import { decryptAESGCM } from "../crypto/";

export const filename = resolve(__filename);

interface DecryptFileInput {
  file: string;
  password: string;
}

export function decryptFile({ file, password }: DecryptFileInput) {
  const content = fs.readFileSync(file);
  const decryptedContent = decryptAESGCM(content, Buffer.from(password));
  fs.writeFileSync(file, decryptedContent);
}
