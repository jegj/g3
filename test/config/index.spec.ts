import assert from "node:assert";
import fs from "node:fs";
import path from "node:path";
import { afterEach, beforeEach, describe, it } from "node:test";
import { z } from "zod";
import { createG3Config, parseG3Config } from "../../src/config";

describe("Config", () => {
  describe("createG3Config", () => {
    it("should throw an error with correct message when GITHUB_TOKEN is missing", () => {
      const emptyConfig = {};
      assert.throws(
        () => createG3Config(emptyConfig),
        (error: Error) => {
          assert.ok(error instanceof z.ZodError);
          return true;
        },
      );
    });

    it("should throw an error with correct message when AES_KEY is missing", () => {
      const partialConfig = {
        GITHUB_TOKEN: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
      };

      assert.throws(
        () => createG3Config(partialConfig),
        (error: Error) => {
          assert.ok(error instanceof z.ZodError);
          return true;
        },
      );
    });

    it("should return a G3Config type if the config properties are valid", () => {
      const partialConfig = {
        GITHUB_TOKEN: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
        AES_KEY: "12345678901234567890123456789012",
      };
      const config = createG3Config(partialConfig);
      assert.strictEqual(config.GITHUB_TOKEN, partialConfig.GITHUB_TOKEN);
      assert.strictEqual(config.AES_KEY, partialConfig.AES_KEY);
    });
  });

  describe("parseG3Config", () => {
    const validConfig = {
      GITHUB_TOKEN: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
      AES_KEY: "12345678901234567890123456789012",
    };

    const tempConfigPath = path.join(__dirname, "temp-config.json");

    beforeEach(() => {
      fs.writeFileSync(tempConfigPath, JSON.stringify(validConfig));
    });

    afterEach(() => {
      if (fs.existsSync(tempConfigPath)) {
        fs.unlinkSync(tempConfigPath);
      }
    });

    it("should throw an error when the file doesn't exist", () => {
      const nonExistentPath = path.join(__dirname, "non-existent-config.json");

      assert.throws(() => parseG3Config(nonExistentPath), {
        name: "Error",
        message: /ENOENT: no such file or directory/,
      });
    });

    it("should throw an error when the file contains invalid JSON", () => {
      fs.writeFileSync(tempConfigPath, "{ invalid json: content }");

      assert.throws(() => parseG3Config(tempConfigPath), {
        name: "SyntaxError",
      });
    });

    it("should parse a valid config file correctly", () => {
      const config = parseG3Config(tempConfigPath);

      assert.strictEqual(config.GITHUB_TOKEN, validConfig.GITHUB_TOKEN);
      assert.strictEqual(config.AES_KEY, validConfig.AES_KEY);
    });
  });
});
