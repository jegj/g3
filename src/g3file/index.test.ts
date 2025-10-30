import { describe, it } from "node:test";
import assert from "node:assert";
import { G3File } from "./index";

describe("G3File.hasMultipleFiles()", () => {
  it("returns true with multiple entries", () => {
    const g3file = new G3File({
      entries: [
        { path: "/path1", files: { "file1.txt": {} } },
        { path: "/path2", files: { "file2.txt": {} } },
      ],
    } as any);

    assert.strictEqual(g3file.hasMultipleFiles(), true);
  });

  it("returns true with three entries", () => {
    const g3file = new G3File({
      entries: [
        { path: "/path1", files: { "file1.txt": {} } },
        { path: "/path2", files: { "file2.txt": {} } },
        { path: "/path3", files: { "file3.txt": {} } },
      ],
    } as any);

    assert.strictEqual(g3file.hasMultipleFiles(), true);
  });

  it("returns true with single entry having multiple files", () => {
    const g3file = new G3File({
      entries: [
        {
          path: "/path1",
          files: {
            "file1.txt": {},
            "file2.txt": {},
          },
        },
      ],
    } as any);

    assert.strictEqual(g3file.hasMultipleFiles(), true);
  });

  it("returns true with single entry having many files", () => {
    const g3file = new G3File({
      entries: [
        {
          path: "/path1",
          files: {
            "file1.txt": {},
            "file2.txt": {},
            "file3.txt": {},
            "file4.txt": {},
          },
        },
      ],
    } as any);

    assert.strictEqual(g3file.hasMultipleFiles(), true);
  });

  it("returns false with single entry having one file", () => {
    const g3file = new G3File({
      entries: [
        {
          path: "/path1",
          files: { "file1.txt": {} },
        },
      ],
    } as any);

    assert.strictEqual(g3file.hasMultipleFiles(), false);
  });

  it("returns false with empty entries", () => {
    const g3file = new G3File({
      entries: [],
    } as any);

    assert.strictEqual(g3file.hasMultipleFiles(), false);
  });
});
