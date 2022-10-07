#!/usr/bin/env sh
# TODO: uglification, preprocessing etc
cd js && \
awk 'FNR==1{print ""}{print}' *.js > main.js && \
uglifyjs --compress --mangle --output main.min.js main.js && \
mv main.min.js .. && \

cd .. && \

cd css && \
awk 'FNR==1{print ""}{print}' *.css > main.css && \
uglifycss --max-line-len 500 --output main.min.css main.css && \
mv main.min.css ..
