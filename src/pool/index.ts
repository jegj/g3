import * as fs from "fs/promises";
import * as os from "os";
import * as path from "path";
import { resolve } from "path";
import Piscina from "piscina";
import { filename } from "./worker";

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
