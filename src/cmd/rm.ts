import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { g3Error } from "../error";
import { deleteG3Entry } from "../fsdata";
import { GistDataEntry } from "../fsdata/types";
import { createG3FileFactory, G3File } from "../g3file";
import { deleteGistFactory } from "../gist";
import { G3Dependecies } from "../types";

export default async function rm(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const deleteGist = deleteGistFactory(dependencies);
  const file = argv.file as string;
  const g3File = await createG3File(file);
  if (g3File.exists) {
    await Promise.all(g3File.gists.map((e: GistDataEntry) => deleteGist(e.id)));
    await deleteG3Entry(g3File);
  } else {
    throw g3Error(`File ${file} does not exist in g3data.`);
  }
}
