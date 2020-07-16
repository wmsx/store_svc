FROM alpine
ADD store_svc-service /store_svc-service
ENTRYPOINT [ "/store_svc-service" ]
