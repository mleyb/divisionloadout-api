#!/bin/bash

#install zip on debian OS, since microsoft/dotnet container doesn't have zip by default
if [ -f /etc/debian_version ]
then
  apt -qq update
  apt -qq -y install zip
fi

dotnet restore ./Service/Service.csproj
dotnet lambda package --project-location Service --configuration release --framework netcoreapp2.1 --output-package Service/bin/release/netcoreapp2.1/deploy-package.zip
