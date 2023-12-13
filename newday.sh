#!/bin/bash

# Copy the Day Template Directory 
cp -r day_template day$1

# Rename the dayX.go & dayX_test.go files.
mv day$1/dayX_test.go day$1/day$1_test.go
mv day$1/dayX.go day$1/day$1.go

# Replace the Package name to the new day in the implementation file.
sed -i '' "s/dayX/day$1/g" day$1/day$1.go
# Replace the Package name to the new day in the test file.
sed -i '' "s/dayX/day$1/g" day$1/day$1_test.go
# Replace the function naming to include the new day.
sed -i '' "s/DayX/Day$1/g" day$1/day$1_test.go
# Replace the Test Error strings with the correct day.
sed -i '' "s/Day X/Day $1/g" day$1/day$1_test.go

# Calculate the day before (day minus 1)
DAY_BEFORE="day$(($1-1))"
# Add the new day's Run function to the array of Run functions in the main.go file.
sed -i '' "s/$DAY_BEFORE\.Run,/$DAY_BEFORE\.Run,\n\t\tday$1.Run,/g" main.go
# Add the import of the new day's package in the main.go file. (Find the previous day import and replace it with that plus the new day).
sed -i '' "s/\"github.com\/mikeconroy\/advent-of-code-23\/$DAY_BEFORE\"/\"github.com\/mikeconroy\/advent-of-code-23\/$DAY_BEFORE\"\n\t\"github.com\/mikeconroy\/advent-of-code-23\/day$1\"/g" main.go