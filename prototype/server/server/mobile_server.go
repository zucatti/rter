package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var phoneIDValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	p, error := ioutil.ReadAll(r.Body)
	if error != nil {
		return
	}

	ioutil.WriteFile(imagePath+"test.png", p, 0600)
}

func MultiUploadHandler(w http.ResponseWriter, r *http.Request) {
	imageFile, header, error := r.FormFile("image")
	checkError(error)

	phoneID := r.FormValue("phone_id")

	if !phoneIDValidator.MatchString(phoneID) {
		http.Error(w, "Malformed Request: Invalid phone_id", http.StatusBadRequest)
		return
	}

	rows, _, err := database.Query("SELECT * FROM whitelist where phone_id = \"%v\"", phoneID)

	if len(rows) == 0 {
		http.Error(w, "Malformed Request: Invalid phone_id", http.StatusBadRequest)
		return
	}

	os.Mkdir(imagePath+phoneID, os.ModeDir|0755)

	lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
	checkError(err)
	long, err := strconv.ParseFloat(r.FormValue("long"), 64)
	checkError(err)

	t := time.Now()
	path := imagePath

	if strings.HasSuffix(header.Filename, ".png") {
		path += fmt.Sprintf("%v/%v.png", phoneID, t.UnixNano())
	} else if strings.HasSuffix(header.Filename, ".jpg") || strings.HasSuffix(header.Filename, "jpeg") {
		path += fmt.Sprintf("%v/%v.jpg", phoneID, t.UnixNano())
	}

	outputFile, error := os.Create(path)
	checkError(error)
	defer outputFile.Close()

	io.Copy(outputFile, imageFile)

	_, _, error = database.Query("INSERT INTO content (phone_id, filepath, geolat, geolong) VALUES(\"%s\", \"%s\", %v, %v)", phoneID, path, lat, long)
	checkError(error)
}
