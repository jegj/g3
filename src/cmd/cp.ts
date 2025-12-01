import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { createG3FileFactory } from "../g3file";
import { uploadGist } from "../pool";
import { G3Dependecies } from "../types";

export default async function cp(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const description = argv.description as string;
  const file = argv.file as string;
  const g3File = await createG3File(file, description);
  await uploadGist(g3File, config);
}
