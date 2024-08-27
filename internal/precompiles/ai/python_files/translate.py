# Use a pipeline as a high-level helper
from transformers import pipeline
from sia import Sia
import sys

loaded_models = {}

def parse_packet(packet: Sia):
    uuid = packet.read_byte_array_n(16)
    fromLang = packet.read_string_n(2)
    toLang = packet.read_string_n(2)
    prompt = packet.read_string16()
    return {"uuid": uuid, "from": fromLang, "to": toLang, "prompt": prompt}


def pack_response_packet(uuid, response: str):
    return Sia().add_byte_array_n(uuid).add_string16(response).content


def request_handler(packet: Sia):
    data = parse_packet(packet)
    fromLang = data["from"]
    toLang = data["to"]
    prompt = data["prompt"]

    key = f"{fromLang}-{toLang}"
    if key not in loaded_models:
        loaded_models[key] = pipeline(
            "translation", model=f"Helsinki-NLP/opus-mt-{fromLang}-{toLang}")

    output = loaded_models[key](prompt)
    response = output[0]["translation_text"]

    return pack_response_packet(data["uuid"], response)

if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Usage: python translate.py <from> <to> <text>")
        sys.exit(1)

    fromLang = sys.argv[1]
    toLang = sys.argv[2]
    text = sys.argv[3]

    pipe = pipeline(
        "translation", model=f"Helsinki-NLP/opus-mt-{fromLang}-{toLang}")

    output = pipe(text)
    print(output[0]["translation_text"])
