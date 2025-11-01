import assert from "node:assert";
import { describe, it } from "node:test";
import { G3Config } from "../../src/config";
import { createG3FileFactory } from "../../src/g3file";
import { G3Dependecies } from "../../src/types";
import { FilesystemDataEntry } from "../../src/fsdata/types";
import { generateMockGistDataEntry } from "../factories/gistDataEntry.mock";

const g3ConfigMock: G3Config = {
  GITHUB_TOKEN: "github_22faketoken_asieowiafhjls2basdjnnmazxwa",
  AES_KEY: "aes_22faketoken_asieowiafhjls2basdjnnmazxwa",
  DATA_FOLDER: "/home/fakeuser/.local/share/g3/files/",
};
const dependencies: G3Dependecies = { config: g3ConfigMock };

describe("G3File", () => {
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

  describe("hasMultipleFiles", () => {
    it("returns true with multiple entries", async () => {
      const getG3FSEntryMock = async () => {
        return {
          entries: [generateMockGistDataEntry(1), generateMockGistDataEntry(1)],
        } as FilesystemDataEntry;
      };
      const g3FileFactory = createG3FileFactory(dependencies, getG3FSEntryMock);
      const g3file = await g3FileFactory("/path1/file1.txt");

      assert.strictEqual(g3file.hasMultipleFiles(), true);
    });

    it("returns false with single entry having multiple files", async () => {
      const getG3FSEntryMock = async () => {
        return {
          entries: [generateMockGistDataEntry(3)],
        } as FilesystemDataEntry;
      };
      const g3FileFactory = createG3FileFactory(dependencies, getG3FSEntryMock);
      const g3file = await g3FileFactory("/path1/file1.txt");

      assert.strictEqual(g3file.hasMultipleFiles(), false);
    });

    it("returns false with single entry having one file", async () => {
      const getG3FSEntryMock = async () => {
        return {
          entries: [generateMockGistDataEntry(1)],
        } as FilesystemDataEntry;
      };
      const g3FileFactory = createG3FileFactory(dependencies, getG3FSEntryMock);

      const g3file = await g3FileFactory("/path1/file1.txt");

      assert.strictEqual(g3file.hasMultipleFiles(), false);
    });

    it("returns false with empty entries", async () => {
      const g3file = await g3FileFactory("/path1/file1.txt");

      assert.strictEqual(g3file.hasMultipleFiles(), false);
    });
  });
});
