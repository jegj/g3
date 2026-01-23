import assert from "node:assert";
import { describe, it } from "node:test";
import { g3Error } from "../../src/error";

describe("Error", () => {
  describe("g3Error", () => {
    it("should create an error with the provided message", () => {
      const error = g3Error("Something went wrong");

      assert.strictEqual(error.message, "Something went wrong");
    });

    it("should set isOperational to true", () => {
      const error = g3Error("Test error");

      assert.strictEqual(error.isOperational, true);
    });

    it("should set name to G3Error", () => {
      const error = g3Error("Test error");

      assert.strictEqual(error.name, "G3Error");
    });

    it("should include details when provided", () => {
      const error = g3Error("Main error", "Additional details here");

      assert.strictEqual(error.details, "Additional details here");
    });

    it("should set details to empty string when not provided", () => {
      const error = g3Error("Main error");

      assert.strictEqual(error.details, "");
    });

    it("should include custom options in the error object", () => {
      const error = g3Error("Test error", "details", {
        code: "ERR_NETWORK",
        statusCode: 500,
      });

      assert.strictEqual(error.code, "ERR_NETWORK");
      assert.strictEqual(error.statusCode, 500);
    });

    it("should have a proper stack trace", () => {
      const error = g3Error("Test error");

      assert.ok(error.stack);
      assert.ok(error.stack.includes("G3Error"));
    });

    it("should be an instance of Error", () => {
      const error = g3Error("Test error");

      assert.ok(error instanceof Error);
    });

    it("should handle empty message", () => {
      const error = g3Error("");

      assert.strictEqual(error.message, "");
      assert.strictEqual(error.name, "G3Error");
      assert.strictEqual(error.isOperational, true);
    });

    it("should handle multiple custom options", () => {
      const error = g3Error("Error", "details", {
        code: "ERR_CUSTOM",
        retryable: true,
        timestamp: 1234567890,
      });

      assert.strictEqual(error.code, "ERR_CUSTOM");
      assert.strictEqual(error.retryable, true);
      assert.strictEqual(error.timestamp, 1234567890);
    });
  });
});
