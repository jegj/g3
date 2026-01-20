import path from "path";
import { getG3FSEntry } from "../fsdata";
import { FilesystemDataEntry, GistDataEntry } from "../fsdata/types";
import { G3Dependecies } from "../types";
import { resolvePath } from "../utils";

const ZERO_PAD = 3;
const NO_DESCRIPTION = "No description";

// A G3File represents a file in the G3 system.
export class G3File {
  g3Filename: string;
  g3Filepath: string;
  filename: string;
  filepath: string;
  filesystemDataEntry: FilesystemDataEntry;
  exists: boolean;
  private findex: number;

  constructor(fpath: string, dataFolder: string, description: string) {
    this.filename = path.basename(fpath);
    this.g3Filename = `${this.filename}.g3.json`;
    this.g3Filepath = path.join(dataFolder, this.g3Filename);
    this.filepath = resolvePath(fpath);
    this.exists = false;
    this.findex = 0;
    this.filesystemDataEntry = {
      entries: [],
      createdAt: new Date().toISOString(),
      description,
    };
  }

  get sortableFileName(): string {
    return `${(this.findex++).toString().padStart(ZERO_PAD, "0")}-${this.filename}`;
  }

  get gists(): GistDataEntry[] {
    return this.filesystemDataEntry.entries;
  }

  get gistCount(): number {
    return this.filesystemDataEntry.entries.length;
  }

  get description(): string {
    return this.filesystemDataEntry.description || NO_DESCRIPTION;
  }
}

export const createG3FileFactory =
  (
    { config }: G3Dependecies,
    getG3Entry: (g3file: G3File) => Promise<FilesystemDataEntry> = getG3FSEntry,
  ) =>
  async (fpath: string, description: string = ""): Promise<G3File> => {
    const g3file: G3File = new G3File(fpath, config.DATA_FOLDER, description);
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

export const parseG3FileFactory =
  (
    { config }: G3Dependecies,
    getG3Entry: (g3file: G3File) => Promise<FilesystemDataEntry> = getG3FSEntry,
  ) =>
  async (dataFilePath: string): Promise<G3File> => {
    const fpath = path.basename(dataFilePath);
    const g3file = new G3File(fpath, config.DATA_FOLDER, "");
    const filesystemDataEntry = await getG3Entry(g3file);
    g3file.filesystemDataEntry = filesystemDataEntry;
    g3file.exists = true;
    return g3file;
  };
