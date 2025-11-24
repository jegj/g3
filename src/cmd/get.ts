import { spawn } from "child_process";
import { createReadStream, createWriteStream } from "fs";
import { readdir } from "fs/promises";
import { join } from "path";
import { pipeline } from "stream/promises";
import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { g3Error } from "../error";
import { createG3FileFactory } from "../g3file";
import { decryptFilesInFolder } from "../pool";
import { G3Dependecies } from "../types";
import {
  createTempFolder,
  deleteFolderIfExists,
  mergeFilesStreaming,
  resolvePath,
} from "../utils";

export default async function get(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const file = argv.file as string;
  const fileDestination = resolvePath(argv.destination as string);
  const g3File = await createG3File(file);
  const temporalFolder = await createTempFolder(
    fileDestination,
    g3File.filename,
  );
  const finalFilePath = join(fileDestination, g3File.filename);
  console.log(`Getting file ${file} to ${temporalFolder}/`);
  if (g3File.exists) {
    if (g3File.hasMultipleGistEntries()) {
    } else {
      await deleteFolderIfExists(temporalFolder);
      await gitClone(
        g3File.filesystemDataEntry.entries[0].gistPullUrl,
        temporalFolder,
      );
      await decryptFilesInFolder(temporalFolder, config.AES_KEY);
      await mergeFilesStreaming(finalFilePath, temporalFolder);
    }
  } else {
    throw g3Error(`File ${file} does not exist in g3data.`);
  }
}

async function gitClone(url: string, folder: string): Promise<void> {
  return new Promise((resolve, reject) => {
    const gitProcess = spawn("git", ["clone", url, folder]);

    let stderr = "";

    gitProcess.stdout.on("data", (data) => {
      console.log(data.toString());
    });

    gitProcess.stderr.on("data", (data) => {
      stderr += data.toString();
      console.error(data.toString());
    });

    gitProcess.on("close", (code) => {
      if (code === 0) {
        resolve();
      } else {
        reject(new Error(`git clone failed with code ${code}: ${stderr}`));
      }
    });

    gitProcess.on("error", (error) => {
      reject(new Error(`Failed to start git process: ${error.message}`));
    });
  });
}
