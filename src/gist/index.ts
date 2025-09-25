import { request } from "undici";
import { GistCreateRequest, GistFiles, GistResponse } from "./types";
import { G3Config } from "../config";
import { G3Dependecies } from "../types";

const API_URL = "https://api.github.com/gists";
const DEFAULT_GITHUB_VERSION = "2022-11-28";
const DEFAULT_ACCEPT_HEADER = "application/vnd.github+json";

export const createGistFactory =
  ({ config }: G3Dependecies) =>
  async (
    description: string,
    files: GistFiles,
    isPublic: boolean,
  ): Promise<GistResponse> => {
    const requestData: GistCreateRequest = {
      description,
      public: isPublic,
      files,
    };

    const { statusCode, body } = await request(API_URL, {
      method: "POST",
      headers: getHeaders(config),
      body: JSON.stringify(requestData),
    });

    if (statusCode !== 201) {
      const errorText = await body.text();
      throw new Error(
        `Gist creation failed with status code ${statusCode}: ${errorText}`,
      );
    }

    return body.json() as Promise<GistResponse>;
  };

const getHeaders = (config: G3Config) => ({
  Accept: DEFAULT_ACCEPT_HEADER,
  Authorization: `Bearer ${config.GITHUB_TOKEN}`,
  "X-GitHub-Api-Version": DEFAULT_GITHUB_VERSION,
  "Content-Type": "application/json",
  "User-Agent": "g3-cli",
});
