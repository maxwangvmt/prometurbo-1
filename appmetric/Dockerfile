FROM alpine:3.7

ARG GIT_COMMIT
ENV GIT_COMMIT ${GIT_COMMIT}

# build.info is generated by scripts/docker/gen.build.info.sh
COPY ./build.info /tmp/

COPY ./icon/favicon.jpg /tmp/
COPY ./_output/appMetric.linux /bin/appMetric
RUN chmod +x /bin/appMetric
EXPOSE 8081

ENTRYPOINT ["/bin/appMetric"]
