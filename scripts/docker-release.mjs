import { writeFileSync, mkdirSync, copyFileSync } from "fs";
import { dirname } from "path";
import { execSync } from "child_process";

const about = [
  "Unchained is a decentralized, peer-to-peer network for data validation.",
  "Unchained nodes work to validate data together and are rewarded in KNS tokens.",
  "The validated data can then be queried by consumer in exchange for KNS tokens.",
  "Learn more about Unchained [here](https://timeleap.swiss/docs/unchained).",
].join(" ");

const questions = [
  "Have any questions? Ask in the [forum](https://forum.timeleap.swiss/c/unchained),",
  "in our [chat](https://t.me/TimeleapTech/85602), or send us an",
  "[email](mailto:hi@timeleap.swiss).",
].join(" ");

const releaseTemplate = () => `\
${about}

${questions}
  `;

const makeReleaseNotes = async () => {
  const release = releaseTemplate();

  writeFileSync("./release-notes.md", release);
};

const makeReleaseFile = (name, files) => {
  const lastTag = process.argv[2];
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
  execSync(`cp release/${dirName}.zip bin/`);
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
