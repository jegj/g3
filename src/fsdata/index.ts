import fs from "fs";
import { G3File } from "../g3file";

export const getFileSize = async (g3File: G3File): Promise<number> => {
  const stats = await fs.promises.stat(g3File.filepath);
  return stats.size;
};
