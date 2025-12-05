import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";
import { createG3FileFactory } from "../g3file";
import { uploadFile } from "../pool";
import { G3Dependecies } from "../types";

const chunkSize = 1024 * 1024 * 5; // 5MB default

export default async function cp(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  const dependencies: G3Dependecies = { config };
  const createG3File = createG3FileFactory(dependencies);
  const description = argv.description as string;
  const file = argv.file as string;
  const g3File = await createG3File(file, description);
  await uploadFile(g3File, config, chunkSize);
}
