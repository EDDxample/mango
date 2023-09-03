# Usage:
#   docker run --rm -v ${PWD}:/input flamegraph cpu_profile.txt output.svg

FROM debian


WORKDIR /flamegraph

RUN apt update && apt install -y git
RUN git clone https://github.com/brendangregg/FlameGraph.git . && \
    echo '#!/bin/bash\n./stackcollapse-go.pl /input/$1 | ./flamegraph.pl > /input/$2' > run.sh && \
    chmod +x run.sh

ENTRYPOINT [ "./run.sh" ]

