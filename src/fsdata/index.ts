import fs from "fs";
import { stat } from "fs/promises";
import { G3File } from "../g3file";

export const getFileSizeMb = async (g3File: G3File): Promise<number> => {
  const stats = await fs.promises.stat(g3File.filepath);
  return stats.size / (1024 * 1024);
};

// For small files
export const getFileContent = async (g3File: G3File): Promise<Buffer> => {
  return fs.promises.readFile(g3File.filepath);
};

export const hasEntry = async (g3FIle: G3File): Promise<boolean> => {
  stat(g3FIle.g3Filepath);
  return true;
};
