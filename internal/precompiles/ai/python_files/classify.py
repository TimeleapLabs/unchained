import sys
from transformers import AutoImageProcessor, AutoModelForImageClassification
from PIL import Image
import torch


def classify_image(image_path):
    # Load the image processor and model
    processor = AutoImageProcessor.from_pretrained(
        "google/vit-base-patch16-224")
    model = AutoModelForImageClassification.from_pretrained(
        "google/vit-base-patch16-224")

    # Open the image file
    image = Image.open(image_path).convert("RGB")

    # Preprocess the image
    inputs = processor(images=image, return_tensors="pt")

    # Forward pass through the model
    with torch.no_grad():
        outputs = model(**inputs)

    # Get the predicted label
    logits = outputs.logits
    predicted_class_idx = logits.argmax(-1).item()
    return model.config.id2label[predicted_class_idx]


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python image_classifier.py <path_to_image>")
        sys.exit(1)

    image_path = sys.argv[1]
    predicted_label = classify_image(image_path)
    print(f"Predicted class: {predicted_label}")
