import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { getG3EntriesFactory } from "../fsdata";
import { G3Dependecies } from "../types";

export default async function ls(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const getG3Entries = getG3EntriesFactory(dependencies);
  const showGist = argv.gist as boolean;
  //TODO: Add gist count and description display
  const entries = await getG3Entries();
  for (const entry of entries) {
    console.log(entry.filename);
  }
}
