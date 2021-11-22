#!/usr/bin/env bash

package=$1

VERSION=$CI_COMMIT_TAG
BUILDTIME=$(date +"%Y.%m.%d.%H%M%S")
REVISION=$(git log --pretty=format:"%h" -n 1)
BRANCH=$(git rev-parse --abbrev-ref HEAD)

if [ -z "$VERSION" ]; then
    VERSION=$BRANCH'-'$REVISION
fi

if [[ -z "$package" ]]; then
  echo "Usage: $0 <package-name>"
  exit 1
fi

package_split=(${package//\// })
package_name=${package_split[-1]}
echo "Package name: ${package_name}"

platforms=("windows/amd64" "linux/amd64")
for platform in "${platforms[@]}"; do
  platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}
  output_name=$package_name'-'$GOOS'-'$GOARCH
  if [ $GOOS = "windows" ]; then
    output_name+='.exe'
  fi

  # env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
  env GOOS=$GOOS GOARCH=$GOARCH go build \
  -ldflags "-extldflags '-static' \
  -X gospl/cmd.VERSION=$VERSION \
  -X gospl/cmd.BUILT=$BUILDTIME \
  -X gospl/cmd.REVISION=$REVISION \
  -X gospl/cmd.BRANCH=$BRANCH \
  " \
  -o ./artifacts/$output_name $package_name
  if [ $? -ne 0 ]; then
    echo 'Error occured'
    exit 1
  fi
  echo "Output name: ${output_name}"
done
