#!/usr/bin/env node
import kleur from "kleur";
import yargs from "yargs";
import { hideBin } from "yargs/helpers";
import cp from "./cmd/cp";
import get from "./cmd/get";
import ls from "./cmd/ls";
import rm from "./cmd/rm";
import {
  createDataFile,
  DEFAULT_CONFIG_FILEPATH,
  parseG3Config,
} from "./config";

const argv = hideBin(process.argv);

yargs(argv)
  .scriptName("g3")
  .version()
  .help()
  .option("verbose", {
    alias: "v",
    type: "boolean",
    description: "Show debug logging",
    default: false,
    global: true,
  })
  .option("config", {
    alias: "c",
    type: "string",
    description: "Path to config file",
    default: DEFAULT_CONFIG_FILEPATH,
    global: true,
  })
  .middleware((argv) => {
    const configPath = argv.config as string;
    try {
      const config = parseG3Config(configPath);
      Object.assign(argv, config);
      createDataFile(config.DATA_FOLDER, Boolean(argv.verbose));
      if (argv.verbose) {
        console.log("Config loaded successfully");
      }
    } catch (error) {
      console.error(
        kleur.red(
          `Failed to load config file at ${configPath}. Check if your configuration file exists and it is a valid JSON.`,
        ),
      );
      if (argv.verbose) {
        console.error(
          kleur.dim(error instanceof Error ? error.stack || "" : String(error)),
        );
      }
      process.exit(1);
    }
  })
  .command({
    command: "ls",
    describe: "Show all the files in your storage",
    builder: (yargs) => {
      return yargs
        .usage("$0 ls")
        .example("$0 ls", "List all files in your storage")
        .epilog("Lists all files stored in your configured storage location");
    },
    handler: async (argv) => {
      await ls(argv);
    },
  })
  .command({
    command: "cp <file>",
    describe:
      "Copy a file to your storage. If the file exists, it will be overwritten.",
    builder: (yargs) => {
      return yargs
        .positional("file", { describe: "Source file path", type: "string" })
        .option("description", {
          alias: "d",
          type: "string",
          description: "A brief description of your file",
        })
        .usage("$0 cp <source>")
        .example(
          // eslint-disable-next-line quotes
          '$0 cp ./myfile.txt -d "My importan file"',
          "Copy myfile.txt to storage with same name",
        )
        .epilog("The file will be saved to your configured storage location");
    },
    handler: async (argv) => {
      await cp(argv);
    },
  })
  .command({
    command: "get <file>",
    describe: "Get file from your storage",
    builder: (yargs) => {
      return yargs
        .positional("file", {
          describe: "File to retrieve",
          type: "string",
        })
        .usage("$0 get <file>")
        .example("$0 get myfile.txt", "Retrieve myfile.txt from storage")
        .epilog(
          "The file will be retrieved from your configured storage location",
        );
    },
    handler: async (argv) => {
      await get(argv);
    },
  })
  .command({
    command: "rm <file>",
    describe: "Delete file from your storage",
    builder: (yargs) => {
      return yargs
        .positional("file", {
          describe: "File to remove",
          type: "string",
        })
        .usage("$0 rm <file>")
        .example("$0 rm myfile.txt", "Remove myfile.txt from storage")
        .epilog(
          "The file will be removed from your configured storage location",
        );
    },
    handler: async (argv) => {
      await rm(argv);
    },
  })
  .fail((msg, err, yargs) => {
    if ("isOperational" in err && err.isOperational) {
      console.error(kleur.red(`Error: ${err.message}`));
    } else {
      console.error(kleur.red("Fatal Error:"), err.message);
      console.error(kleur.dim(err.stack || ""));
    }
    process.exit(1);
  })
  .epilog("For more information, visit https://github.com/jegj/g3")
  .alias("h", "help")
  .usage("$0 <command> [options]")
  .demandCommand(1, "You need at least one command before moving on")
  .strict()
  .parse();
