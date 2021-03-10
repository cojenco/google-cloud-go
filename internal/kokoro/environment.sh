#!/bin/bash
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# A suite of tests offered by https://github.com/googleapis/env-tests-logging
# That allows deploying and testing features in live GCP services.
# Currently only configured to test logging & error reporting

set -eox pipefail

# Test prechecks
if [[ -z "${ENVIRONMENT:-}" ]]; then
  echo "ENVIRONMENT not set. Exiting"
  exit 1
fi

if [[ -z "${PROJECT_ROOT:-}"  ]]; then
    PROJECT_ROOT="github/google-cloud-go"
fi

# Add the test module as a submodule to the repo.
cd "${KOKORO_ARTIFACTS_DIR}/github/google-cloud-go/internal/"
git submodule add https://github.com/googleapis/env-tests-logging
cd "env-tests-logging/"

# Disable buffering, so that the logs stream through.
# TODO double check if I need this
export PYTHONUNBUFFERED=1

# Debug: show build environment
env | grep KOKORO

# Set up service account credentials
# TODO(nicolezhu) Check with Go team to use dulcet-port-762 project instead
# export GOOGLE_APPLICATION_CREDENTIALS=$KOKORO_KEYSTORE_DIR/72523_go_integration_service_account
export PROJECT_ID="log-bench"
# gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS

# Remove old nox
python3 --version
python3 -m pip uninstall --yes --quiet nox-automation
# Install nox
python3 -m pip install --upgrade --quiet nox
python3 -m nox --version

# create a unique id for this run
UUID=$(python  -c 'import uuid; print(uuid.uuid1())' | head -c 7)
export ENVCTL_ID=ci-$UUID
echo $ENVCTL_ID

# Run the specified environment test
set +e
python3 -m nox --session "tests(language='go', platform='$ENVIRONMENT')"
TEST_STATUS_CODE=$?

# destroy resources
echo "cleaning up..."
${PROJECT_ROOT}/tests/environment/envctl/envctl python $ENVIRONMENT destroy

# exit with proper status code
exit $TEST_STATUS_CODE
