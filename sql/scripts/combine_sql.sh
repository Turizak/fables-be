#!/bin/bash

# Check if enough parameters are passed
if [[ -z "$1" || -z "$2" ]]; then
  echo "Usage: $0 <directory> <output_filename>"
  exit 1
fi

# Input directory containing SQL files
input_dir="$1"

# Output file for the combined SQL script
output_file="$2"

# Check if the directory exists
if [[ ! -d "$input_dir" ]]; then
  echo "Error: Directory '$input_dir' does not exist."
  exit 1
fi

# Create or overwrite the output file
> "$output_file"

# Loop through all .sql files in the directory
for file in "$input_dir"/*.sql; do
  if [[ -f "$file" ]]; then
    echo "-- Begin $(basename "$file")" >> "$output_file"
    cat "$file" >> "$output_file"
    echo "" >> "$output_file"
    echo "-- End $(basename "$file")" >> "$output_file"
    echo "" >> "$output_file"
  else
    echo "Warning: No .sql files found in '$input_dir'."
  fi
done

echo "Combined SQL script created as $output_file."
