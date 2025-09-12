import fs from "fs";
import path from "path";
import os from "os";
import { z } from "zod";

const HOME_DIR = os.homedir();

export const DEFAULT_CONFIG_FILEPATH = path.join(
  HOME_DIR,
  ".config/g3/config.json",
);
export const DEFAULT_DATA_FILEPATH = path.join(
  HOME_DIR,
  ".local/share/g3/files/",
);

export const G3ConfigSchema = z.object({
  GITHUB_TOKEN: z.string().min(1, "GitHubToken is required"),
  AES_KEY: z.string().min(1, "AesKey is required"),
  DATA_FOLDER: z
    .string()
    .transform((val) => (val?.trim() === "" ? undefined : val))
    .default(DEFAULT_DATA_FILEPATH),
});

export type G3Config = z.infer<typeof G3ConfigSchema>;

export function createG3Config(config: Partial<G3Config>): G3Config {
  return validateG3Config(config);
}

export function parseG3Config(configPath: string): G3Config {
  const resolvedConfigPath = resolvePath(configPath);
  const content = JSON.parse(fs.readFileSync(resolvedConfigPath, "utf-8"));
  return createG3Config(content);
}

function resolvePath(filePath: string): string {
  if (filePath.startsWith("~")) {
    return path.join(os.homedir(), filePath.slice(1));
  }
  return path.resolve(filePath);
}

function validateG3Config(config: Partial<G3Config>): G3Config {
  return G3ConfigSchema.parse(config);
}

export function createDataFile(dataFolder: string, verbose: boolean = false) {
  const dir = resolvePath(path.dirname(dataFolder));
  if (!fs.existsSync(dir)) {
    if (verbose) {
      console.log(`Creating data directory at ${dir}`);
    }
    fs.mkdirSync(dir, { recursive: true });
  }
}

export function createConfigFromArgv<T extends object>(source: T): G3Config {
  return createG3Config(source);
}
