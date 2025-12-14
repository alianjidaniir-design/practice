package main

import (
	"encoding/json"
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

func Procces(file string, s []float64) Information {
	c := Information{}
	c.Name = file
	c.Len = len(s)
	c.Max = slices.Max(s)
	c.Min = slices.Min(s)
	meanValue, standardDeviation := stdDev(s)
	c.Mean = meanValue
	c.StdDev = standardDeviation
}
