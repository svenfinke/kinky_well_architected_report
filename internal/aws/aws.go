package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"log"
)

var aws_config aws.Config
var client wellarchitected.Client

func createConfigAndClient() {
	if &client != nil {
		return
	}

	var err error
	aws_config, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client = *wellarchitected.NewFromConfig(aws_config)
}

func FetchLenses() (*wellarchitected.ListLensesOutput, error) {
	createConfigAndClient()
	return client.ListLenses(context.TODO(), &wellarchitected.ListLensesInput{
		LensStatus: "ALL",
		LensType:   "AWS_OFFICIAL",
		MaxResults: 10,
	})
}

func FetchWorkloads() (*wellarchitected.ListWorkloadsOutput, error) {
	createConfigAndClient()
	return client.ListWorkloads(context.TODO(), &wellarchitected.ListWorkloadsInput{
		MaxResults: 10,
	})
}

func FetchImprovements(workloadId *string, lensAlias *string) (*wellarchitected.ListLensReviewImprovementsOutput, error) {
	createConfigAndClient()
	return client.ListLensReviewImprovements(context.TODO(), &wellarchitected.ListLensReviewImprovementsInput{
		LensAlias:  lensAlias,
		WorkloadId: workloadId,
	})
}
