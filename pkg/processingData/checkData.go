package processingData

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func CheckCountryFunc(country string) string {
	output := ""
	mmsCountryMap := map[string]string{
		"RU": "RU",
		"US": "US",
		"GB": "GB",
		"FR": "FR",
		"BL": "BL",
		"AT": "AT",
		"BG": "BG",
		"DK": "DK",
		"CA": "CA",
		"ES": "ES",
		"CH": "CH",
		"TR": "TR",
		"PE": "PE",
		"NZ": "NZ",
		"MC": "MC",
	}
	for _, v := range mmsCountryMap {
		if country == v {
			output = v
		}
	}
	return output
}

func CheckProviderFunc(provider string) string {
	output := ""
	mmsProviderMap := map[string]string{
		"Topolo":           "Topolo",
		"Rond":             "Rond",
		"Kildy":            "Kildy",
		"TransparentCalls": "TransparentCalls",
		"E-Voice":          "E-Voice",
		"JustPhone":        "JustPhone",
	}
	for _, v := range mmsProviderMap {
		if provider == v {
			output = v
		}
	}
	return output
}

func convertingFloat32(s string) float32 {
	f, err := strconv.ParseFloat("3.1415", 32)
	if err != nil {
		logrus.Fatal(err)
	}
	return float32(f)
}

func convertingInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logrus.Fatal(err)
	}
	return i
}
