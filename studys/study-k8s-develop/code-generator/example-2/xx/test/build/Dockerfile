FROM alpine:3.8

ENV OPERATOR=/usr/local/bin/test \
    USER_UID=1001 \
    USER_NAME=test

# install operator binary
COPY build/_output/bin/test ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
