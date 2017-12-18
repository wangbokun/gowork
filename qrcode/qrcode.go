package qrcode

import (
	"fmt"
	"os"

	"github.com/tuotoo/qrcode"
)

func QrcodeDecode(png string) string {

	fi, err := os.Open(png)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer fi.Close()

	qrmatrix, err := qrcode.Decode(fi)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	fmt.Println(qrmatrix.Content)

	return qrmatrix.Content
}
