import fs from "fs";
import { z } from "zod";

export const DEFAULT_CONFIG_FILEPATH = "~/.config/g3/config.json";
export const DEFAULT_DATA_FILEPATH = "~/.local/share/g3";

export const G3ConfigSchema = z.object({
  GitHubToken: z.string().min(1, "GitHubToken is required"),
  AesKey: z.string().min(1, "AesKey is required"),
});

export type G3Config = z.infer<typeof G3ConfigSchema>;

export function createG3Config(config: Partial<G3Config>): G3Config {
  validateG3Config(config);
  return {
    GitHubToken: config.GitHubToken!,
    AesKey: config.AesKey!,
  };
}

export function parseG3Config(configPath: string): G3Config {
  const content = JSON.parse(fs.readFileSync(configPath, "utf-8"));
  validateG3Config(content);
  return content as G3Config;
}

function validateG3Config(config: Partial<G3Config>): void {
  G3ConfigSchema.parse(config);
}
