
# Homework Policy Verifier

## Overview
This project is a Homework Policy Verifier tool developed in Go, designed to check AWS IAM Role Policies against specified guidelines. It evaluates policies to identify if they are assigned to specific resources or use wildcards, based on the structure defined in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html).

## Features
- **Resource Verification**: Determines whether IAM Role Policies are assigned to specific resources or utilize wildcards.
- **Compliance Checking**: Ensures policies comply with defined practices.

## Prerequisites
Before you begin, ensure you have met the following requirement:
- Go (version 1.22.1) installed on your machine.

## Installation
Clone the repository to your local machine:
```bash
git clone https://github.com/Aros24/HomeworkPolicyVerifier.git
cd HomeworkPolicyVerifier
```
Install the necessary dependencies:
```bash
go mod tidy
```
## Usage
To use the Homework Policy Verifier, run the tool with the -input flag followed by the path to your policy JSON file:
```bash
go run main.go -input "/path/to/your/policy/file.json"
```
Ensure to replace "/path/to/your/policy/file.json" with the actual path to the IAM policy JSON file you wish to verify. You can find examplery JSON file in folder examples.
## Running Tests
This project includes tests to ensure the reliability and correctness of the policy verification logic. To run these tests, execute the following command in the project root directory:
```bash
go test ./... -v
```
This will run all tests within the project and output the results, including any failures or errors encountered during testing.