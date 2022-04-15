#!/bin/sh

set -e

while true; do
  read -p "Insert challenge-kebab-case: " challenge
  if [[ -z $challenge ]]; then
    continue
  fi
  break
done

mkdir ${challenge}
touch ${challenge}/${challenge}.go

echo "Insert your function in ${challenge}/${challenge}.go and create new test file with:\n\tgotests --all --parallel -w ${challenge}/${challenge}.go"

exit 0