#!/usr/bin/env sh

# TODO: uglification, preprocessing etc
cd js
awk 'FNR==1{print ""}{print}' *.js > main.js
mv main.js ..

cd ..

cd css
awk 'FNR==1{print ""}{print}' *.css > main.css
mv main.css ..
