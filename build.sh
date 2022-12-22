#!/bin/bash

make tfgen
make build_sdks
make install_nodejs_sdk

rm -f ~/.pulumi/bin/pulumi-resource-statuscake
cp ./bin/pulumi-resource-statuscake ~/.pulumi/bin/pulumi-resource-statuscake

cd examples/my-example/ts
yarn link "@pulumiverse/statuscake"
