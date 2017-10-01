#!/usr/bin/env bash

#creates the css
sass --update resources/app/static/sass:resources/app/static/css --style compressed --scss

#builds the application
astilectron-bundler -v
