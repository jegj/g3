import { spawn } from "child_process";

export async function gitClone(
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
