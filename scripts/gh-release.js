import { writeFileSync, readFileSync, mkdirSync, copyFileSync } from "fs";
import { dirname } from "path";
import { execSync } from "child_process";

const packageJson = JSON.parse(readFileSync("./package.json"));

const makeReleaseNotes = async () => {
  const url =
    "https://app.whatthediff.ai/changelog/github/KenshiTech/unchained.json";
  const req = await fetch(url);
  const notes = await req.json();
  const wtd = notes.changelog.data[0];

  const changes = wtd.comment.replace(/^## PR Summary/, "## Changes");
  const release = `\
# Unchained ${packageJson.version}

${changes}
`;

  writeFileSync("./release-notes.md", release);
};

const makeReleaseFile = (name, files) => {
  const dirName = `unchained-v${packageJson.version}-${name}`;
  mkdirSync(`release/${dirName}`, { recursive: true });
  for (const { source, target } of files) {
    const dirToMake = dirname(target);
    if (dirToMake !== ".") {
      mkdirSync(`release/${dirName}/${dirToMake}`, { recursive: true });
    }
    copyFileSync(source, `release/${dirName}/${target}`);
  }
  execSync(`cd release && zip -r ${dirName}.zip ${dirName}`);
};

await makeReleaseNotes();
mkdirSync("./release", { recursive: true });
makeReleaseFile("docker", [
  { source: "docker/compose.yml", target: "compose.yml" },
  { source: "docker/Dockerfile", target: "Dockerfile" },
  { source: "docker/start_node.sh", target: "scripts/start_node.sh" },
  { source: "docker/stop_node.sh", target: "scripts/stop_node.sh" },
  { source: "docker/view_node.sh", target: "scripts/view_node.sh" },
  { source: ".env.template", target: ".env.template" },
  { source: "conf.yaml.template", target: "conf.yaml.template" },
]);
