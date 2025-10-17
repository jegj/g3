import { GistFilesResponse } from "../gist/types";

export type FilesystemDataEntry = {
  entries: GistDataEntry[];
  createdAt: string;
};

export type GistDataEntry = {
  id: string;
  gistUrl: string;
  files: Record<string, GistFilesResponse>;
};
