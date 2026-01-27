import assert from "node:assert";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import { resolvePath } from "../../src/utils";

describe("Utils", () => {
  describe("resolvePath", () => {
    it("should resolve path starting with ~ to home directory", () => {
      const result = resolvePath("~/documents/file.txt");
      const expected = path.join(os.homedir(), "documents/file.txt");
      assert.strictEqual(result, expected);
    });

    it("should resolve ~ alone to home directory", () => {
      const result = resolvePath("~");
      const expected = os.homedir();
      assert.strictEqual(result, expected);
    });

    it("should resolve relative path to absolute path", () => {
      const result = resolvePath("./file.txt");
      const expected = path.resolve("./file.txt");
      assert.strictEqual(result, expected);
    });

    it("should return absolute path unchanged", () => {
      const absolutePath = "/home/user/documents/file.txt";
      const result = resolvePath(absolutePath);
      assert.strictEqual(result, absolutePath);
    });

    it("should resolve empty string to current directory", () => {
      const result = resolvePath("");
      const expected = path.resolve("");
      assert.strictEqual(result, expected);
    });

    it("should resolve path with parent directory references", () => {
      const result = resolvePath("../parent/file.txt");
      const expected = path.resolve("../parent/file.txt");
      assert.strictEqual(result, expected);
    });
  });
});
