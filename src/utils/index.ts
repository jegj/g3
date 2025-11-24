import { createReadStream, createWriteStream } from "fs";
import { mkdir, readdir, rm } from "fs/promises";
import os from "os";
import path, { join } from "path";
import { pipeline } from "stream/promises";

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

export async function mergeFilesStreaming(
  outputFile: string,
  searchDir: string = ".",
): Promise<void> {
  const files = await readdir(searchDir);
  const matchedFiles = files
    .filter((file) => file !== ".git")
    .sort()
    .map((file) => join(searchDir, file));

  const output = createWriteStream(outputFile);

  for (const file of matchedFiles) {
    const input = createReadStream(file);
    await pipeline(input, output, { end: false });
  }

  output.end();
}
