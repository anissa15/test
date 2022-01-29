package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	r := mux.NewRouter()
	r.Methods("POST").Path("/convert").Handler(http.HandlerFunc(handleConvert))

	log.Println("Application is starting at port", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, r))
}

type ConvertRequest struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ConvertResponse struct {
	Result string `json:"result"`
}

func handleConvert(w http.ResponseWriter, r *http.Request) {
	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "failed decode request")
		return
	}

	var resp ConvertResponse
	switch req.Type {
	case "binary":
		match, err := validateBinary(req.Value)
		if err != nil {
			writeError(w, "failed validate binary value")
			return
		}
		if !match {
			writeError(w, "invalid binary value")
			return
		}
		decimal, err := convertBinary(req.Value)
		if err != nil {
			writeError(w, "failed convert binary")
			return
		}
		resp.Result = fmt.Sprint(decimal)
	case "decimal":
		decimal, err := validateDecimal(req.Value)
		if err != nil {
			writeError(w, "failed validate decimal value")
			return
		}
		resp.Result = convertDecimal(decimal)
	default:
		writeError(w, "type option are 'binary' or 'decimal'")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func validateBinary(value string) (bool, error) {
	matchOtherThan01, err := regexp.MatchString("([^0-1])+", value)
	if err != nil {
		return false, err
	}
	return !matchOtherThan01, nil
}

func validateDecimal(value string) (int, error) {
	decimal, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

func convertDecimal(decimal int) string {
	var binary []int
	for decimal != 0 {
		binary = append(binary, decimal%2)
		decimal = decimal / 2
	}
	var result string
	if len(binary) == 0 {
		result = fmt.Sprint(0)
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			result += fmt.Sprint(binary[i])
		}
	}
	return result
}

func convertBinary(binary string) (int, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(decimal), nil
}

func writeError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
