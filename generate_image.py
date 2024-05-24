import argparse
import os
from diffusers import StableDiffusionPipeline


if __name__ == "__main__":
    try:
        os.remove("output.jpeg")
    except FileNotFoundError:
        pass
    parser = argparse.ArgumentParser(description='Diffusers CLI')
    parser.add_argument('--prompt', metavar='prompt', required=True,
                        help='Prompt for the image generation')
    args = parser.parse_args()
    prompt = args.prompt
    
    pipe = StableDiffusionPipeline.from_pretrained("CompVis/stable-diffusion-v1-4")
    pipe = pipe.to("mps")

    _ = pipe(prompt, num_inference_steps=1)
    image = pipe(prompt).images[0]
    image.save("output.jpeg")