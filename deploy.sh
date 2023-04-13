#!/bin/sh
# github pages deploy

# If a command fails then the deploy stops
set -e

printf "\033[0;32mDeploying updates to GitHub...\033[0m\n"

if [ ! -d public ];
then
	mkdir public
	cd public

	git init
	git remote add origin "git@github.com:rdallman/rdallman.github.io"
	git pull origin master
	
	cd ..
fi

# Build the project.
hugo # if using a theme, replace with `hugo -t <YOURTHEME>`

# Go To Public folder
cd public

# Add changes to git.
git add .

# Commit changes.
msg="rebuilding site $(date)"
if [ -n "$*" ]; then
	msg="$*"
fi
git commit -m "$msg"

# Push source and build repos.
git push origin master
