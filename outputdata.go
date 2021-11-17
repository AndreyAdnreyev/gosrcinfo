package main

import (
	"flag"
	"fmt"
)

type MapData map[string][]string

func NewMapData() MapData {
	return map[string][]string{}
}

func (d MapData) print() {
	i := 1
	for k := range d {
		for _, elem := range d[k] {
			fmt.Printf("%3d : %s : %s\n", i, k, elem)
		}
		i++
	}
}

func (d MapData) add(key string, file string) {
	if _, ok := d[key]; !ok {
		d[key] = make([]string, 0)
	}
	d[key] = append(d[key], file)
}

func printSlice(input []string) {
	for i, v := range input {
		fmt.Println(i+1, ":", v)
	}
}

func printHelp() {
	flag.PrintDefaults()

}
