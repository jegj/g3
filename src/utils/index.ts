import { mkdir, rm } from "fs/promises";
import os from "os";
import path, { join } from "path";

export function resolvePath(filePath: string): string {
  if (filePath.startsWith("~")) {
    return path.join(os.homedir(), filePath.slice(1));
  }
  return path.resolve(filePath);
}

export async function createTempFolder(
  location: string,
  folderName: string,
): Promise<string> {
  const hiddenFolderName = folderName.startsWith(".")
    ? folderName
    : `.${folderName}`;
  const folderPath = join(location, hiddenFolderName);
  await mkdir(folderPath, { recursive: true });
  return folderPath;
}

export async function deleteFolderIfExists(folder: string): Promise<void> {
  await rm(folder, { recursive: true, force: true });
}
