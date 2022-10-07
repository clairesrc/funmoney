FROM node:alpine
RUN npm install http-server  uglify-js uglifycss -g
EXPOSE 8080
RUN mkdir /frontend && cd /frontend && mkdir js && mkdir css
WORKDIR /frontend
COPY css/* /frontend/css/
COPY js/* /frontend/js/
COPY * /frontend/
RUN ./build.sh
CMD ["http-server", "-o", "--cors"]