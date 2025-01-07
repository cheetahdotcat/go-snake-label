package lib

import (
	"fmt"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func DHLPrivat(inputFile, outputFile string) (err error) {
	canvasWidth := uint(1642)
	canvasHeight := uint(696)
	//
	srcWand := imagick.NewMagickWand()
	defer srcWand.Destroy()

	if err = srcWand.ReadImage(inputFile); err != nil {
		return err
	}

	var (
		barcodeSizeX int = 710
		barcodeSizeY int = 320
		scSize       int = 296
	)
	sections := []struct {
		srcX, srcY, srcW, srcH,
		tgtX, tgtY, tgtW, tgtH,
		rot int
	}{

		{ // Kopf
			1964, 106, 1124, 94,
			0, 0, 890, 74,
			0,
		},
		{ // Adresse
			1964, 210, 785, 625,
			0, 95, 580, 465,
			0,
		},
		{ // Sicherheitscode
			2763, 215, scSize, scSize,
			594, 90, scSize, scSize,
			0,
		},
		{ // Sicherheitscode Text
			3075, 244, 20, 194,
			645, 395, 200, 20,
			270,
		},
		{ // Bahntransport
			2802, 679, 234, 154,
			666, 465, 152, 100,
			0,
		},
		{ //Sendungsdaten
			1964, 933, 1124, 152,
			0, 576, 890, 120,
			0,
		},
		{ // Leitcode/Routingcode
			2181, 1526, barcodeSizeX, barcodeSizeY,
			930, 20, barcodeSizeX, barcodeSizeY,
			0,
		},
		{ // Identcode/Sendungsnummer
			2181, 1940, barcodeSizeX, barcodeSizeY,
			930, 376, barcodeSizeX, barcodeSizeY,
			0,
		},
	}

	outputWand := imagick.NewMagickWand()
	defer outputWand.Destroy()
	drawingWand := imagick.NewDrawingWand()
	defer drawingWand.Destroy()
	pixelWand := imagick.NewPixelWand()
	defer pixelWand.Destroy()

	// Set canvas size for the output image
	pixelWand.SetColor("white")
	drawingWand.SetFillColor(pixelWand)
	err = outputWand.NewImage(canvasWidth, canvasHeight, pixelWand)
	if err != nil {
		return err
	}

	for i, section := range sections {
		cropped := srcWand.Clone()
		defer cropped.Destroy()

		if err := cropped.CropImage(uint(section.srcW), uint(section.srcH), int(section.srcX), int(section.srcY)); err != nil {
			return fmt.Errorf("failed to crop section %d: %v", i, err)
		}

		if section.rot != 0 {
			err = cropped.RotateImage(pixelWand, float64(section.rot))
			if err != nil {
				return fmt.Errorf("failed to rotate section %d: %v", i, err)
			}
		}
		// Resize cropped section (if needed)
		if err := cropped.ResizeImage(uint(section.tgtW), uint(section.tgtH), imagick.FILTER_LANCZOS, 1); err != nil {
			return fmt.Errorf("failed to resize section %d: %v", i, err)
		}

		//cropped.WriteImage(fmt.Sprintf("section%d.png", i))
		// Composite the cropped section onto the output canvas
		if err := outputWand.CompositeImage(cropped, imagick.COMPOSITE_OP_COPY, int(section.tgtX), int(section.tgtY)); err != nil {
			return fmt.Errorf("failed to composite section %d: %v", i, err)
		}
	}
	{ // Line
		strokeWand := imagick.NewPixelWand()
		defer strokeWand.Destroy()
		strokeWand.SetColor("black")
		lineWand := imagick.NewDrawingWand()
		defer lineWand.Destroy()

		lineWand.SetStrokeColor(strokeWand)
		lineWand.Line(910, 0, 910, float64(canvasHeight))

		err = outputWand.DrawImage(lineWand)
		if err != nil {
			return err
		}
	}

	err = outputWand.SetImageFormat("png")
	if err != nil {
		return err
	}
	err = outputWand.WriteImage(outputFile)
	if err != nil {
		return err
	}
	//
	return nil
}
func DHLPrivatInternational(inputFile, outputFile string) (err error) {
	canvasWidth := uint(2232)
	canvasHeight := uint(696)
	//
	srcWand := imagick.NewMagickWand()
	defer srcWand.Destroy()

	if err = srcWand.ReadImage(inputFile); err != nil {
		return err
	}

	var (
		barcodeSizeX int = 1124
		barcodeSizeY int = 280
		scSize       int = 296
	)
	sections := []struct {
		srcX, srcY, srcW, srcH,
		tgtX, tgtY, tgtW, tgtH,
		rot int
	}{
		{ // Kopf
			1964, 106, 1124, 94,
			0, 0, 890, 74,
			0,
		},
		{ // Adresse
			1964, 210, 785, 625,
			0, 95, 580, 465,
			0,
		},
		{ // Sicherheitscode
			2763, 215, scSize, scSize,
			594, 90, scSize, scSize,
			0,
		},
		{ // Sicherheitscode Text
			3075, 244, 20, 194,
			645, 395, 200, 20,
			270,
		},
		{ // Telefonnummer & E-Mail
			2768, 743, 300, 60,
			590, 495, 300, 60,
			0,
		},
		{ //Sendungsdaten
			1964, 933, 1124, 152,
			0, 576, 890, 120,
			0,
		},
		{ // Icons / Go Green
			1964, 854, 1124, 70,
			920, 696, 696, 43,
			270,
		},
		{ // Tracked
			1964, 1197, 1124, 122,
			986, 696, 696, 76,
			270,
		},
		{ // Tracked Icon
			2846, 530, 130, 202,
			930, 10, 130, 202,
			0,
		},
		{ // Unzustellbarkeit
			1964, 1422, barcodeSizeX, 70,
			1100, 4, barcodeSizeX, 70,
			0,
		},
		{ // Leitcode/Routingcode
			1964, 1634, barcodeSizeX, barcodeSizeY,
			1100, 100, barcodeSizeX, barcodeSizeY,
			0,
		},
		{ // Category Letter
			1964, 2000, 160, barcodeSizeY - 20,
			1100, 416, 160, barcodeSizeY - 20,
			0,
		},
		{ // Identcode/Sendungsnummer
			1964, 2048 + 260, 160, 20,
			1100, 416 + 260, 160, 20,
			0,
		},
		{ // Identcode/Sendungsnummer (2)
			1964 + 160, 2048, barcodeSizeX - 160, barcodeSizeY,
			1100 + 160, 416, barcodeSizeX - 160, barcodeSizeY,
			0,
		},
	}

	outputWand := imagick.NewMagickWand()
	defer outputWand.Destroy()
	drawingWand := imagick.NewDrawingWand()
	defer drawingWand.Destroy()
	pixelWand := imagick.NewPixelWand()
	defer pixelWand.Destroy()

	// Set canvas size for the output image
	pixelWand.SetColor("white")
	drawingWand.SetFillColor(pixelWand)
	err = outputWand.NewImage(canvasWidth, canvasHeight, pixelWand)
	if err != nil {
		return err
	}

	for i, section := range sections {
		cropped := srcWand.Clone()
		defer cropped.Destroy()

		if err := cropped.CropImage(uint(section.srcW), uint(section.srcH), int(section.srcX), int(section.srcY)); err != nil {
			return fmt.Errorf("failed to crop section %d: %v", i, err)
		}

		if section.rot != 0 {
			err = cropped.RotateImage(pixelWand, float64(section.rot))
			if err != nil {
				return fmt.Errorf("failed to rotate section %d: %v", i, err)
			}
		}
		// Resize cropped section (if needed)
		if err := cropped.ResizeImage(uint(section.tgtW), uint(section.tgtH), imagick.FILTER_LANCZOS, 1); err != nil {
			return fmt.Errorf("failed to resize section %d: %v", i, err)
		}

		// Composite the cropped section onto the output canvas
		if err := outputWand.CompositeImage(cropped, imagick.COMPOSITE_OP_COPY, int(section.tgtX), int(section.tgtY)); err != nil {
			return fmt.Errorf("failed to composite section %d: %v", i, err)
		}
	}
	{ // Line 1
		strokeWand := imagick.NewPixelWand()
		defer strokeWand.Destroy()
		strokeWand.SetColor("black")
		lineWand := imagick.NewDrawingWand()
		defer lineWand.Destroy()

		lineWand.SetStrokeColor(strokeWand)
		lineWand.Line(910, 0, 910, float64(canvasHeight))
		lineWand.Line(920, 230, 1070, 230)
		lineWand.Line(970, 240, 970, float64(canvasHeight))
		lineWand.Line(1080, 0, 1080, float64(canvasHeight))

		err = outputWand.DrawImage(lineWand)
		if err != nil {
			return err
		}
	}

	err = outputWand.SetImageFormat("png")
	if err != nil {
		return err
	}
	err = outputWand.WriteImage(outputFile)
	if err != nil {
		return err
	}
	//
	return nil
}
