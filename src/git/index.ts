import { spawn } from "child_process";

export async function gitClone(url: string, folder: string): Promise<void> {
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
