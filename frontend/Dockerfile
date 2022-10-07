FROM node:alpine
RUN npm install http-server  uglify-js uglifycss html-minifier-terser -g
EXPOSE 8080
ENV ENV "dev"
RUN mkdir -p /frontend/favicons /frontend/js /frontend/css
WORKDIR /frontend
COPY css/sections/* /frontend/css/sections/
COPY css/* /frontend/css/
COPY js/sections/* /frontend/js/sections/
COPY js/* /frontend/js/
COPY favicons/* /frontend/favicons/
COPY build.sh /build.sh
COPY dev.html /frontend/dev.html
COPY index.html /frontend/index.html
RUN /build.sh
CMD ["http-server", "-o", "--cors"]
