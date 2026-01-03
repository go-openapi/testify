#!/usr/bin/env bash
# Local development script for Hugo documentation

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "${SCRIPT_DIR}"

echo "==> Preparing Hugo documentation site..."

# Extract version information
LATEST_RELEASE=$(git tag --list --sort -version:refname 'v*' 2>/dev/null | head -1 || echo "dev")
REQUIRED_GO_VERSION=$(grep "^go\s" ../../../go.mod | awk '{print $2}')
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION_MESSAGE="Documentation test for latest master"

echo "    Latest release: ${LATEST_RELEASE}"
echo "    Go version: ${REQUIRED_GO_VERSION}"
echo "    Build time: ${BUILD_TIME}"

# Generate dynamic config
cat testify.yaml.template \
  | sed "s|{{ GO_VERSION }}|${REQUIRED_GO_VERSION}|g" \
  | sed "s|{{ LATEST_RELEASE }}|${LATEST_RELEASE}|g" \
  | sed "s|{{ VERSION_MESSAGE }}|${VERSION_MESSAGE}|g" \
  | sed "s|{{ BUILD_TIME }}|${BUILD_TIME}|g" \
  > testify.yaml

echo "==> Generated testify.yaml"

# Check if theme exists
if [ ! -d "themes/hugo-relearn" ]; then
  echo "ERROR: Relearn theme not found at themes/hugo-relearn"
  echo "Run: unzip hugo-theme-relearn-main.zip -d themes/ && mv themes/hugo-theme-relearn-main themes/hugo-relearn"
  exit 1
fi

# Check if generated docs exist
if [ ! -d "../../../docs/doc-site" ]; then
  echo "WARNING: Generated docs not found at ../../../docs/doc-site"
  echo "You may need to run: go generate ./..."
  echo ""
  echo "Creating placeholder content directory..."
  mkdir -p content
fi

echo "==> Starting Hugo development server..."
echo "    Visit: http://localhost:1313/testify/"
echo ""

# Start Hugo server with both configs
hugo server \
  --config hugo.yaml,testify.yaml \
  --buildDrafts \
  --disableFastRender \
  --navigateToChanged \
  --bind 0.0.0.0 \
  --port 1313 \
  --baseURL http://localhost:1313/testify/ \
  --appendPort=false \
  --logLevel info \
  --cleanDestinationDir
