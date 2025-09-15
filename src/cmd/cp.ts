import { ArgumentsCamelCase } from "yargs";
import { createConfigFromArgv } from "../config";

export default async function cp(argv: ArgumentsCamelCase) {
  const config = createConfigFromArgv(argv);
  console.log("cp command: listing files in storage", config);
}
