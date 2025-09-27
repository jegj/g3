export type GistFiles = Record<string, { content: string }>;

export type GistCreateRequest = {
  description: string;
  public: boolean;
  files: GistFiles;
};

export type GistResponse = {
  id: string;
  url: string;
  description: string;
  public: boolean;
  files: Record<
    string,
    {
      filename: string;
      type: string;
      language: string;
      raw_url: string;
      size: number;
      content: string;
      truncated: boolean;
    }
  >;
  truncated: boolean;
  created_at: string;
  updated_at: string;
};
