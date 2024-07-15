## Step-by-Step Instructions

## 1.Building the Project
cd /path/to/project
cd cmd
go build

## 2.Running the Project
./cmd

## 3.Running Tests:

## For handler package:
cd ../pkg/handler  # Move up one directory from cmd to project root, then into handler
go test

## For service package:
cd ../service  # Move up one directory from cmd to project root, then into service
go test

## Running All Tests:
cd ..
go test ./...