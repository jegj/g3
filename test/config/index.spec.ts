import { describe, test } from "node:test";
import assert from "node:assert";
import { validateG3Config } from "../../src/config";

describe("Config", () => {
  test("test", () => {
    assert.strictEqual(1, 1);
  });

  describe("validateG3Config", () => {
    test("should throw if GitHubToken is missing", () => {
      assert.throws(() => validateG3Config({}), {
        name: "Error",
        message: "GithubToken is required",
      });
    });

    test("should throw if GitHubToken is missing", () => {
      assert.throws(
        () =>
          validateG3Config({
            GitHubToken: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
          }),
        {
          name: "Error",
          message: "AesKey is required",
        },
      );
    });
  });
});
