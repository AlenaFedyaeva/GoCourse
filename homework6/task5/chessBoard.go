package task5

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)
//Rectangle - структура квадрас
type Rectangle struct {
	p image.Point
	a int
}
//ColorModel from Image interface
func (c *Rectangle) ColorModel() color.Model {
	return color.AlphaModel
}
//Bounds from Image interface
func (c *Rectangle) Bounds() image.Rectangle {
	return image.Rect(c.p.X, c.p.Y, c.p.X+c.a, c.p.Y+c.a)
}
//At from Image interface
func (c *Rectangle) At(x, y int) color.Color {
	return color.Alpha{255}
}
//ChessBoardImg - create png file with chess board
func ChessBoardImg() {
	fname:="task5/chessBoard.png"
	file, err := os.Create(fname)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer file.Close()

	cellNum := 8
	step := 100
	//Белая доска
	img := image.NewRGBA(image.Rect(0, 0, step*cellNum, step*cellNum))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	mask := image.NewRGBA(image.Rect(0, 0, step*cellNum, step*cellNum))
	draw.Draw(mask, mask.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	//На белой доске рисуем только черные квардраты
	for i := 0; i < cellNum; i++ {
		// нечетный ряд x
		if i%2 == 0 {
			for j := 0; j < cellNum; j++ {
				if j%2 == 0 { //нечетный ряд у
					fmt.Print("1| ", i, j, " |")
					draw.DrawMask(img, img.Bounds(), mask, image.ZP,
					&Rectangle{image.Point{i*step, j*step}, step}, image.ZP, draw.Over)
				} else { //четный ряд у
					fmt.Print("0")
				}
			}
		} else { //четный ряд х
			for j := 0; j < cellNum; j++ {
				if j%2 == 0 { //нечетный ряд у
					fmt.Print("0")
				} else {
					fmt.Print("1| ", i, j, " |")
					draw.DrawMask(img, img.Bounds(), mask, image.ZP,
					&Rectangle{image.Point{i*step, j*step}, step}, image.ZP, draw.Over)
				}
			}
		}
		fmt.Println("")
	}
	//Сохраняем в файл
	png.Encode(file, img)
	//Оставила fmt.printf() и кейсы с ним, чтобы было легче понять как вычислены координаты
}