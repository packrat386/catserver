FROM scratch
ADD catserver_linux_amd64 /catserver
ENV PATH=/
ENTRYPOINT ["catserver"]
