import * as fs from "fs/promises";
import * as os from "os";
import * as path from "path";
import { resolve } from "path";
import Piscina from "piscina";
import { G3Config } from "../config";
import { G3File } from "../g3file";
import { filename, ProcessFileChunkParam } from "./worker";

const piscina = new Piscina({
  filename: resolve(__dirname, "./workerWrapper.js"),
  minThreads: 0,
  maxThreads: os.cpus().length * 4,
  concurrentTasksPerWorker: 1,
  idleTimeout: 1000,
  workerData: { fullpath: filename },
});

piscina.on("message", (event) => {
  console.log("Message received from worker: ", event);
});

export async function decryptFilesInFolder(
  folder: string,
  descryptPassword: string,
) {
  const files = await fs.readdir(folder);
  for (const file of files) {
    const fullPath = path.join(folder, file);
    const stats = await fs.stat(fullPath);
    if (stats.isFile()) {
      await piscina.run(
        { file: fullPath, password: descryptPassword },
        { name: "decryptFile" },
      );
    }
  }
}

export async function uploadGist(
  g3File: G3File,
  config: G3Config,
  chunkSize: number = 1024 * 1024, // 1MB default
) {
  const stats = await fs.stat(g3File.filepath);
  const fileSize = stats.size;
  const numChunks = Math.ceil(fileSize / chunkSize);
  console.log(`File size: ${fileSize} bytes`);
  console.log(`Reading in ${numChunks} chunks ...`);
  const tasks: Promise<void>[] = [];
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
  console.log(`Uploading chunks in parallel ...`);
  await Promise.all(tasks);
}
