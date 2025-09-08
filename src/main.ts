#!/usr/bin/env node
import yargs from "yargs";
import { hideBin } from "yargs/helpers";
import {
  createDataFile,
  DEFAULT_CONFIG_FILEPATH,
  parseG3Config,
} from "./config";
import ls from "./cmd/ls";

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
    try {
      const config = parseG3Config(argv.config as string);
      argv.GITHUB_TOKEN = config.GITHUB_TOKEN;
      argv.AES_KEY = config.AES_KEY;
      createDataFile(argv.verbose as boolean);
      if (argv.verbose) {
        console.log("Config loaded successfully");
      }
    } catch (error) {
      if (argv.verbose) {
        console.error("Error loading config:", error);
      }
      console.error(
        `Failed to load config file at ${argv.config}. Check if your configuration file exists and is valid JSON.`,
      );
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
        .usage("$0 cp <source>")
        .example(
          "$0 cp ./myfile.txt",
          "Copy myfile.txt to storage with same name",
        )
        .epilog("The file will be saved to your configured storage location");
    },
    handler: (argv) => {
      console.log(
        `cp command: copying ${argv.source} to ${argv.destination || "default name"}`,
      );
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
    handler: (argv) => {
      console.log(`get command: retrieving ${argv.file}`);
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
    handler: (argv) => {
      console.log(`rm command: removing ${argv.file}`);
    },
  })
  .epilog("For more information, visit https://github.com/jegj/g3")
  .alias("h", "help")
  .usage("$0 <command> [options]")
  .demandCommand(1, "You need at least one command before moving on")
  .strict()
  .parse();
