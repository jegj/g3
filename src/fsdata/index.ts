import fs from "fs/promises";
import { G3File } from "../g3file";
import { GistDataEntry, FilesystemDataEntry } from "./types";

export const getFileSizeMb = async (g3File: G3File): Promise<number> => {
  const stats = await fs.stat(g3File.filepath);
  return stats.size / (1024 * 1024);
};

export const getFileContent = async (g3File: G3File): Promise<Buffer> => {
  return fs.readFile(g3File.filepath);
};

export const hasEntry = async (g3FIle: G3File): Promise<boolean> => {
  fs.stat(g3FIle.g3Filepath);
  return true;
};

export const appendG3Entry = async (
  entries: GistDataEntry[],
  g3file: G3File,
): Promise<void> => {
  const dataEntry: FilesystemDataEntry = {
    entries,
    createdAt: new Date().toISOString(),
  };
  const jsonData = JSON.stringify(dataEntry);
  await fs.writeFile(g3file.g3Filepath, jsonData);
  return;
};
