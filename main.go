package main

import (
	"flag"
	"homeworkPolicyVerifier/policy"
	"homeworkPolicyVerifier/utils"
	"log"
	"os"
)

func main() {
	inputFlag := flag.String("input", "", "Path to a JSON file or a JSON string")
	flag.Parse()

	if *inputFlag == "" {
		log.Println("No input provided. Please use the -input option to specify a JSON file path or a JSON string.")
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

	log.Printf("Policy Name: %s\nPolicy Document: %+v\n", parsedPolicy.PolicyName, parsedPolicy.PolicyDocument)

	if parsedPolicy.PolicyDocument.HasSpecificResources() {
		log.Println("The policy specifies specific resources for all statements.")
	} else {
		log.Println("At least one statement in the policy specifies all resources using '*'.")
	}
}
