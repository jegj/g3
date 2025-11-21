import * as crypto from "crypto";

const ALG = "aes-256-cbc";

export function encryptAESGCM(text: string, password: string): string {
  const key = Buffer.from(password, "hex");
  const iv = crypto.randomBytes(16); // Generate a random IV for each encryption
  const cipher = crypto.createCipheriv(ALG, key, iv);
  let encrypted = cipher.update(text, "utf8", "hex");
  encrypted += cipher.final("hex");
  return iv.toString("hex") + encrypted; // Prepend IV for decryption
}

export function decryptAESGCM(encryptedText: string, password: string): string {
  const key = Buffer.from(password, "hex");
  const ivFromCipher = Buffer.from(encryptedText.slice(0, 32), "hex");
  const encryptedData = encryptedText.slice(32);
  const decipher = crypto.createDecipheriv(ALG, key, ivFromCipher);
  let decrypted = decipher.update(encryptedData, "hex", "utf8");
  decrypted += decipher.final("utf8");
  return decrypted;
}
