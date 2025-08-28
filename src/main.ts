#!/usr/bin/env node
import yargs from "yargs";

const argv = yargs(process.argv.slice(2))
  .options({
    a: { type: "boolean", default: false },
    b: { type: "string", demandOption: true },
    c: { type: "number", alias: "chill" },
    d: { type: "array" },
    e: { type: "count" },
    f: { choices: ["1", "2", "3"] },
  })
  .parse();

const myFunction = (name: string, age: number) => {
  console.log("Hello, " + name + "! You are " + age + " years old.");

  let unusedVariable = 42; // Unused variable (ESLint should warn about this)

  return { name: name, age: age };
};

console.log(myFunction("Alice", 25)); //
