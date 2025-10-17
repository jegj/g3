export type GistFilesRequest = Record<string, { content: string }>;

export type GistCreateRequest = {
  description: string;
  public: boolean;
  files: GistFilesRequest;
};

export type GistResponse = {
  id: string;
  url: string;
  description: string;
  public: boolean;
  files: Record<string, GistFilesResponse>;
  truncated: boolean;
  created_at: string;
  updated_at: string;
};

export type GistFilesResponse = {
  filename: string;
  type: string;
  language: string;
  raw_url: string;
  size: number;
  content: string;
  truncated: boolean;
  encoding: string;
};
