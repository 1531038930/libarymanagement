FROM alpine:3.12
LABEL authors="15310"
WORKDIR /app
COPY librarymanagement /app
COPY static/ /app/static
EXPOSE 8080
ENTRYPOINT ["./librarymanagement"]