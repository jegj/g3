import { G3Config } from "../config";

export type GistDependecies = {
  config: G3Config;
};

export type GistConfig = {
  description: string;
  public: boolean;
  filenamePrefix: string;
};
export type GistFile = {
  content: string;
};

export type GistCreateRequest = {
  description: string;
  public: boolean;
  files: Record<string, GistFile>;
};

export type GistResponse = {
  id: string;
  html_url: string;
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
    }
  >;
  created_at: string;
  updated_at: string;
};
