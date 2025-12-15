import { spawn } from "child_process";
import { GistDataEntry } from "../fsdata/types";
import { createTempFolder } from "../utils";

// TODO: Optimize to clone all gists in parallel with a limit
export async function cloneGistEntries(
  entries: GistDataEntry[],
  temporalFolder: string,
  token: string,
): Promise<void> {
  let index = 1;
  for (const entry of entries) {
    const paddedIndex = index.toString().padStart(3, "0");
    const subTemporalFolder = await createTempFolder(
      temporalFolder,
      `${paddedIndex}_gist_${entry.id}`,
    );
    await gitClone(entry.gistPullUrl, subTemporalFolder, token);
    index++;
  }
}

async function gitClone(
  url: string,
  destination: string,
  token: string,
): Promise<void> {
  return new Promise((resolve, reject) => {
    // Inject token into the URL for authentication
    const authenticatedUrl = url.replace("https://", `https://${token}@`);

    const gitProcess = spawn("git", ["clone", authenticatedUrl, destination]);

    let stderr = "";

    gitProcess.stderr.on("data", (data) => {
      stderr += data.toString();
    });

    gitProcess.on("close", (code) => {
      if (code !== 0) {
        reject(new Error(`Git clone failed: ${stderr}`));
      } else {
        resolve();
      }
    });

    gitProcess.on("error", (error) => {
      reject(new Error(`Failed to spawn git process: ${error.message}`));
    });
  });
}
