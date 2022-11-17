package renderer

import (
	"fmt"
	"github.com/svenfinke/kinky_well_architected_report/internal/aws"
	"log"
)

func PrintData() {
	lenses, err := aws.FetchLenses()
	if err != nil {
		log.Fatal(err)
	}

	print("\n Found the following lenses: \n")
	for _, lens := range lenses.LensSummaries {
		fmt.Printf(" - %s (%s:%s) \n", *lens.LensAlias, *lens.LensName, *lens.LensVersion)
	}

	workloads, err := aws.FetchWorkloads()
	if err != nil {
		log.Fatal(err)
	}

	print("\n Found the following workloads: \n")
	for _, workload := range workloads.WorkloadSummaries {
		fmt.Printf(" - %s (%s)\n", *workload.WorkloadName, *workload.WorkloadId)
	}

	for _, workload := range workloads.WorkloadSummaries {
		for _, lens := range lenses.LensSummaries {
			improvements, err := aws.FetchImprovements(workload.WorkloadId, lens.LensAlias)
			if err != nil {
				log.Fatal(err)
			}

			print("\n Printing all improvements: \n")
			for _, improvement := range improvements.ImprovementSummaries {
				fmt.Printf(" - %s \n", *improvement.QuestionTitle)
			}
		}
	}
}
