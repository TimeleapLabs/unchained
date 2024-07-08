import sys
import torch
from diffusers import DiffusionPipeline
import io
import os
from sia import Sia
from torch_device import get_device

pipelines = {}

OPEN_SOURCE_MODELS = [
    "segmind/SSD-1B",
    "Corcelio/mobius",
    "segmind/Segmind-Vega",
    "Corcelio/openvision",
    "SimianLuo/LCM_Dreamshaper_v7",
    "OEvortex/PixelGen"
]

NON_FREE_MODELS = [
    "fluently/Fluently-XL-Final",
    "alvdansen/littletinies",
    "cagliostrolab/animagine-xl-3.1",
    "SG161222/Realistic_Vision_V6.0_B1_noVAE",
    "Lykon/dreamshaper-xl-v2-turbo",
    "UnfilteredAI/NSFW-gen-v2.1",
]


def get_pipeline(model_name, lora_weights=None):
    key = (model_name + ":::" + lora_weights) if lora_weights else model_name

    if key not in pipelines:
        pipelines[key] = DiffusionPipeline.from_pretrained(
            model_name, torch_dtype=torch.float16)

        if lora_weights:
            pipelines[key].load_lora_weights(lora_weights)

        pipelines[key].to(get_device())

        pipelines[key].safety_checker = lambda images, **kwargs: (
            images, [False] * len(images))

    return pipelines[key]


def image_to_bytes(image):
    byte_io = io.BytesIO()
    image.save(byte_io, format="PNG")
    return byte_io.getvalue()


def parse_packet(packet: Sia):
    # get uuid v7 from packet
    uuid = packet.read_byte_array_n(16)
    model = packet.read_string8()
    lora_weights = packet.read_string8()
    steps = packet.read_uint8()
    # get prompt length from packet (little endian uint16 at offset 17)
    prompt = packet.read_string16()
    negative_prompt = packet.read_string16()
    return {
        "uuid": uuid,
        "prompt": prompt,
        "model": model,
        "lora_weights": lora_weights,
        "negative_prompt": negative_prompt,
        "steps": steps
    }


def pack_response_packet(uuid, response: bytes):
    return Sia().add_byte_array_n(uuid).add_byte_array32(response).content


def request_handler(packet):
    parsed = parse_packet(packet)
    pipe = get_pipeline(parsed["model"], parsed["lora_weights"])

    images = pipe(
        prompt=parsed["prompt"],
        num_inference_steps=parsed["steps"],
        guidance_scale=5.0,
        lcm_origin_steps=50,
        height=1024,
        width=1024,
        negative_prompt=parsed["negative_prompt"],
        num_images_per_prompt=1,
        output_type="pil").images

    response = image_to_bytes(images[0])

    return pack_response_packet(parsed["uuid"], response)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py '<caption>'")
        sys.exit(1)

    caption = sys.argv[1]

    for model_name in [*OPEN_SOURCE_MODELS, *NON_FREE_MODELS]:
        pipe = DiffusionPipeline.from_pretrained(
            model_name, torch_dtype=torch.float32)

        pipe.to(get_device())

        lora_weights = os.getenv("IMAGE_TO_TEXT_LORA_WEIGHTS")

        if lora_weights:
            pipe.load_lora_weights(lora_weights)

        pipe.safety_checker = lambda images, **kwargs: (
            images, [False] * len(images))

        num_inference_steps_str = os.getenv(
            "IMAGE_TO_TEXT_STEPS") or "32"
        num_inference_steps = int(num_inference_steps_str)

        images = pipe(
            prompt=caption,
            num_inference_steps=num_inference_steps,
            guidance_scale=4.0,
            lcm_origin_steps=50,
            output_type="pil").images

        # Save the generated images
        for idx, img in enumerate(images):
            model_name_stripped = model_name.replace("/", "_")
            img.save(f"generated_image_{model_name_stripped}_{idx}.png")

    print("Images saved successfully.")
