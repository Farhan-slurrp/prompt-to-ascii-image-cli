package internal

import (
	"image"
	"os/exec"
)

func GenerateImage(prompt string) (image.Image, error) {
	out, err := exec.Command("python3", "generate_image.py", "--prompt", prompt).Output()
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, nil
	}

	image, err := LoadImage("output.jpeg")
	if err != nil {
		return nil, err
	}

	return image, nil
}
