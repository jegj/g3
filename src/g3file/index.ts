import path from "path";
import { G3Dependecies } from "../types";
import { resolvePath } from "../utils";

// A G3File represents a file in the G3 system.
export type G3File = {
  g3Filename: string;
  g3Filepath: string;
  filename: string;
  filepath: string;
};

export const createG3FileFactory =
  ({ config }: G3Dependecies) =>
  (fpath: string): G3File => {
    const filename = path.basename(fpath);
    const g3Filename = `${filename}.g3.json`;
    const g3Filepath = path.join(config.DATA_FOLDER, g3Filename);
    const filepath = resolvePath(fpath);

    return {
      g3Filename,
      g3Filepath,
      filename: filename,
      filepath,
    };
  };
