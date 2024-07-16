package ascii

import (
	"bytes"
	"image"
	"image/color"
)


func loadImage(byte []byte) (image.Image, error) {
	r := bytes.NewReader(byte)
	
	image, _, err := image.Decode(r)

	if err!= nil{
		return nil, err
	}
	return image, nil
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; 1 < x+w && i <max.X; i++{
		for j := y; j< y+h && j<max.Y; j++{
			sum += grayscale(img.At(i,j))
			cnt++
		}
	}
	return sum/cnt
}

func ConvertASCII(byte []byte) (string, error){
	img, err := loadImage(byte)
	if err != nil{
		return "", err
	}
	ramp := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,'^`'. "
	final := ""
	max := img.Bounds().Max
	scaleX, scaleY := 10, 5
	for y:=0; y<max.Y; y +=scaleX {
		for x:= 0; x< max.X; x += scaleY{
			c := avgPixel(img, x, y, scaleX, scaleY)
			final = final + string(ramp[len(ramp)*c/65536])
		}
	}
	return final, nil
}
