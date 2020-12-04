package main

import (
	"fmt"
	"strconv"
)

func main() {
	        // Create and register a OpenCensus Stackdriver Trace exporter.
        exporter, err := stackdriver.NewExporter(stackdriver.Options{
                ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
        })
        if err != nil {
                log.Fatal(err)
        }
        trace.RegisterExporter(exporter)

        // By default, traces will be sampled relatively rarely. To change the
        // sampling frequency for your entire program, call ApplyConfig. Use a
        // ProbabilitySampler to sample a subset of traces, or use AlwaysSample to
        // collect a trace on every run.
        //
        // Be careful about using trace.AlwaysSample in a production application
        // with significant traffic: a new trace will be started and exported for
        // every request.
        trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	a := strconv.Itoa(12)
	s := strconv.FormatInt(int64(33333), 10)
	fmt.Printf("%q %q\n", a, s)
}
