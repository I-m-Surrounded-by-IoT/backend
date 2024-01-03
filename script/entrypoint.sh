#!/bin/bash

chown -R ${PUID}:${PGID} /backend

umask ${UMASK}

exec su-exec ${PUID}:${PGID} backend $@
