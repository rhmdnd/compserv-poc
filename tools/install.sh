#!/bin/bash

sudo dnf update -y
sudo dnf install java-latest-openjdk wget -y

# NOTE: Versions 5.3.0 and 5.4.0 are unable to validate EasyDynamic's OpenAPI
# specification for OSCAL
# (https://github.com/EasyDynamics/oscal-rest/blob/develop/openapi.yaml). This
# is likely due to something in the specification, but the validation process
# doesn't provide any debug output besides the following:
# 
#  Errors:
#          - Cannot invoke "io.swagger.v3.parser.util.OpenAPIDeserializer$ParseResult.extra(String,
#            String, com.fasterxml.jackson.databind.JsonNode)" because "result" is null
#  
#  [error] Spec has 1 errors.
#  
# Version 6.0.0 beta validates the specification. I'm leaving the variables
# here to show where I installed the CLI tool from and the versions.

GENERATOR_V5_3=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.3.0/openapi-generator-cli-5.3.0.jar
GENERATOR_V5_4=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.4.0/openapi-generator-cli-5.4.0.jar
GENERATOR_V6=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.0.0-beta/openapi-generator-cli-6.0.0-beta.jar

# Install version 6.0.0 beta since it works with the specification we're using.
wget $GENERATOR_V6
