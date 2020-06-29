#!/bin/bash

# https://gist.github.com/davemo/88de90577a57698dd72d722bcfc44964
# works with a file called VERSION in the current directory,
# the contents of which should be a semantic version number
# such as "1.2.3"

NOW="$(date +'%B %d, %Y')"
RED="\033[1;31m"
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
BLUE="\033[1;34m"
CYAN="\033[1;36m"
WHITE="\033[1;37m"
RESET="\033[0m"

LATEST_HASH=`git log --pretty=format:'%h' -n 1`

WARNING_FLAG="${YELLOW}!"
NOTICE_FLAG="${CYAN}❯"

PUSHING_MSG="${NOTICE_FLAG} Pushing new version to the ${WHITE}origin${CYAN}..."

if [ -f VERSION ]; then
  BASE_STRING=`cat VERSION`
  BASE_LIST=(`echo $BASE_STRING | tr '.' ' '`)
  V_MAJOR=${BASE_LIST[0]}
  V_MINOR=${BASE_LIST[1]}
  V_PATCH=${BASE_LIST[2]}
  echo -e "${NOTICE_FLAG} Current version: ${WHITE}$BASE_STRING"
  echo -e "${NOTICE_FLAG} Latest commit hash: ${WHITE}$LATEST_HASH"

  if [[ $1 -eq '--patch' ]]; then
      NEW_VERSION="$V_MAJOR.$V_MINOR.$((V_PATCH + 1))"
  elif [[ $1 -eq '--minor' ]]; then
      NEW_VERSION="$V_MAJOR.$((V_MINOR + 1)).$V_PATCH"
  elif [[ $1 -eq '--major' ]]; then
      NEW_VERSION="$((V_MAJOR + 1)).$V_MINOR.$V_PATCH"
  else
      NEW_VERSION="$V_MAJOR.$V_MINOR.$((V_PATCH + 1))"
  fi

  echo -e "${NOTICE_FLAG} Bumping to: ${GREEN}$NEW_VERSION"

  echo $NEW_VERSION > VERSION
  git add -- VERSION
  git commit -m "${NEW_VERSION}" -- VERSION
  git tag -a -m "${NEW_VERSION}" "v$NEW_VERSION"
  git push origin --follow-tags
else
  echo -e "${WARNING_FLAG} Could not find a VERSION file."
  exit 1;
fi

echo -e "${NOTICE_FLAG} Finished."