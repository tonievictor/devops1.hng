package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors"
	"github.com/tonievictor/dotenv"
)

type Response struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

type ErrResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func main() {
	dotenv.Config(".env")

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/classify-number", handleroot)
	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		fmt.Println("An error occured while starting up the server")
	}
}

func handleroot(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	n, err := strconv.Atoi(number)
	if err != nil {
		res := ErrResponse{
			Number: number,
			Error:  true,
		}
		fmt.Printf("An error occured %v", err)
		writeRes(w, http.StatusBadRequest, res)
		return
	}

	funfact, err := GetFunFact(n)
	if err != nil {
		res := ErrResponse{
			Number: number,
			Error:  true,
		}
		fmt.Printf("An error occured %v", err)
		writeRes(w, http.StatusBadRequest, res)
		return
	}

	res := Response{
		Number:     n,
		IsPrime:    IsPrimeFn(n),
		IsPerfect:  IsPerfectFn(n),
		DigitSum:   SumOfDigits(n),
		Properties: IsArmStrongFn(n),
		FunFact:    funfact,
	}
	writeRes(w, http.StatusOK, res)
}

func writeRes(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Printf("An error occured while encoding the json response: %v", err)
	}
}
