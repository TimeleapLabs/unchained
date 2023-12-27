import { spawn } from "child_process";
import { fileURLToPath } from "url";
import { dirname, join } from "path";
import { config } from "../../constants.js";
const __filename = fileURLToPath(import.meta.url);
const __root = new Array(5)
    .fill(null)
    .reduce((path) => dirname(path), __filename);
export const runPrismaCommand = async (args) => {
    const prisma = `prisma${process.platform === "win32" ? ".cmd" : ""}`;
    const prismaCliPath = join(__root, "node_modules", ".bin", prisma);
    const prismaProcess = spawn(prismaCliPath, args, {
        stdio: "inherit",
        env: { ...process.env, UNCHAINED_DATABASE_URL: config.database?.url },
        shell: true,
        cwd: __root,
    });
    return new Promise((resolve) => prismaProcess.on("close", resolve));
};
