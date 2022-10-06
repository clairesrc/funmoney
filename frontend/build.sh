#!/usr/bin/env sh

# TODO: uglification
awk 'FNR==1{print ""}{print}' *.js > main.js
awk 'FNR==1{print ""}{print}' *.css > style.css

