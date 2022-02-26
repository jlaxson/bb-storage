#!/bin/sh -ex

# if test "${GITHUB_ACTIONS}" = "true"; then
  echo "BUILD_SCM_REVISION $(git rev-parse --short HEAD)"
  echo "BUILD_SCM_TIMESTAMP $(TZ=UTC date -r "$(git show -s --format=%ct HEAD)" +%Y%m%dT%H%M%SZ)"
# fi
