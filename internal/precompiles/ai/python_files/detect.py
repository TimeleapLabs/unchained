import sys
from transformers import AutoImageProcessor, AutoModelForObjectDetection
from PIL import Image, ImageDraw, ImageFont
import torch


def detect_objects(image_path, font_path='noto.ttf', font_size=32):
    # Load the image processor and model
    processor = AutoImageProcessor.from_pretrained("hustvl/yolos-small")
    model = AutoModelForObjectDetection.from_pretrained("hustvl/yolos-small")

    # Open the image file
    image = Image.open(image_path).convert("RGB")

    # Preprocess the image
    inputs = processor(images=image, return_tensors="pt")

    # Forward pass through the model
    with torch.no_grad():
        outputs = model(**inputs)

    # Process outputs
    target_sizes = torch.tensor([image.size[::-1]])
    results = processor.post_process_object_detection(
        outputs, target_sizes=target_sizes, threshold=0.9)[0]

    # Load the font
    font = ImageFont.truetype(font_path, font_size)

    # Draw bounding boxes and labels on the image
    draw = ImageDraw.Draw(image)
    for score, label, box in zip(results["scores"], results["labels"], results["boxes"]):
        box = [round(i, 2) for i in box.tolist()]
        draw.rectangle(box, outline="red", width=3)
        text = f"{model.config.id2label[label.item()]}: {
            round(score.item(), 3)}"
        text_bbox = draw.textbbox((box[0], box[1]), text, font=font)
        text_location = (box[0], box[1] - (text_bbox[3] - text_bbox[1]))
        draw.rectangle(
            [text_location, (text_bbox[2], text_bbox[3])], fill="red")
        draw.text((box[0], box[1] - (text_bbox[3] - text_bbox[1])),
                  text, fill="white", font=font)

    return image


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python object_detector.py <path_to_image>")
        sys.exit(1)

    image_path = sys.argv[1]
    detected_image = detect_objects(image_path)
    detected_image.show()  # Display the image with detections
