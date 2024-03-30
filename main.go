package main

import (
	"flag"
	"homeworkPolicyVerifier/policy"
	"homeworkPolicyVerifier/utils"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("No flags provided. Please use the -input option to specify a JSON file path.")
		os.Exit(1)
	}

	inputFlag := flag.String("input", "", "Path to a JSON file")
	flag.Parse()

	if *inputFlag == "" {
		log.Println("No input provided. Please use the -input option to specify a JSON file path.")
		os.Exit(1)
	}

	policyData, err := utils.ReadInput(*inputFlag)
	if err != nil {
		log.Printf("Failed to read input: %v\n", err)
		os.Exit(1)
	}

	parsedPolicy, err := policy.UnmarshalPolicy(policyData)
	if err != nil {
		log.Printf("Failed to parse policy: %v\n", err)
		os.Exit(1)
	}

	if err := parsedPolicy.ValidatePolicy(); err != nil {
		log.Printf("Policy validation failed: %v\n", err)
		os.Exit(1)
	}

	if parsedPolicy.PolicyDocument.HasSpecificResources() {
		log.Println("The policy specifies specific resources for all statements.")
	} else {
		log.Println("At least one statement in the policy specifies all resources using '*'.")
	}
}
