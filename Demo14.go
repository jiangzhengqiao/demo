package main

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

func main() {

	// // input files
	// files := []string{"01.jpg", "02.jpg", "03.jpg"}

	// // load images and make 100x100 thumbnails of them
	// var thumbnails []image.Image
	// for _, file := range files {
	// 	img, err := imaging.Open(file)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	thumb := imaging.Thumbnail(img, 100, 100, imaging.CatmullRom)
	// 	thumbnails = append(thumbnails, thumb)
	// }

	// // create a new blank image
	// dst := imaging.New(100*len(thumbnails), 100, color.NRGBA{0, 0, 0, 0})

	// // paste thumbnails into the new image side by side
	// for i, thumb := range thumbnails {
	// 	dst = imaging.Paste(dst, thumb, image.Pt(i*100, 0))
	// }

	// // save the combined image to file
	// err := imaging.Save(dst, "dst.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	img, err := imaging.Open("/Users/jiangzhengqiao/Desktop/1.jpg")
	if err != nil {
		panic(err)
	}

	dstImage := imaging.Blur(img, 0.5)
	err = imaging.Save(dstImage, "/Users/jiangzhengqiao/Desktop/5.jpg")
	if err != nil {
		panic(err)
	}
}
