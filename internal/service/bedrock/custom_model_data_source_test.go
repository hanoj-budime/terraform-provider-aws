// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrock_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccBedrockCustomModelDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_custom_model.test"
	datasourceName := "data.aws_bedrock_custom_model.test"
	var v bedrock.GetModelCustomizationJobOutput

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID) },
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomModelConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomModelExists(ctx, resourceName, &v),
				),
			},
			{
				PreConfig: func() {
					testAccWaitModelCustomizationJobCompleted(ctx, t, &v)
				},
				Config: testAccCustomModelDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "hyperparameters", datasourceName, "hyperparameters"),
					resource.TestCheckResourceAttrPair(resourceName, "job_arn", datasourceName, "job_arn"),
					resource.TestCheckResourceAttrPair(resourceName, "job_name", datasourceName, "job_name"),
					resource.TestCheckResourceAttrPair(resourceName, "model_arn", datasourceName, "model_arn"),
					resource.TestCheckResourceAttrPair(resourceName, "model_kms_key_arn", datasourceName, "model_kms_key_arn"),
					resource.TestCheckResourceAttrPair(resourceName, "model_name", datasourceName, "model_name"),
					resource.TestCheckResourceAttrPair(resourceName, "output_data_config.#", datasourceName, "output_data_config.#"),
					resource.TestCheckResourceAttrPair(resourceName, "training_data_config.#", datasourceName, "training_data_config.#"),
					resource.TestCheckResourceAttrPair(resourceName, "training_metrics.#", datasourceName, "training_metrics.#"),
					resource.TestCheckResourceAttrPair(resourceName, "validation_data_config.#", datasourceName, "validation_data_config.#"),
					resource.TestCheckResourceAttrPair(resourceName, "validation_metrics.#", datasourceName, "validation_metrics.#"),
				),
			},
		},
	})
}

func testAccCustomModelDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccCustomModelConfig_basic(rName), `
data "aws_bedrock_custom_model" "test" {
  model_id = aws_bedrock_custom_model.test.model_arn
}
`)
}
