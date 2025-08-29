#!/usr/bin/env node
import yargs from "yargs";
import { hideBin } from "yargs/helpers";

const argv = hideBin(process.argv);

yargs(argv)
  .scriptName("g3")
  .version()
  .help()
  .command({
    command: "ls",
    describe: "Show all the files in your storage",
    handler: () => {
      console.log("ls command");
    },
  })
  .command({
    command: "cp <source> [destination]",
    describe: "Copy a file to your storage",
    builder: (yargs) => {
      return yargs
        .positional("source", { describe: "Source file path", type: "string" })
        .positional("destination", {
          describe: "Destination name (optional)",
          type: "string",
        });
    },
    handler: (argv) => {
      console.log(
        `cp command: copying ${argv.source} to ${argv.destination || "default name"}`,
      );
    },
  })
  .command({
    command: "get <file>",
    describe: "Retrieve a file from your storage",
    builder: (yargs) => {
      return yargs.positional("file", {
        describe: "File to retrieve",
        type: "string",
      });
    },
    handler: (argv) => {
      console.log(`get command: retrieving ${argv.file}`);
    },
  })
  .command({
    command: "rm <file>",
    describe: "Remove a file from your storage",
    builder: (yargs) => {
      return yargs.positional("file", {
        describe: "File to remove",
        type: "string",
      });
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
