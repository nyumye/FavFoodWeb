#!/bin/sh

# m is for checking the need for mongod
m="m"

# start mongo db
if [ "$1" = "$m" ]
then
    start cmd.exe "/C bash -c 'mongod --config ./mongo/mongodb.config'"
fi


# build go file
cd ./FavFoodWeb
go build

# run main app
./FavFoodWeb
# cmd.exe "/K bash -c './FavFoodWeb'"

cd ..