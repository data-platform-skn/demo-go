FROM scratch
EXPOSE 8080
ENTRYPOINT ["/demo-go"]
COPY ./bin/ /