import * as crypto from "crypto";

/**
 * Encrypts data using AES-GCM algorithm
 * @param plaintext The data to encrypt
 * @param key The encryption key (should be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
 * @returns Encrypted data with prepended nonce
 */
export function encryptAESGCM(plaintext: Buffer, key: Buffer): Buffer {
  // Create initialization vector (nonce)
  const nonce = crypto.randomBytes(12); // 12 bytes for GCM

  // Create cipher
  const cipher = crypto.createCipheriv("aes-256-gcm", key, nonce);

  // Encrypt data
  const encrypted = Buffer.concat([cipher.update(plaintext), cipher.final()]);

  // Get authentication tag
  const authTag = cipher.getAuthTag();

  // Return nonce + encrypted data + authentication tag
  return Buffer.concat([nonce, encrypted, authTag]);
}

/**
 * Decrypts data that was encrypted using AES-GCM
 * @param ciphertext The encrypted data (with prepended nonce and appended auth tag)
 * @param key The encryption key used for encryption
 * @returns Decrypted data
 */
export function decryptAESGCM(ciphertext: Buffer, key: Buffer): Buffer {
  // Extract nonce (first 12 bytes)
  const nonce = ciphertext.subarray(0, 12);

  // Extract authentication tag (last 16 bytes)
  const authTag = ciphertext.subarray(ciphertext.length - 16);

  // Extract the actual encrypted content (everything between nonce and authTag)
  const encryptedContent = ciphertext.subarray(12, ciphertext.length - 16);

  // Create decipher
  const decipher = crypto.createDecipheriv("aes-256-gcm", key, nonce);

  // Set authentication tag
  decipher.setAuthTag(authTag);

  // Decrypt data
  return Buffer.concat([decipher.update(encryptedContent), decipher.final()]);
}
