export type GistDataEntry = {
  id: string;
  gistUrl: string;
};

export type FilesystemDataEntry = {
  entries: GistDataEntry[];
  createdAt: string;
};
