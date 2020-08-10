package main

import (
	"log"
	"os"

	_ "github.com/horance-liu/gohit/matcher"
	"github.com/horance-liu/gohit/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
