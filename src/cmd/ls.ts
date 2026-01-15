import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { getG3EntriesFactory } from "../fsdata";
import { G3Dependecies } from "../types";

export default async function ls(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const getG3Entries = getG3EntriesFactory(dependencies);
  const showFullData = argv.full as boolean;

  const entries = await getG3Entries();
  for (const entry of entries) {
    if (showFullData) {
      console.log(
        `${entry.filename} - ${entry.description || "No description"} (${entry.gists?.length || 0} gists)`,
      );
    } else {
      console.log(entry.filename);
    }
  }
}
