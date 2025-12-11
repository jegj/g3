import { createReadStream, createWriteStream } from "fs";
import { mkdir, readdir, rm, stat } from "fs/promises";
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
  const output = createWriteStream(outputFile);
  const entries = await readdir(searchDir, { withFileTypes: true });
  for (const entry of entries) {
    const fullPath = path.join(searchDir, entry.name);
    if (entry.isDirectory() && entry.name.includes("gist_")) {
      const subEntries = await readdir(fullPath);
      for (const subFile of subEntries) {
        const subFullPath = path.join(fullPath, subFile);
        const stats = await stat(subFullPath);
        if (stats.isFile() && subFile !== ".git") {
          const input = createReadStream(subFullPath);
          await pipeline(input, output, { end: false });
        }
      }
    }
  }
  output.end();
}
