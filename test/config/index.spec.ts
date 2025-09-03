import { describe, it } from "node:test";
import assert from "node:assert";
import { createG3Config } from "../../src/config";
import { z } from "zod";

describe("Config", () => {
  describe("createG3Config", () => {
    it("should throw an error with correct message when GitHubToken is missing", () => {
      const emptyConfig = {};
      assert.throws(
        () => createG3Config(emptyConfig),
        (error: Error) => {
          assert.ok(error instanceof z.ZodError);
          return true;
        },
      );
    });

    it("should throw an error with correct message when AesKey is missing", () => {
      const partialConfig = {
        GitHubToken: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
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
        GitHubToken: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
        AesKey: "12345678901234567890123456789012",
      };
      const config = createG3Config(partialConfig);
      assert.strictEqual(config.GitHubToken, partialConfig.GitHubToken);
      assert.strictEqual(config.AesKey, partialConfig.AesKey);
    });
  });
});
