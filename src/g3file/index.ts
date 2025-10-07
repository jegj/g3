import path from "path";
import { G3Dependecies } from "../types";
import { resolvePath } from "../utils";
import { getG3Entry } from "../fsdata";
import { FilesystemDataEntry } from "../fsdata/types";

// A G3File represents a file in the G3 system.
export type G3File = {
  g3Filename: string;
  g3Filepath: string;
  filename: string;
  filepath: string;
  filesystemDataEntry?: FilesystemDataEntry | null;
};

export const createG3FileFactory =
  ({ config }: G3Dependecies) =>
  async (fpath: string): Promise<G3File> => {
    const filename = path.basename(fpath);
    const g3Filename = `${filename}.g3.json`;
    const g3Filepath = path.join(config.DATA_FOLDER, g3Filename);
    const filepath = resolvePath(fpath);

    const g3file: G3File = {
      g3Filename,
      g3Filepath,
      filename: filename,
      filepath,
    };

    try {
      const filesystemDataEntry = await getG3Entry(g3file);
      g3file.filesystemDataEntry = filesystemDataEntry;
    } catch (e) {
      if (
        e instanceof Error &&
        (e.message.includes("ENOENT") || e.message.includes("not found"))
      ) {
        g3file.filesystemDataEntry = null;
      } else {
        throw e;
      }
    }
    return g3file;
  };
