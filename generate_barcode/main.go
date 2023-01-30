package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	//"os"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

func main() {

	now := time.Now()
	expiry_date_string := now.Format("2006-01-02 15:04:05")
	date, err := time.Parse("2006-01-02 15:04:05", expiry_date_string)
    if err != nil {
        fmt.Println("           [*] ERROR while parsing datetime", err)
    }
	expiryDate := date.Add(time.Minute * 2).Format("2006-01-02 15:04:05")
	fmt.Println(expiryDate)

	// Create the qrcode
	qrCode, _ := qr.Encode(expiryDate, qr.M, qr.Auto)
	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 500, 500)

	// barcode
	enc := oned.NewCode128Writer()
	img, _ := enc.Encode("Hello, Gophers! bbbbbh", gozxing.BarcodeFormat_CODE_128, 400, 200, nil)

	// *BitMatrix implements the image.Image interface,
	// so it is able to be passed to png.Encode directly.

	// http request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data1 := make(map[string]interface{})
		bpng, _ := DecodePNGBase64(img)
		bjpeg, _ := DecodeJPEGBase64(img)
		qpng, _ := DecodePNGBase64(qrCode)
		qjpeg, _ := DecodeJPEGBase64(qrCode)
		data1["barcode"] = map[string]interface{} {
			"base64_png": bpng,
			"base64_jpeg": bjpeg,
		}
		data1["qr_code"] = map[string]interface{} {
			"base64_png": qpng,
			"base64_jpeg": qjpeg,
		}
	
		d, _ := json.Marshal(data1)
		fmt.Println(string(d))
		json.NewEncoder(w).Encode(data1)
		// need to decode the string in base64 to convert to an image
	})
	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	})
	http.HandleFunc("/qr_code", func(w http.ResponseWriter, r *http.Request) {
		jpeg.Encode(w, qrCode, nil)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func DecodeJPEGBase64(img image.Image) (string, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}	

func DecodePNGBase64(img image.Image) (string, error) {
	buff := new(bytes.Buffer)
	if err := png.Encode(buff, img); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buff.Bytes()), nil
}