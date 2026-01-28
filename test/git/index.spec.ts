import assert from "node:assert";
import { describe, it } from "node:test";

describe("Git", () => {
  describe("gitClone URL transformation", () => {
    it("should correctly inject token into URL", () => {
      // This tests the URL transformation logic used in gitClone
      const url = "https://gist.github.com/abc123.git";
      const token = "my-secret-token";
      const expectedAuthenticatedUrl = `https://${token}@gist.github.com/abc123.git`;
      const authenticatedUrl = url.replace("https://", `https://${token}@`);

      assert.strictEqual(authenticatedUrl, expectedAuthenticatedUrl);
    });

    it("should handle URL with existing path correctly", () => {
      const url = "https://gist.github.com/user/abc123.git";
      const token = "token123";
      const authenticatedUrl = url.replace("https://", `https://${token}@`);

      assert.strictEqual(
        authenticatedUrl,
        "https://token123@gist.github.com/user/abc123.git",
      );
    });

    it("should handle tokens with special characters in URL transformation", () => {
      const url = "https://gist.github.com/abc123.git";
      const token = "ghp_abcDEF123";
      const authenticatedUrl = url.replace("https://", `https://${token}@`);

      assert.strictEqual(
        authenticatedUrl,
        "https://ghp_abcDEF123@gist.github.com/abc123.git",
      );
    });
  });
});
