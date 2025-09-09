import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";

export default async function ls(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  console.log("ls command: listing files in storage", config);
}
