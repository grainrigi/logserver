if ! [ -e /workspace/.env ]; then
  cp /workspace/example.env /workspace/.env
fi

source /workspace/.env

alias migrate-dev='migrate -database "${DB_URL}" -path /workspace/migrations'

export LANG=C
