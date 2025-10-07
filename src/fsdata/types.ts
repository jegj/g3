export type FilesystemDataEntry = {
  entries: GistDataEntry[];
  createdAt: string;
};

export type GistDataEntry = {
  id: string;
  gistUrl: string;
};
