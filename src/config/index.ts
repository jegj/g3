import fs from "fs";
import { z } from "zod";

export const DEFAULT_CONFIG_FILEPATH = "~/.config/g3/config.json";
export const DEFAULT_DATA_FILEPATH = "~/.local/share/g3";

export const G3ConfigSchema = z.object({
  GITHUB_TOKEN: z.string().min(1, "GitHubToken is required"),
  AES_KEY: z.string().min(1, "AesKey is required"),
  NO_DEFAULT_CONFIG_FILE: z.boolean().optional(),
});

export type G3Config = z.infer<typeof G3ConfigSchema>;

export function createG3Config(config: Partial<G3Config>): G3Config {
  return validateG3Config(config);
}

export function parseG3Config(configPath: string): G3Config {
  const content = JSON.parse(fs.readFileSync(configPath, "utf-8"));
  return createG3Config(content);
}

function validateG3Config(config: Partial<G3Config>): G3Config {
  return G3ConfigSchema.parse(config);
}
