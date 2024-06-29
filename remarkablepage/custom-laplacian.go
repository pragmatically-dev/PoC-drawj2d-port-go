package remarkablepage

import (
	"errors"

	"image"
)

var kernel4 = Kernel{Content: [][]float64{
	{0, 1, 0},
	{1, -4, 1},
	{0, 1, 0},
}, Width: 3, Height: 3}

var kernel8 = Kernel{Content: [][]float64{
	{1, 1, 1},
	{1, -8, 1},
	{1, 1, 1},
}, Width: 3, Height: 3}

var sharpen = Kernel{Content: [][]float64{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}, Width: 3, Height: 3}

var kernel9 = Kernel{Content: [][]float64{
	{1, 1, 1},
	{1, -9, 1},
	{1, 1, 1},
}, Width: 3, Height: 3}

var kernel13 = Kernel{Content: [][]float64{
	{2, 2, 2},
	{2, -13, 2},
	{2, 2, 2},
}, Width: 3, Height: 3}

var sobelY = Kernel{Content: [][]float64{
	{-1, -2, -1},
	{0, 0, 0},
	{1, 2, 1},
}, Width: 3, Height: 3}

var sobelX = Kernel{Content: [][]float64{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}, Width: 3, Height: 3}

var gaussianBlur = Kernel{Content: [][]float64{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}, Width: 3, Height: 3}

// LaplacianKernel - constant type for differentiating Laplacian kernels
type LaplacianKernel int

const (
	// K4 Laplacian kernel:
	//	{0, 1, 0},
	//	{1, -4, 1},
	//	{0, 1, 0},
	K4 LaplacianKernel = iota
	// K8 Laplacian kernel:
	//	{0, 1, 0},
	//	{1, -8, 1},
	//	{0, 1, 0},
	K8

	Sharpen

	K9

	K13
	SobelY
	SobelX
	Gaussian
)

// LaplacianGray applies Laplacian filter to a grayscale image. The kernel types are: K4 and K8 (see LaplacianKernel)
// Example of usage:
//
//	res, err := edgedetection.LaplacianGray(img, paddding.BorderReflect, edgedetection.K8)
func LaplacianGray(gray *image.Gray, border CBorder, kernel LaplacianKernel) (*image.Gray, error) {
	var laplacianKernel Kernel
	switch kernel {
	case K4:
		laplacianKernel = kernel4
	case K8:
		laplacianKernel = kernel8
	case Sharpen:
		laplacianKernel = sharpen
	case K9:
		laplacianKernel = kernel9
	case K13:
		laplacianKernel = kernel13
	case SobelY:
		laplacianKernel = sobelY
	case SobelX:
		laplacianKernel = sobelX
	case Gaussian:
		laplacianKernel = gaussianBlur
	default:
		return nil, errors.New("invalid kernel")
	}
	return ConvolveGray(gray, &laplacianKernel, image.Point{X: 1, Y: 1}, border)
}