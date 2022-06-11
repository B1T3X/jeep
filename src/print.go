package main

import (
    "os"
    "bytes"
    "fmt"
    "github.com/hennedo/escpos"
    qrcode "github.com/skip2/go-qrcode"
    "image"
    // Get jpeg and png support for images
    _ "image/png"
    _ "image/jpeg"
    //"golang.org/x/image/webp"
)

// The escpos library function for printing QR codes doesn't for me so the QR
// code is instead implemented as an image
func createQRCode(url string) (png []byte, err error) {
    png, err = qrcode.Encode(url, qrcode.Medium, 256)
    return
}

func getLogo (logoPath string) (logoImage image.Image, err error) {
//    file, err := os.OpenFile(logoPath, os.O_RDONLY, 0)
//    if err != nil {
//        panic(err)
//    }
    //logoImage, _, err = image.Decode(file)
    return
}

// Takes a jiraIssue struct and prints to the printer connected to printerPath
func (conf *jeepConfig) printIssue(issue jiraIssue) (err error) {
    if err != nil {
        panic(err)
    }
    issue_qrcode, err := createQRCode(fmt.Sprintf("%v/browse/%v", conf.JiraAddress, issue.Key))
    if err != nil {
        panic(err)
    }
    issue_qrcode_reader := bytes.NewReader(issue_qrcode)
    image, _, err := image.Decode(issue_qrcode_reader)
    if err != nil {
        fmt.Println("Error decoding image...")
        panic(err)
    }
    file, err := os.OpenFile(conf.PrinterPath, os.O_RDWR, 0)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    printer := escpos.New(file)
    printer.Justify(escpos.JustifyCenter).Underline(2).Bold(true).Size(2, 2).Write(fmt.Sprintf("New %v\n", issue.Type))
    if err = printer.Print(); err != nil {
        panic(err)
    }
    printer.Justify(escpos.JustifyCenter).Underline(0).Bold(true).Size(1, 1).Write(fmt.Sprintf("%v\n",issue.Key))
    if err = printer.Print(); err != nil {
        panic(err)
    }
    printer.Justify(escpos.JustifyCenter).Underline(0).Bold(false).Size(1, 1).Write(fmt.Sprintf("%v\n",issue.Description))
    if err = printer.Print(); err != nil {
        panic(err)
    }
    printer.Justify(escpos.JustifyCenter).Underline(0).Bold(true).Size(1, 1).Write("Reporter: ")
    printer.Bold(false).Write(fmt.Sprintf("%v\n", issue.Reporter))
    if escpos_int, err := printer.PrintImage(image); err != nil {
       fmt.Println("Error printing image...")
        fmt.Println(escpos_int)
        panic(err)
    }
    // Line feeds so that the QR code will fully come out of the printer
    for i := 0; i < 5; i++ {
        printer.Write("\n")
        err = printer.Print()
    }
    return
}

