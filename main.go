package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"git.cheetah.cat/cheetah/go-snake-label/lib"
	"gopkg.in/gographics/imagick.v2/imagick"
)

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, (n+1)/2)
	if _, err := src.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)[:n]
}

func tempPNGPath() string {
	return path.Join(os.TempDir(), fmt.Sprintf(RandStringBytesMaskImprSrc(16), ".png"))
}

// sudo apt-get install libmagickwand-dev
func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	inputFile := flag.String("input", "", "input filename")
	outputFile := flag.String("output", "", "output filename")
	modeSelect := flag.String("mode", "dhlprivat", "mode select (dhlprivat, dhlprivatint)")

	flag.Parse()

	if len(*inputFile) < 3 {
		panic("invalid input file")
	}
	if len(*outputFile) == 0 {
		panic("invalid output file")
	}

	tempFileName := tempPNGPath()
	defer os.Remove(tempFileName)

	outputFileName := *outputFile
	outputSTDOUT := false
	if outputFileName == "-" || outputFileName == "/dev/stdout" {
		outputSTDOUT = true
		outputFileName = tempPNGPath()
		defer os.Remove(outputFileName)
	}
	switch *modeSelect {
	case "dhlprivat":
		// scale: 4,    // optional, defaults to 4 (288dpi) -- 4.1666 for 300dpi
		lib.ConvertPDF2PNG(*inputFile, tempFileName, 288, 90)
		lib.DHLPrivat(tempFileName, outputFileName)
	case "dhlprivatint":
		lib.ConvertPDF2PNG(*inputFile, tempFileName, 288, 90)
		lib.DHLPrivatInternational(tempFileName, outputFileName)
	}

	if outputSTDOUT {
		outputBytes, _ := os.ReadFile(outputFileName)
		binary.Write(os.Stdout, binary.LittleEndian, outputBytes)
	}

}
