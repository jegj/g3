import * as crypto from "crypto";
import { g3Error } from "../error";

const algorithm = "aes-256-cbc";
const iterations = 100000;
const keyLength = 32;
const ivLength = 16;

function deriveKeyAndIV(password: Buffer, salt: Buffer) {
  const key = crypto.pbkdf2Sync(
    password,
    salt,
    iterations,
    keyLength,
    "sha256",
  );
  const iv = key.slice(0, ivLength); // Use the first 16 bytes as the IV
  return { key, iv };
}

export function encryptAESGCM(plaintext: Buffer, password: Buffer): Buffer {
  const salt = crypto.randomBytes(16);
  const { key, iv } = deriveKeyAndIV(password, salt);
  const cipher = crypto.createCipheriv(algorithm, key, iv);
  const encryptedData = Buffer.concat([
    cipher.update(plaintext),
    cipher.final(),
  ]);

  return encryptedData;
}

export function decryptAESGCM(ciphertext: Buffer, password: Buffer): Buffer {
  const salt = ciphertext.subarray(0, 16);
  const encryptedData = ciphertext.subarray(16);
  const { key, iv } = deriveKeyAndIV(password, salt);
  const decipher = crypto.createDecipheriv(algorithm, key, iv);

  try {
    const decryptedData = Buffer.concat([
      decipher.update(encryptedData),
      decipher.final(),
    ]);
    return decryptedData;
  } catch (error) {
    const errorMessage =
      error instanceof Error ? error.message : "Unknown error";
    throw g3Error(
      "Decryption failed. Invalid password or corrupted data.",
      errorMessage,
    );
  }
}
