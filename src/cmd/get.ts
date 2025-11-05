import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { encryptAESGCM } from "../crypto";
import { g3Error } from "../error";
import { appendG3FSEntry, getFileContent, getFileSizeMb } from "../fsdata";
import { GistDataEntry } from "../fsdata/types";
import { createG3FileFactory } from "../g3file";
import { createGistFactory } from "../gist";
import { GistResponse } from "../gist/types";
import { G3Dependecies } from "../types";

// The idea to use the gist_pull_url and do the merging locally
export default async function get(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const file = argv.file as string;
  const g3File = await createG3File(file);

  if (g3File.exists) {
    if (g3File.hasMultipleFiles()) {
    } else {
      //TODO: CHECK IF USER WORKER OR PROCESS
    }
  } else {
    throw g3Error(`File ${file} does not exist in g3data.`);
  }
}
