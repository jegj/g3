import path from "path";
import { getG3FSEntry } from "../fsdata";
import { FilesystemDataEntry } from "../fsdata/types";
import { G3Dependecies } from "../types";
import { resolvePath } from "../utils";

// A G3File represents a file in the G3 system.
export class G3File {
  g3Filename: string;
  g3Filepath: string;
  filename: string;
  filepath: string;
  filesystemDataEntry: FilesystemDataEntry;
  exists: boolean;

  constructor(fpath: string, dataFolder: string) {
    this.filename = path.basename(fpath);
    this.g3Filename = `${this.filename}.g3.json`;
    this.g3Filepath = path.join(dataFolder, this.g3Filename);
    this.filepath = resolvePath(fpath);
    this.exists = false;
    this.filesystemDataEntry = {
      entries: [],
      createdAt: new Date().toISOString(),
    };
  }

  hasMultipleFiles(): boolean {
    if (this.filesystemDataEntry.entries.length > 1) {
      return true;
    } else if (this.filesystemDataEntry.entries.length === 1) {
      const entry = this.filesystemDataEntry.entries[0];
      return Object.keys(entry.files).length > 1;
    } else {
      return false;
    }
  }
}

export const createG3FileFactory =
  (
    { config }: G3Dependecies,
    getG3Entry: (g3file: G3File) => Promise<FilesystemDataEntry> = getG3FSEntry,
  ) =>
  async (fpath: string): Promise<G3File> => {
    const g3file: G3File = new G3File(fpath, config.DATA_FOLDER);
    try {
      const filesystemDataEntry = await getG3Entry(g3file);
      g3file.filesystemDataEntry = filesystemDataEntry;
      g3file.exists = true;
    } catch (e) {
      // Throw only when the error is different than file not found
      if (
        !(e instanceof Error) ||
        (!e.message.includes("ENOENT") && !e.message.includes("not found"))
      ) {
        throw e;
      }
    }
    return g3file;
  };
