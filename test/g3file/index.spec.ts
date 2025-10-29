import assert from "node:assert";
import { describe, it } from "node:test";
import { G3Config } from "../../src/config";
import { createG3FileFactory } from "../../src/g3file";
import { G3Dependecies } from "../../src/types";

describe("G3File", () => {
  const g3ConfigMock: G3Config = {
    GITHUB_TOKEN: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
    AES_KEY: "aes_22faketoken_asieowiafhjls2basdjnnmazxwa",
    DATA_FOLDER: "/home/fakeuser/.local/share/g3/files/",
  };
  const dependencies: G3Dependecies = { config: g3ConfigMock };
  const g3FileFactory = createG3FileFactory(dependencies);

  describe("createG3FileFactory", () => {
    it("should create a G3File with absolute path", async () => {
      const absolutePath = "/home/user/documents/report.pdf";
      const result = await g3FileFactory(absolutePath);

      assert.strictEqual(result.filename, "report.pdf");
      assert.strictEqual(result.filepath, absolutePath);
      assert.strictEqual(result.g3Filename, "report.pdf.g3.json");
      assert.strictEqual(
        result.g3Filepath,
        "/home/fakeuser/.local/share/g3/files/report.pdf.g3.json",
      );
    });
  });
});
