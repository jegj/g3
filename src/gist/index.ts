import { request } from "undici";
import {
  GistCreateRequest,
  GistDependecies,
  GistFile,
  GistResponse,
} from "./types";
import { G3Config } from "../config";

const API_URL = "https://api.github.com/gists";
const DEFAULT_GITHUB_VERSION = "2022-11-28";
const DEFAULT_ACCEPT_HEADER = "application/vnd.github+json";

export const createGist =
  ({ config }: GistDependecies) =>
  async (
    description: string,
    files: Record<string, GistFile>,
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
});

/*
const createGistService = (cfg: Config): GistProvider => {
  const getHeaders = () => ({
    Accept: DEFAULT_ACCEPT_HEADER,
    Authorization: `Bearer ${cfg.ghToken}`,
    "X-GitHub-Api-Version": DEFAULT_GITHUB_VERSION,
    "Content-Type": "application/json",
  });

  const createGist = async (
    description: string,
    files: Record<string, GistFile>,
    isPublic: boolean,
  ): Promise<GistResponse> => {
    const requestData: GistCreateRequest = {
      description,
      public: isPublic,
      files,
    };

    const { statusCode, body } = await request(API_URL, {
      method: "POST",
      headers: getHeaders(),
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

  const deleteGist = async (id: string): Promise<void> => {
    const url = `${API_URL}/${id}`;

    const { statusCode, body } = await request(url, {
      method: "DELETE",
      headers: getHeaders(),
    });

    if (statusCode !== 204) {
      const errorText = await body.text();
      throw new Error(
        `Gist deletion failed with status code ${statusCode}: ${errorText}`,
      );
    }
  };

  const getGist = async (id: string): Promise<GistResponse> => {
    const url = `${API_URL}/${id}`;

    const { statusCode, body } = await request(url, {
      method: "GET",
      headers: getHeaders(),
    });

    if (statusCode !== 200) {
      const errorText = await body.text();
      throw new Error(
        `Gist retrieval failed with status code ${statusCode}: ${errorText}`,
      );
    }

    return body.json() as Promise<GistResponse>;
  };

  const updateGist = async (
    id: string,
    description: string,
    files: Record<string, GistFile>,
    isPublic: boolean,
  ): Promise<GistResponse> => {
    const url = `${API_URL}/${id}`;
    const requestData: GistCreateRequest = {
      description,
      public: isPublic,
      files,
    };

    const { statusCode, body } = await request(url, {
      method: "PATCH",
      headers: getHeaders(),
      body: JSON.stringify(requestData),
    });

    if (statusCode !== 200) {
      const errorText = await body.text();
      throw new Error(
        `Gist update failed with status code ${statusCode}: ${errorText}`,
      );
    }

    return body.json() as Promise<GistResponse>;
  };

  return {
    createGist,
    deleteGist,
    getGist,
    updateGist,
  };
};

export { createGistService, GistProvider, Config, GistResponse, GistFile };
*/
