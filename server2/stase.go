package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"slices"
	"time"
)

type Information struct {
	Name   string
	Len    int
	Max    float64
	Min    float64
	Mean   float64
	StdDev float64
}

func Procces(file string, f2 int, s []float64) Information {
	c := Information{}
	c.Name = file
	c.Len = len(s)
	c.Max = slices.Max(s)
	c.Min = slices.Min(s)
	meanValue, standardDeviation := stdDev(s)
	c.Mean = meanValue
	c.StdDev = standardDeviation

	return c
}

func stdDev(s []float64) (float64, float64) {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	mean := sum / float64(len(s))

	var squared float64
	for i := 0; i < len(s); i++ {
		squared += math.Pow(s[i]-mean, 2)
	}
	standardDeviation := math.Sqrt(squared / float64(len(s)))
	return mean, standardDeviation

}

var JSONFILE = "data.json"

type PhoneBook []Information

var data = PhoneBook{}
var index map[string]int

func Des(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

func S(s interface{}, r io.Writer) error {
	e := json.NewEncoder(r)
	return e.Encode(s)
}

func saveJSON(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = S(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func readJSON(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
			return nil
		}
		return err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = Des(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		index[k.Name] = i
	}
}

func insert(ps *Information) error {
	_, ok := index[(*ps).Name]
	if ok {
		return errors.New("Duplicate name")
	}
	data = append(data, *ps)
	createIndex()
	err := saveJSON(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}

func DElete(key string) error {
	i, ok := index[key]
	if !ok {
		return errors.New("Not Found")

	}
	data = append(data[:i], data[i+1:]...)
	delete(index, key)
	err := saveJSON(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) (*Information, error) {
	i, ok := index[key]
	if !ok {
		return nil, errors.New("Not Found")
	}
	return &data[i], nil
}

func list() string {
	var all string
	for _, k := range data {
		all += all + fmt.Sprintf("%s\t%d\t%f\t%f\n", k.Name, k.Len, k.StdDev, k.Mean)
	}
	return all
}

func main() {
	err := readJSON(JSONFILE)
	if err != nil && err != io.EOF {
		fmt.Println("Error", err)
		return
	}
	createIndex()

}
