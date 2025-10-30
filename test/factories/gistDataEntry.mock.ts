import { GistDataEntry } from "../../src/fsdata/types";
import { GistFilesResponse } from "../../src/gist/types";

export function generateMockGistDataEntry(
  numFiles?: number,
  overrides?: Partial<GistDataEntry>,
): GistDataEntry {
  const randomId = Math.random().toString(36).substring(2, 15);
  const username = `user${Math.floor(Math.random() * 1000)}`;

  const fileTypes = [
    { ext: "ts", type: "application/typescript", language: "TypeScript" },
    { ext: "js", type: "application/javascript", language: "JavaScript" },
    { ext: "md", type: "text/markdown", language: "Markdown" },
    { ext: "json", type: "application/json", language: "JSON" },
    { ext: "py", type: "text/x-python", language: "Python" },
    { ext: "txt", type: "text/plain", language: "Text" },
  ];

  const validNumFoles = numFiles ? numFiles : Math.floor(Math.random() * 3) + 1; // 1-3 files
  const files: Record<string, GistFilesResponse> = {};

  for (let i = 0; i < validNumFoles; i++) {
    const fileType = fileTypes[Math.floor(Math.random() * fileTypes.length)];
    const filename = `file${i + 1}.${fileType.ext}`;
    const content = `Sample content for ${filename}\nLine 2\nLine 3`;

    files[filename] = {
      filename,
      type: fileType.type,
      language: fileType.language,
      raw_url: `https://gist.githubusercontent.com/${username}/${randomId}/raw/${filename}`,
      size: content.length,
      truncated: false,
      content,
      encoding: "utf-8",
    };
  }

  return {
    id: randomId,
    gistUrl: `https://gist.github.com/${username}/${randomId}`,
    files,
    ...overrides,
  };
}
