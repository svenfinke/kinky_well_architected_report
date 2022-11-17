package main

import (
	"flag"
	"github.com/svenfinke/kinky_well_architected_report/internal/renderer"
)

var (
	printList   bool
	workloadId  string
	lensAliases []string
)

func main() {
	print(`

 █  █▀ ▄█    ▄   █  █▀ ▀▄    ▄
 █▄█   ██     █  █▄█     █  █
 █▀▄   ██ ██   █ █▀▄      ▀█
 █  █  ▐█ █ █  █ █  █     █
   █    ▐ █  █ █   █    ▄▀
  ▀       █   ██  ▀

  Rendering AWS Well Architected Reports as kinky LaTeX documents  
	`)

	printList = *flag.Bool("print", false, "print the available data instead of rendering the LaTeX document")
	workloadId = *flag.String("workloadId", "", "the workload id from which the LaTeX document should be created.")
	lensAliases = flag.Args()

	if printList {
		renderer.PrintData()
		return
	}
}
