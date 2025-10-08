import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { deleteG3Entry } from "../fsdata";
import { deleteGistFactory } from "../gist";
import { G3Dependecies } from "../types";
import { createG3FileFactory } from "../g3file";
import { GistDataEntry } from "../fsdata/types";

export default async function rm(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const deleteGist = deleteGistFactory(dependencies);
  const createG3File = createG3FileFactory(dependencies);
  const file = argv.file as string;
  const g3File = await createG3File(file);

  //TODO: What if fail because one call fail
  await Promise.all(
    g3File.filesystemDataEntry.entries.map((e: GistDataEntry) =>
      deleteGist(e.id),
    ),
  );
  await deleteG3Entry(g3File);
}
