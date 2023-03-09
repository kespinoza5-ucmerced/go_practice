package main

import (
	"flag"

	"github.com/kespinoza5-ucmerced/go_practice/test_experiment/runner"
	"gitlab.com/akita/mgpusim/v3/benchmarks/dnn/relu"
)

var numData = flag.Int("length", 4096, "The number of samples to filter.")

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := relu.NewBenchmark(runner.Driver())
	benchmark.Length = *numData

	runner.AddBenchmark(benchmark)

	runner.Run()
}
