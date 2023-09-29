FROM alpine

RUN mkdir reviewbot
WORKDIR /reviewbot

COPY build/reviewbot .

ENV GRPC_PORT=50051
ENV TELEGRAM_API_KEY=""
EXPOSE 50051
ENTRYPOINT [ "./reviewbot" ]