import fs from "fs/promises";
import { G3File, parseG3FileFactory } from "../g3file";
import { G3Dependecies } from "../types";
import { FilesystemDataEntry, GistDataEntry } from "./types";

export const getFileContent = async (g3File: G3File): Promise<string> => {
  return fs.readFile(g3File.filepath, { encoding: "utf-8" });
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
    description: g3file.description,
  };
  const jsonData = JSON.stringify(dataEntry);
  await fs.writeFile(g3file.g3Filepath, jsonData);
  return;
};

export const getG3EntriesFactory =
  ({ config }: G3Dependecies) =>
  async (): Promise<G3File[]> => {
    const dependencies: G3Dependecies = { config };
    const parseG3File = parseG3FileFactory(dependencies);

    const files = await fs.readdir(config.DATA_FOLDER, {
      withFileTypes: true,
    });

    if (files.length === 0) {
      return [];
    }

    return await Promise.all(
      files
        .filter((file) => file.isFile() && file.name.endsWith(".g3.json"))
        .map((file) => file.name.replace(/\.g3\.json$/, ""))
        .map((filename) => parseG3File(filename)),
    );
  };

export const deleteG3Entry = async (g3File: G3File): Promise<void> => {
  await fs.unlink(g3File.g3Filepath);
};
