package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {

		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL: ¨%v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		var Addres Cep
		err = json.Unmarshal(res, &Addres)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshalling JSON: %v\n", err)
			continue
		}

		println(Addres.Cep, Addres.Logradouro, Addres.Bairro, Addres.Localidade, Addres.Uf)

		file, err := os.Create("address.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\n",
			Addres.Cep, Addres.Logradouro, Addres.Bairro, Addres.Localidade, Addres.Uf))
	}
}
