import fs from "fs/promises";
import { G3File } from "../g3file";
import { G3Dependecies } from "../types";
import { FilesystemDataEntry, GistDataEntry } from "./types";

const G3_FILE_EXTENSION_INDEX_REMOVAL = -8;

export const getFileSizeMb = async (g3File: G3File): Promise<number> => {
  const stats = await fs.stat(g3File.filepath);
  return stats.size / (1024 * 1024);
};

export const getFileContent = async (g3File: G3File): Promise<string> => {
  return fs.readFile(g3File.filepath, { encoding: "utf-8" });
};

export const hasEntry = async (g3FIle: G3File): Promise<boolean> => {
  fs.stat(g3FIle.g3Filepath);
  return true;
};

export const getG3FSEntry = async (
  g3file: G3File,
): Promise<FilesystemDataEntry> => {
  const data = await fs.readFile(g3file.g3Filepath, "utf8");
  const dataEntry: FilesystemDataEntry = JSON.parse(data);
  return dataEntry;
};

export const appendG3FSEntry = async (
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

export const createG3EntriesFactory =
  ({ config }: G3Dependecies) =>
  async (): Promise<string[]> => {
    const files = await fs.readdir(config.DATA_FOLDER, {
      withFileTypes: true,
    });

    if (files.length === 0) {
      return [];
    }

    return files
      .filter((file) => file.isFile() && file.name.endsWith(".g3.json"))
      .map((file) => file.name.slice(0, G3_FILE_EXTENSION_INDEX_REMOVAL));
  };

export const deleteG3Entry = async (g3File: G3File): Promise<void> => {
  await fs.unlink(g3File.g3Filepath);
};
