#!/bin/bash

# Input and output file paths
input_file="./swagger.yaml"
output_file="./deploy/api-gateway-swagger.yml"

rm $output_file 2> /dev/null || true
sed '/example:/d' "$input_file" > "$output_file"
