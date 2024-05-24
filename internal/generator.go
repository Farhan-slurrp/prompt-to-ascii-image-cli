package internal

import (
	"image"
	"os/exec"
)

func GenerateImage(prompt string) (image.Image, error) {
	_, err := exec.Command("python3", "generate_image.py", "--prompt", prompt).Output()
	if err != nil {
		return nil, err
	}

	image, err := LoadImage("output.png")
	if err != nil {
		return nil, err
	}

	return image, nil
}
