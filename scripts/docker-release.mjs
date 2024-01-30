import { writeFileSync, readFileSync, mkdirSync, copyFileSync } from "fs";
import { dirname } from "path";
import { execSync } from "child_process";

const about = [
  "Unchained is a decentralized, peer-to-peer network for data validation.",
  "Unchained nodes work to validate data together and are rewarded in KNS tokens.",
  "The validated data can then be queried by consumer in exchange for KNS tokens.",
  "Learn more about Unchained [here](https://kenshi.io/docs/unchained).",
].join(" ");

const questions = [
  "Have any questions? Ask in the [forum](https://forum.kenshi.io/c/unchained),",
  "in our [chat](https://t.me/KenshiTech/85602), or send us an",
  "[email](mailto:hi@kenshi.io).",
].join(" ");

const releaseTemplate = (changes) => `\
${about}

${questions}

${changes}
  `;

const makeReleaseNotes = async () => {
  const url =
    "https://app.whatthediff.ai/changelog/github/KenshiTech/unchained.json";
  const req = await fetch(url);
  const notes = await req.json();
  const wtd = notes.changelog.data[0];

  const changes = wtd.comment.replace(/^## PR Summary/, "## Changes");
  const release = releaseTemplate(changes);

  writeFileSync("./release-notes.md", release);
};

const makeReleaseFile = (name, files) => {
  const lastTag = execSync(`git tag | head -1`).toString().trim();
  const dirName = `unchained-${lastTag}-${name}`;
  mkdirSync(`release/${dirName}`, { recursive: true });
  for (const { source, target } of files) {
    const dirToMake = dirname(target);
    if (dirToMake !== ".") {
      mkdirSync(`release/${dirName}/${dirToMake}`, { recursive: true });
    }
    copyFileSync(source, `release/${dirName}/${target}`);
  }
  execSync(`cd release && zip -r ${dirName}.zip ${dirName}`);
  execSync(`cp release/${dirName}.zip src/bin/`);
};

await makeReleaseNotes();
mkdirSync("./release", { recursive: true });

makeReleaseFile("docker", [
  { source: "docker/compose.yaml", target: "compose.yaml" },
  { source: "docker/Dockerfile", target: "Dockerfile" },
  { source: "docker/unchained.sh", target: "unchained.sh" },
  // { source: "docker/grafana/dashboard.yaml", target: "grafana/dashboard.yaml" },
  // {
  //   source: "docker/grafana/datasource.yaml",
  //   target: "grafana/datasource.yaml",
  // },
  // {
  //   source: "docker/grafana/dashboards/unchained.json",
  //   target: "grafana/dashboards/unchained.json",
  // },
  // { source: ".env.template", target: ".env.template" },
  // { source: "conf.broker.yaml.template", target: "conf.broker.yaml.template" },
  // { source: "conf.remote.yaml.template", target: "conf.remote.yaml.template" },
  { source: "conf.worker.yaml.template", target: "conf.worker.yaml.template" },
  { source: "LICENSE", target: "LICENSE" },
]);
