FROM python:latest
RUN mkdir /frontend && cd /frontend
WORKDIR /frontend
COPY * /frontend/
EXPOSE 8000
CMD ["python","-m","http.server"]