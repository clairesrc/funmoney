FROM python:latest
RUN mkdir /frontend && cd /frontend && mkdir js && mkdir css
WORKDIR /frontend
COPY css/* /frontend/css/
COPY js/* /frontend/js/
COPY * /frontend/
RUN ./build.sh
EXPOSE 8000
CMD ["python","-m","http.server"]