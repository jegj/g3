import { spawn } from "child_process";
import { GistDataEntry } from "../fsdata/types";
import { createTempFolder } from "../utils";

// TODO: Optimize to clone all gists in parallel with a limit
export async function cloneGistEntries(
  entries: GistDataEntry[],
  temporalFolder: string,
): Promise<void> {
  let index = 1;
  for (const entry of entries) {
    console.log(`Cloning gist ${entry.gistPullUrl} ...`);
    const paddedIndex = index.toString().padStart(3, "0");
    const subTemporalFolder = await createTempFolder(
      temporalFolder,
      `${paddedIndex}_gist_${entry.id}`,
    );
    await gitClone(entry.gistPullUrl, subTemporalFolder);
    index++;
  }
}

//TODO: This use the git local installation, try to use GIT token instead
async function gitClone(url: string, folder: string): Promise<void> {
  return new Promise((resolve, reject) => {
    const gitProcess = spawn("git", ["clone", url, folder]);

    let stderr = "";

    gitProcess.stdout.on("data", (data) => {
      console.log(data.toString());
    });

    gitProcess.stderr.on("data", (data) => {
      stderr += data.toString();
      console.error(data.toString());
    });

    gitProcess.on("close", (code) => {
      if (code === 0) {
        resolve();
      } else {
        reject(new Error(`git clone failed with code ${code}: ${stderr}`));
      }
    });

    gitProcess.on("error", (error) => {
      reject(new Error(`Failed to start git process: ${error.message}`));
    });
  });
}
