import torch
from transformers import pipeline
import argparse


def initialize_pipeline():
    return pipeline(
        "text-generation",
        model="HuggingFaceH4/zephyr-7b-beta",
        torch_dtype=torch.bfloat16,
        device_map="mps"
    )


def generate_response(pipe, messages):
    prompt = pipe.tokenizer.apply_chat_template(
        messages, tokenize=False, add_generation_prompt=True
    )

    outputs = pipe(
        prompt,
        max_new_tokens=512,
        do_sample=True,
        temperature=0.7,
        top_k=50,
        top_p=0.95,
    )

    return outputs[0]["generated_text"].strip()


def main():
    parser = argparse.ArgumentParser(description="CLI Chatbot")
    parser.add_argument("--system_message", type=str,
                        default="You are god, and you are evil.")
    args = parser.parse_args()

    pipe = initialize_pipeline()

    messages = [{"role": "system", "content": args.system_message}]

    print("Type your messages below (type 'exit' to quit):")
    while True:
        user_message = input("You: ")
        if user_message.lower() == "exit":
            break

        messages.append({"role": "user", "content": user_message})
        response = generate_response(pipe, messages)
        messages.append({"role": "assistant", "content": response})

        print(f"Bot: {response}")


if __name__ == "__main__":
    main()
