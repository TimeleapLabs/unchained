import enquirer from "enquirer";

const prompt = new enquirer.Password({
  name: "key",
  message: "What is your private key?",
});

export const ask = () => prompt.run();
