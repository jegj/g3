import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { getG3EntriesFactory } from "../fsdata";
import { G3File } from "../g3file";
import { G3Dependecies } from "../types";

export default async function ls(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const getG3Entries = getG3EntriesFactory(dependencies);
  const showFullData = argv.full as boolean;

  const entries: G3File[] = await getG3Entries();
  for (const entry of entries) {
    if (showFullData) {
      console.log(
        `${entry.filename} - ${entry.description} (${entry.gistCount} gists)`,
      );
    } else {
      console.log(entry.filename);
    }
  }
}
