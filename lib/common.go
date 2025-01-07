package lib

import "gopkg.in/gographics/imagick.v2/imagick"

func ConvertPDF2PNG(pdfFile, pngFile string, dpi int, rot int) (err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err = mw.SetResolution(float64(dpi), float64(dpi))
	if err != nil {
		return err
	}

	err = mw.ReadImage(pdfFile)
	if err != nil {
		return err
	}

	if mw.GetNumberImages() > 1 {
		mw = mw.MergeImageLayers(imagick.IMAGE_LAYER_FLATTEN)
	}

	drawingWand := imagick.NewDrawingWand()
	defer drawingWand.Destroy()
	pixelWand := imagick.NewPixelWand()
	defer pixelWand.Destroy()
	pixelWand.SetColor("white")
	drawingWand.SetFillColor(pixelWand)

	err = mw.RotateImage(pixelWand, float64(rot))
	if err != nil {
		return err
	}

	//mw.SetIteratorIndex(0)
	err = mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	if err != nil {
		return err
	}

	err = mw.SetImageFormat("png")
	if err != nil {
		return err
	}

	err = mw.WriteImage(pngFile)
	if err != nil {
		return err
	}
	//
	return nil
}
