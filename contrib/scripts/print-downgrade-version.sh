#!/usr/bin/env bash
#
# A utility script to print the branch name of the previous stable or patch
# release.

set -o errexit
set -o nounset

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
VERSION=${VERSION-"$(cat "$SCRIPT_DIR/../../VERSION")"}
if [[ $VERSION =~ ([0-9^]+)\.([0-9^]+)\.([0-9^]+).* ]] ; then
    major=${BASH_REMATCH[1]}
    minor=${BASH_REMATCH[2]}
    patch=${BASH_REMATCH[3]}
else
  >&2 echo "ERROR: failed to parse version '$VERSION'"
  exit 1
fi

if [[ ${1-} == "patch" ]] ; then
    # If user passed "patch" as first argument, print the latest patch version

    # Hack: When working on a patch release preparation commit, file VERSION
    # contains the new value for the release that is yet to be tagged and
    # published. In this case, we want to downgrade to the previous patch
    # release.
    # Skip this step if we're not in a Git repository or if there is no more
    # than one commit (likely a shallow clone).
    if git rev-parse --is-inside-work-tree &> /dev/null && \
        git rev-parse --verify HEAD^ &> /dev/null && \
        git diff --name-only HEAD^..HEAD | grep -q "^VERSION$"; then
        patch=$((patch - 1))
    fi
    if [[ "${patch}" -le "0" ]] ; then
        >&2 echo "ERROR: failed to deduce patch release previous to version '$VERSION'"
        exit 1
    else
        echo "v${major}.${minor}.${patch}${TAG_SUFFIX:-}"
    fi
else
    if [[ "${minor}" == "0" ]] ; then
        >&2 echo "ERROR: failed to deduce release previous to version '$VERSION'"
        exit 1
    fi
    # Else print the previous stable version by decrementing the minor version
    # and trimming the patch version.
    ((minor--))
    echo "v${major}.${minor}${BRANCH_SUFFIX:-}"
fi
