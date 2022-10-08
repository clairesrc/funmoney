#!/usr/bin/env sh
set -e
cd frontend

# Concatenate + uglify js
# Sections render functions followed by main app runtime.
cd js/sections
SECTIONSJS=`awk 'FNR==1{print ""}{print}' *.js`
cd ..
UTILSJS=`awk 'FNR==1{print ""}{print}' utils.main.js`
MAINJS=`awk 'FNR==1{print ""}{print}' app.main.js`
echo "$UTILSJS$SECTIONSJS$MAINJS" > main.js
uglifyjs --compress --mangle --output main.min.js main.js
mv main.min.js .. 

cd .. 

# Concatenate + uglify css. 
# Uses filenames as a way to control the order in which they are concatenated.
cd css 
STARTCSS=`awk 'FNR==1{print ""}{print}' *.start.css`
MAINCSS=`awk 'FNR==1{print ""}{print}' *.main.css` 
cd sections
SECTIONSCSS=`awk 'FNR==1{print ""}{print}' *.css`
cd ..
ENDCSS=`awk 'FNR==1{print ""}{print}' *.end.css`
echo "$STARTCSS$MAINCSS$SECTIONSCSS$ENDCSS" > main.css
uglifycss --max-line-len 500 --output main.min.css main.css 
mv main.min.css .. 

# Uglify index.html
cd .. 
html-minifier-terser --collapse-whitespace --remove-comments index.html > index.min.html
mv index.min.html index.html

# Add live-reload if dev mode enabled
if [[ "$ENV" == "dev" ]]; then
    echo "Dev mode enabled, adding dev snippet to index.html"
    INDEXCONTENT=`cat index.html`
    DEVPOSTFIX=`cat dev.html`

    echo "$INDEXCONTENT$DEVPOSTFIX" > index.html
fi

# Prepare favicons
mv favicons/* .