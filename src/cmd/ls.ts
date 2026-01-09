import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { createG3EntriesFactory } from "../fsdata";
import { G3Dependecies } from "../types";

export default async function ls(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const getG3Entries = createG3EntriesFactory(dependencies);
  const entries = await getG3Entries();
  const showGist = argv.gist as boolean;
  for (const entry of entries) {
    console.log(entry);
  }
}
