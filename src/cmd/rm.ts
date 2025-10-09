import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { deleteG3Entry } from "../fsdata";
import { deleteGistFactory } from "../gist";
import { G3Dependecies } from "../types";
import { createG3FileFactory, G3File } from "../g3file";
import { GistDataEntry } from "../fsdata/types";

export default async function rm(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const deleteG3FileGists = deleteG3FileGistsFactory(dependencies);
  const file = argv.file as string;
  const g3File = await createG3File(file);
  await deleteG3FileGists(g3File);
  await deleteG3Entry(g3File);
}

function deleteG3FileGistsFactory(dependencies: G3Dependecies) {
  const deleteGist = deleteGistFactory(dependencies);
  return async function deleteG3FileGists(g3File: G3File): Promise<void> {
    //TODO: What if fail because one call fail
    await Promise.all(
      g3File.filesystemDataEntry.entries.map((e: GistDataEntry) =>
        deleteGist(e.id),
      ),
    );
  };
}
