import { join } from "path";
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
import { gitClone } from "../git";

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
      //TODO: DO LOGIC FOR MORE FILES
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
