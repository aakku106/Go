package main

import (
	"fmt"
)

type PriceCalculator interface {
	CalculatePrice() int64
}
type ProductSetter interface {
	SetProduct()
}
type ProductGetter interface {
	GetProduct()
}

type ProductDetails struct {
	price int64
	brand string
}

type ShirtStruct struct {
	ProductDetails
	size  string
	color string
}

type MonitorStruct struct {
	ProductDetails
	size       string
	resolution string
}

type PantStruct struct {
	ShirtStruct
}
type CatStruct struct {
	ProductDetails
	color  string
	weight int8
	age    int8
}

var (
	discountOnShirt   int8 = 2
	discountOnPant    int8 = 1
	discountOnMonitor int8 = 5
	taxOnCat          int8 = 6
	VAT               int8 = 13
)

func (s *ShirtStruct) SetProduct() {
	s.brand = "aakkuWares"
	s.price = 100
	s.color = "black"
	s.size = "XXL"
}
func (s *ShirtStruct) GetProduct() *ShirtStruct {
	return s
}
func (shirt *ShirtStruct) CalculatePrice() int64 {
	discount := (shirt.price / int64(discountOnShirt)) / 100
	vat := (shirt.price / int64(VAT)) / 100
	return (shirt.price - discount + vat)
}

func (s *PantStruct) SetProduct() {
	s.brand = "aakkuWares"
	s.price = 150
	s.color = "Dark Blue"
	s.size = "XXL"
}
func (s *PantStruct) GetProduct() *PantStruct {
	return s
}
func (pant *PantStruct) CalculatePrice() int64 {
	discount := (pant.price / int64(discountOnPant)) / 100
	vat := (pant.price / int64(VAT)) / 100
	return (pant.price - discount + vat)
}

func (s *MonitorStruct) SetProduct() {
	s.brand = "aakkuMoni"
	s.price = 20_000
	s.size = "umm Well its big"
	s.resolution = "1080p"
}
func (s *MonitorStruct) GetProduct() *MonitorStruct {
	return s
}
func (m *MonitorStruct) CalculatePrice() int64 {
	discount := (m.price / int64(discountOnMonitor)) / 100
	vat := (m.price / int64(VAT)) / 100
	return (m.price - discount + vat)
}

func (s *CatStruct) SetProduct() {
	s.brand = "aakkuNamels"
	s.price = 1000
	s.age = 3
	s.weight = 9
}
func (s *CatStruct) GetProduct() *CatStruct {
	return s
}
func (cat *CatStruct) CalculatePrice() int64 {
	tax := (cat.price / int64(taxOnCat)) / 100
	vat := (cat.price / int64(VAT)) / 100
	return (cat.price - tax + vat)
}

func SetAll() (a [4]PriceCalculator) {
	s := &ShirtStruct{}
	s.SetProduct()

	p := &PantStruct{}
	p.SetProduct()

	m := &MonitorStruct{}
	m.SetProduct()

	c := &CatStruct{}
	c.SetProduct()

	a[0] = s
	a[1] = p
	a[2] = m
	a[3] = c

	return
}

// func GetAll() ProductSetter

func FinalPriceCalculator(n ...PriceCalculator) int64 {
	var s int64 = 0
	for _, v := range n {
		s += v.CalculatePrice()
	}
	return s
}

func main() {
	fmt.Println("Starting program")
	fmt.Println("Setting all")
	structsB := SetAll()
	GT := FinalPriceCalculator(structsB[0], structsB[1], structsB[2], structsB[3])
	fmt.Printf("GT: %v\n", GT)
	fmt.Printf("structsB: %v\n", structsB)
}
