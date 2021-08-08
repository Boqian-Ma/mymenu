#!/bin/bash

## declare an array variable
declare -a array=("restaurants" "tables" "restaurant_members" "categories" "menu_items" "orders" "orders_items")

# set password
set pass "password"
export PGPASSWORD=password

# get length of an array
arraylength=${#array[@]}

# Get path to data files
dot="$(cd "$(dirname "$0")"; pwd)"
path="$dot/migrations"

# use for loop to read all values and indexes
for (( i=0; i<${arraylength}; i++ ));
do
echo Loading ${array[$i]}.csv
echo $(psql -U postgres -h localhost example -c "\copy ${array[$i]} FROM '$path/${array[$i]}.csv' DELIMITER ',' CSV HEADER";)
done
