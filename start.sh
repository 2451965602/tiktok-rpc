#!/bin/bash

#!/bin/bash

go run ./cmd/api &
go run ./cmd/interact &
go run ./cmd/social &
go run ./cmd/user &
go run ./cmd/video &

