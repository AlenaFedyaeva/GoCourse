package task2

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var red  color.Color = color.RGBA{200, 30, 30, 255}
//DrawLines - draw horizontal & vertical lines
func DrawLines()  {
	file, err := os.Create("task2/task2.png")
    if err != nil {
        fmt.Errorf("%s", err)
    }
	defer file.Close()

	size:=800
	step:=100
    img := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for x := 0; x < size; x++ {
		for y:=step; y<size; y+=step {
			img.Set(x, y, red)
			img.Set(y, x, color.Black)
		}
	}
	//Сохраняем в файл
    png.Encode(file, img)
}
