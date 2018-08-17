#!/usr/bin/env bash

set -eou pipefail
#set -x  # useful for debugging

docker_cleanup() {
    echo "cleaning up existing network and containers..."
    CONTAINERS='document'
    docker ps | grep -E ${CONTAINERS} | awk '{print $1}' | xargs -I {} docker stop {} || true
    docker ps -a | grep -E ${CONTAINERS} | awk '{print $1}' | xargs -I {} docker rm {} || true
    docker network list | grep ${CONTAINERS} | awk '{print $2}' | xargs -I {} docker network rm {} || true
}

# optional settings (generally defaults should be fine, but sometimes useful for debugging)
DOCUMENT_LOG_LEVEL="${DOCUMENT_LOG_LEVEL:-INFO}"  # or DEBUG
DOCUMENT_TIMEOUT="${DOCUMENT_TIMEOUT:-5}"  # 10, or 20 for really sketchy network

# local and filesystem constants
LOCAL_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# container command constants
DOCUMENT_IMAGE="gcr.io/elixir-core-prod/document:snapshot" # develop

echo
echo "cleaning up from previous runs..."
docker_cleanup

echo
echo "creating document docker network..."
docker network create document

# TODO start and healthcheck dependency services if necessary

echo
echo "starting document..."
port=10100
name="document-0"
docker run --name "${name}" --net=document -d -p ${port}:${port} ${DOCUMENT_IMAGE} \
    start \
    --logLevel "${DOCUMENT_LOG_LEVEL}" \
    --serverPort ${port}
    # TODO add other relevant args if necessary
document_addrs="${name}:${port}"
document_containers="${name}"

echo
echo "testing document health..."
docker run --rm --net=document ${DOCUMENT_IMAGE} test health \
    --addresses "${document_addrs}" \
    --logLevel "${DOCUMENT_LOG_LEVEL}"

echo
echo "testing document ..."
docker run --rm --net=document ${DOCUMENT_IMAGE} test io \
    --addresses "${document_addrs}" \
    --logLevel "${DOCUMENT_LOG_LEVEL}"
    # TODO add other relevant args if necessary

echo
echo "cleaning up..."
docker_cleanup

echo
echo "All tests passed."
