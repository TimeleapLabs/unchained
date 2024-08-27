import torch


def get_device():
    if torch.cuda.is_available():
        return "cuda"
    # detect if m1/m2/m3
    if torch.backends.mps.is_available():
        return "mps"
    # detect if vulkan (for android, raspberry pi, etc.)
    if torch.is_vulkan_available():
        return "vulkan"
    return "cpu"
