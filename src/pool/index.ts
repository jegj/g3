import * as fs from "fs/promises";
import * as os from "os";
import * as path from "path";
import { resolve } from "path";
import Piscina from "piscina";
import { G3Config } from "../config";
import { appendG3FSEntry } from "../fsdata";
import { GistDataEntry } from "../fsdata/types";
import { G3File } from "../g3file";
import { createTempFolder } from "../utils";
import { filename, ProcessFileChunkParam } from "./worker";

const piscina = new Piscina({
  filename: resolve(__dirname, "./workerWrapper.js"),
  minThreads: 0,
  maxThreads: os.cpus().length * 4,
  concurrentTasksPerWorker: 1,
  idleTimeout: 1000,
  workerData: { fullpath: filename },
});

export async function decryptFilesInFolder(
  folder: string,
  descryptPassword: string,
) {
  const entries = await fs.readdir(folder, { withFileTypes: true });
  for (const entry of entries) {
    const fullPath = path.join(folder, entry.name);
    if (entry.isDirectory() && entry.name.includes("gist_")) {
      const subEntries = await fs.readdir(fullPath);
      for (const subFile of subEntries) {
        const subFullPath = path.join(fullPath, subFile);
        const stats = await fs.stat(subFullPath);
        if (stats.isFile()) {
          await piscina.run(
            { file: subFullPath, password: descryptPassword },
            { name: "decryptFile" },
          );
        }
      }
    }
  }
}

export async function uploadFile(
  g3File: G3File,
  config: G3Config,
): Promise<void> {
  const chunkSize = config.CHUNK_SIZE;
  const stats = await fs.stat(g3File.filepath);
  const fileSize = stats.size;
  const numChunks = Math.ceil(fileSize / chunkSize);
  const tasks: Promise<GistDataEntry>[] = [];
  for (let i = 0; i < numChunks; i++) {
    const start = i * chunkSize;
    const end = Math.min(start + chunkSize, fileSize);
    const task: ProcessFileChunkParam = {
      filePath: g3File.filepath,
      sortableFileName: g3File.sortableFileName,
      start,
      end,
      chunkIndex: i,
      g3File,
      config,
    };
    tasks.push(
      piscina.run(
        {
          ...task,
        },
        { name: "processGistChunk" },
      ),
    );
  }
  const entries = await Promise.all(tasks);
  await appendG3FSEntry(entries, g3File);
}

export async function cloneGistEntries(
  entries: GistDataEntry[],
  temporalFolder: string,
  token: string,
): Promise<void> {
  let index = 1;
  const tasks: Promise<void>[] = [];
  for (const entry of entries) {
    const paddedIndex = index.toString().padStart(3, "0");
    const subTemporalFolder = await createTempFolder(
      temporalFolder,
      `${paddedIndex}_gist_${entry.id}`,
    );
    tasks.push(
      piscina.run(
        {
          gistPullUrl: entry.gistPullUrl,
          subTemporalFolder,
          token,
        },
        { name: "cloneRepo" },
      ),
    );
    index++;
  }
  await Promise.all(tasks);
}
