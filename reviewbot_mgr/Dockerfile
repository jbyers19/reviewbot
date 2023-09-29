FROM python

RUN mkdir reviewbot-mgr
WORKDIR /reviewbot-mgr
COPY api/*.py* api/
COPY requirements.txt .
RUN pip install -U pip
RUN pip install -r requirements.txt

ENV GRPC_PORT=50051
ENV GRPC_HOST="host.docker.internal"
ENV PYTHONPATH=/reviewbot-mgr/api

WORKDIR /reviewbot-mgr/api
EXPOSE 5000
ENTRYPOINT [ "flask", "run", "--host=0.0.0.0" ]
