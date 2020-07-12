package handlers

import (
	"encoding/json"
	"gomicroservice/model"
	"gomicroservice/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Products struct {
	l *log.Logger
}

var (
	Fields []string
	UrlLen int
)

func NewProduct(l *log.Logger) *Products  {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// handle the request for a list of products

	Fields = strings.Split(r.URL.String(), "/")
	UrlLen = len(Fields)

	if r.Method == http.MethodGet {
		switch UrlLen {
		case 2 :
			p.getProducts(w, r)
			return
		case 3 :
			p.showProduct(w)
			return
		default :
			break
		}

	}

	if r.Method == "POST" || r.Method == "OPTIONS" {
		p.insertProduct(w, r)
		return
	}

	if r.Method == "PUT" || r.Method == "OPTIONS" {
		p.updateProduct(w, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	utils.ResponseJSON(w, "Method Not Aloowed", http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

		p.l.Println("Handle GET Products")

		// fetch the products from the datastore
		lp, err := model.GetProductAll()
		if err != nil {
			p.l.Printf("%v \n", err)
		}

		utils.ResponseJSON(w, lp, http.StatusOK)
		p.l.Println(lp)
		return
}

func (p *Products) showProduct(w http.ResponseWriter) {

	p.l.Println("Handle Show Products")
	id, err := strconv.ParseInt(Fields[UrlLen-1], 10, 64)

		if err != nil {
			utils.ResponseJSON(w, "wrong endpoint", http.StatusBadRequest)
			return
		}

	lp, err := model.ShowProduct(id)
		if err != nil {
			p.l.Printf("%v \n", err)
			return
		}

		utils.ResponseJSON(w, lp, http.StatusOK)
		p.l.Println(lp)
		return
}

func (p *Products) insertProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")

	var dataproduct model.Product

	if err := json.NewDecoder(r.Body).Decode(&dataproduct);

	err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		p.l.Printf("Error : %v \n", err)
		return
	}

	if err := model.PostProduct(dataproduct);
	err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		p.l.Println(err)
		return
	}

	utils.ResponseJSON(w, dataproduct, http.StatusCreated)
	p.l.Println(dataproduct)
	return
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	var dataproduct model.Product

	if err := json.NewDecoder(r.Body).Decode(&dataproduct);
	err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		p.l.Println(err)
		return
	}

	if err := model.UpdateProduct(dataproduct);
		err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			p.l.Println(err)
			return
		}

		utils.ResponseJSON(w, dataproduct, http.StatusCreated)
		p.l.Println(dataproduct)
		return
}

