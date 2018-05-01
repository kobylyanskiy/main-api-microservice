DOCKER=docker
KUBECTL=kubectl
IMAGE_LOCAL_PREFIX=kobylyanskiy
DEPLOYMENT=main-api
VERSION=$1
PROJECT_ID=spy-crowd

${DOCKER} build -t ${IMAGE_LOCAL_PREFIX}/${DEPLOYMENT}:${VERSION} .
${DOCKER} tag ${IMAGE_LOCAL_PREFIX}/${DEPLOYMENT}:${VERSION} gcr.io/${PROJECT_ID}/${DEPLOYMENT}:${VERSION}
${DOCKER} push gcr.io/${PROJECT_ID}/${DEPLOYMENT}:${VERSION}

${KUBECTL} set image deployment/${DEPLOYMENT}-deployment ${DEPLOYMENT}=gcr.io/${PROJECT_ID}/${DEPLOYMENT}:${VERSION}
${KUBECTL} rollout status deployment/${DEPLOYMENT}-deployment
